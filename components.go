package tic

import "strings"

type BoxStyle struct {
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
	Horizontal  string
	Vertical    string
}

var (
	BoxSingle = BoxStyle{"┌", "┐", "└", "┘", "─", "│"}
	BoxDouble = BoxStyle{"╔", "╗", "╚", "╝", "═", "║"}
	BoxRetro  = BoxStyle{"┌", "┐", "└", "┘", "─", "│"}
)

func Box(title string, body string, style BoxStyle) string {
	lines := strings.Split(body, "\n")
	width := VisibleLen(title)

	for _, line := range lines {
		if l := VisibleLen(line); l > width {
			width = l
		}
	}

	top := style.TopLeft + strings.Repeat(style.Horizontal, width+2) + style.TopRight
	titleLine := style.Vertical + " " + PadRightVisible(title, width) + " " + style.Vertical
	sep := style.Vertical + " " + strings.Repeat(style.Horizontal, width) + " " + style.Vertical

	out := []string{top, titleLine, sep}
	for _, line := range lines {
		out = append(out, style.Vertical+" "+PadRightVisible(line, width)+" "+style.Vertical)
	}

	bottom := style.BottomLeft + strings.Repeat(style.Horizontal, width+2) + style.BottomRight
	out = append(out, bottom)

	return strings.Join(out, "\n")
}

func Header(title string) string {
	return Box(title, "AMS TIC", BoxRetro)
}
