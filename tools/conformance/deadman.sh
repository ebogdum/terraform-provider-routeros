#!/usr/bin/env bash
#
# Dead-man switch for live-device conformance runs.
#
# A conformance sweep mutates real menus on a real router. Any one of those
# writes can sever management access -- a firewall rule, an address change, a
# disabled service. This installs a watchdog on the device itself: a scheduler
# decrements a counter every minute, and if nothing re-arms it the router
# restores the pre-run backup and reboots.
#
# The backup is taken BEFORE the switch is installed, so restoring it also
# removes the switch. That is what makes a reboot loop impossible.
#
# Layers, because one is not enough:
#   1. the dead-man switch itself (recovers from anything, ~45s downtime)
#   2. MGMT-ALWAYS-ALLOW at the head of the input chain (survives a test that
#      adds a drop rule)
#   3. a static secondary address (survives a test that clobbers the primary,
#      which matters here because the primary is usually DHCP-assigned)
#   4. two services, SSH and REST, re-armed independently
#
# Usage:
#   ROUTEROS_HOST=192.168.10.2 ROUTEROS_PASSWORD=... ./deadman.sh backup
#   ...                                              ./deadman.sh install
#   ...                                              ./deadman.sh arm      # blocks
#   ...                                              ./deadman.sh disarm
#   ...                                              ./deadman.sh verify   # REBOOTS
#   ...                                              ./deadman.sh remove
#
# `verify` deliberately lets the switch expire. Run it once before trusting a
# campaign to the switch -- an untested watchdog is worse than none, because it
# buys confidence it has not earned.

set -euo pipefail

HOST="${ROUTEROS_HOST:?ROUTEROS_HOST is required}"
HOST="${HOST#http://}"; HOST="${HOST#https://}"
USER="${ROUTEROS_USER:-admin}"
PASS="${ROUTEROS_PASSWORD:?ROUTEROS_PASSWORD is required}"
SECOND_IP="${ROUTEROS_SECONDARY_IP:-192.168.10.3/24}"
GRACE="${ROUTEROS_DMS_GRACE:-5}"          # minutes without a re-arm before restore
STATE="${ROUTEROS_DMS_STATE:-.conformance}"

ros()  { ssh -o BatchMode=yes -o ConnectTimeout=15 "$USER@$HOST" "$@"; }
rest() { curl -s -m 10 -u "$USER:$PASS" "http://$2/rest$1" ; }

cmd_backup() {
  mkdir -p "$STATE"
  local ts; ts="$(date +%Y%m%d-%H%M%S)"
  ros "/system/backup/save name=preconf-$ts dont-encrypt=yes; :delay 3s; /export file=preconf-$ts; :delay 2s"
  scp -O -o BatchMode=yes "$USER@$HOST:preconf-$ts.backup" "$STATE/" >/dev/null
  scp -O -o BatchMode=yes "$USER@$HOST:preconf-$ts.rsc"    "$STATE/" >/dev/null
  echo "$ts" > "$STATE/LATEST"
  echo "backup preconf-$ts saved on device and pulled to $STATE/"
}

cmd_install() {
  local ts; ts="$(cat "$STATE/LATEST")"
  local myip; myip="$(ros ':put [/ip/arp/get [find where mac-address=[/interface/ethernet/get [find] mac-address]] address]' 2>/dev/null || true)"
  myip="${ROUTEROS_MGMT_IP:-$(ip_of_this_host)}"
  cat > "$STATE/dms-install.rsc" <<EOF
/system/script/remove [find name~"^dms-"]
/system/scheduler/remove [find name="dms"]
/system/script/add name=dms-rearm policy=read,write,policy,test source=":global dmsCounter $GRACE"
/system/script/add name=dms-tick policy=read,write,policy,test,reboot,ftp,password,sensitive source={
  :global dmsCounter
  :if ([:typeof \$dmsCounter] = "nothing") do={ :set dmsCounter $GRACE }
  :set dmsCounter (\$dmsCounter - 1)
  :if (\$dmsCounter <= 0) do={
    :log error "DMS EXPIRED - restoring preconf-$ts and rebooting"
    /system/backup/load name="preconf-$ts" password=""
  }
}
/system/scheduler/add name=dms disabled=yes interval=1m policy=read,write,policy,test,reboot,ftp,password,sensitive on-event="/system/script/run dms-tick"
/ip/firewall/filter/remove [find comment="MGMT-ALWAYS-ALLOW"]
/ip/firewall/filter/add chain=input action=accept src-address=$myip place-before=0 comment="MGMT-ALWAYS-ALLOW"
:if ([:len [/ip/address/find where comment="MGMT-SECONDARY"]] = 0) do={
  /ip/address/add address=$SECOND_IP interface=[/ip/address/get [find] interface] comment="MGMT-SECONDARY"
}
:put "DMS installed (disabled)"
EOF
  scp -O -o BatchMode=yes "$STATE/dms-install.rsc" "$USER@$HOST:dms-install.rsc" >/dev/null
  ros '/import file=dms-install.rsc'
}

ip_of_this_host() {
  # the address this machine actually reaches the router from
  local dev; dev="$(route -n get "$HOST" 2>/dev/null | awk '/interface:/{print $2}')"
  ifconfig "$dev" 2>/dev/null | awk '/inet /{print $2; exit}'
}

cmd_arm() {
  ros '/system/scheduler/enable dms; /system/script/run dms-rearm'
  local second; second="${SECOND_IP%%/*}"
  echo "DMS ARMED (${GRACE}m grace). Re-arming over SSH and REST until interrupted."
  trap 'cmd_disarm; exit 0' INT TERM
  while true; do
    ros '/system/script/run dms-rearm' 2>/dev/null || true
    curl -s -m 5 -u "$USER:$PASS" -X POST "http://$second/rest/system/script/run" \
         -H 'Content-Type: application/json' -d '{".id":"dms-rearm"}' >/dev/null 2>&1 || true
    sleep 20
  done
}

cmd_disarm() { ros '/system/scheduler/disable dms; /system/script/run dms-rearm; :put "DMS DISARMED"'; }

cmd_verify() {
  echo "This lets the switch expire: the router WILL restore and reboot (~45s)."
  ros '/system/scheduler/enable dms'
  ros ':global dmsCounter 1'
  local t0; t0=$SECONDS
  until ! curl -s -m 3 -u "$USER:$PASS" "http://$HOST/rest/system/resource" >/dev/null 2>&1; do
    sleep 2; [ $((SECONDS-t0)) -gt 180 ] && { echo "FAIL: switch never fired"; exit 1; }
  done
  echo "  switch fired at +$((SECONDS-t0))s, router going down"
  until curl -s -m 3 -u "$USER:$PASS" "http://$HOST/rest/system/resource" >/dev/null 2>&1; do
    sleep 3; [ $((SECONDS-t0)) -gt 420 ] && { echo "FAIL: router did not return"; exit 1; }
  done
  echo "  router back at +$((SECONDS-t0))s"
  ros '/export file=postdms'; sleep 2
  scp -O -o BatchMode=yes "$USER@$HOST:postdms.rsc" "$STATE/" >/dev/null
  ros '/file/remove [find name="postdms.rsc"]'
  if diff <(grep -v '^#' "$STATE/preconf-$(cat "$STATE/LATEST").rsc") \
          <(grep -v '^#' "$STATE/postdms.rsc") >/dev/null; then
    echo "PASS: restored config is identical to the pre-run baseline"
  else
    echo "FAIL: restored config differs from baseline"; exit 1
  fi
}

cmd_remove() {
  ros '/system/scheduler/remove [find name="dms"]
       /system/script/remove [find name~"^dms-"]
       /ip/firewall/filter/remove [find comment="MGMT-ALWAYS-ALLOW"]
       /ip/address/remove [find comment="MGMT-SECONDARY"]
       /file/remove [find name~"^preconf-"]
       /file/remove [find name="dms-install.rsc"]
       :put "DMS fully removed, device clean"'
}

case "${1:-}" in
  backup) cmd_backup ;;  install) cmd_install ;;  arm) cmd_arm ;;
  disarm) cmd_disarm ;;  verify)  cmd_verify  ;;  remove) cmd_remove ;;
  *) sed -n '2,40p' "$0"; exit 2 ;;
esac
