package client

import (
	_ "embed"
	"encoding/json"
	"path"
	"strings"
	"sync"
)

//go:embed matrix.json
var matrixJSON []byte

// matrixRule says what a property is called on a family of boards, or that the
// hardware does not have it at all.
type matrixRule struct {
	Boards []string `json:"boards"`
	Wire   string   `json:"wire"`
	Absent string   `json:"absent"`
}

type deviceMatrix struct {
	Menus map[string]map[string][]matrixRule `json:"menus"`
}

var (
	matrixOnce sync.Once
	matrix     deviceMatrix
)

func loadMatrix() deviceMatrix {
	matrixOnce.Do(func() {
		_ = json.Unmarshal(matrixJSON, &matrix)
	})
	return matrix
}

// boardMatches supports a leading "!" for negation, so a property present on
// exactly one family is one rule rather than a list of every other board.
func boardMatches(patterns []string, board string) bool {
	for _, p := range patterns {
		neg := strings.HasPrefix(p, "!")
		ok, _ := path.Match(strings.TrimPrefix(p, "!"), board)
		if neg {
			if ok {
				return false
			}
			continue
		}
		if ok {
			return true
		}
	}
	for _, p := range patterns {
		if strings.HasPrefix(p, "!") {
			return true
		}
	}
	return false
}

// resolveProperty maps a property to the name this board uses, or reports that
// the board does not have it.
func resolveProperty(menuPath, prop, board string) (wire, absent string) {
	props, ok := loadMatrix().Menus[menuPath]
	if !ok || board == "" {
		return prop, ""
	}
	for _, r := range props[prop] {
		if !boardMatches(r.Boards, board) {
			continue
		}
		if r.Absent != "" {
			return prop, r.Absent
		}
		if r.Wire != "" {
			return r.Wire, ""
		}
	}
	return prop, ""
}

// KnownAbsent reports why menuPath.prop does not exist on board, or "".
func KnownAbsent(menuPath, prop, board string) string {
	_, absent := resolveProperty(menuPath, prop, board)
	return absent
}

// applyMatrix rewrites body's property names to what this board calls them, and
// reports the ones the board is known not to have.
func applyMatrix(menuPath, board string, body Object) (Object, map[string]string) {
	if board == "" || len(body) == 0 {
		return body, nil
	}
	if _, ok := loadMatrix().Menus[menuPath]; !ok {
		return body, nil
	}
	out := make(Object, len(body))
	var absent map[string]string
	for k, v := range body {
		wire, why := resolveProperty(menuPath, k, board)
		if why != "" {
			if absent == nil {
				absent = map[string]string{}
			}
			absent[k] = why
			continue
		}
		out[wire] = v
	}
	return out, absent
}
