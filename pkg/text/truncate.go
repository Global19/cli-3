package text

import (
	"github.com/rivo/uniseg"
)

// DisplayWidth calculates what the rendered width of a string may be
func DisplayWidth(s string) int {
	return uniseg.GraphemeClusterCount(s)
}

const (
	ellipsisWidth       = 3
	minWidthForEllipsis = 5
)

// Truncate shortens a string to fit the maximum display width
func Truncate(max int, s string) string {
	w := DisplayWidth(s)
	if w <= max {
		return s
	}

	useEllipsis := false
	if max >= minWidthForEllipsis {
		useEllipsis = true
		max -= ellipsisWidth
	}

	gr := uniseg.NewGraphemes(s)
	res := ""
	for i := 0; i <= max; i++ {
		gr.Next()
		res += gr.Str()
	}

	if useEllipsis {
		res += "..."
	}
	return res
}
