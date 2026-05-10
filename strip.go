package tic

import (
	"regexp"

	"golang.org/x/term"
)

var ansiRegex = regexp.MustCompile(`[\x1b\x9b][[\]()#;?]*(?:(?:(?:[a-zA-Z\d]*(?:;[a-zA-Z\d]*)*)?\x07)|(?:(?:\d{1,4}(?:;\d{0,4})*)?[\dA-PR-TZcf-nq-uy=><~]))`)

// StripANSI removes common ANSI/VT escape sequences from a string.
func StripANSI(s string) string {
	return ansiRegex.ReplaceAllString(s, "")
}

// HasANSI reports whether the string contains ANSI/VT escape sequences.
func HasANSI(s string) bool {
	return ansiRegex.MatchString(s)
}

// VisibleLen returns the rune count after removing ANSI sequences.
//
// This intentionally counts runes, not display cells. For most TIC headers
// this is enough; wide Unicode cell handling can be added later without
// breaking the API.
func VisibleLen(s string) int {
	return len([]rune(StripANSI(s)))
}

// PadRightVisible pads s to width based on VisibleLen.
func PadRightVisible(s string, width int) string {
	l := VisibleLen(s)
	if l >= width {
		return s
	}

	return s + repeatSpace(width-l)
}

// CenterVisible centers s within width based on VisibleLen.
func CenterVisible(s string, width int) string {
	l := VisibleLen(s)
	if l >= width {
		return s
	}

	left := (width - l) / 2
	return repeatSpace(left) + s
}

func repeatSpace(n int) string {
	if n <= 0 {
		return ""
	}

	b := make([]byte, n)
	for i := range b {
		b[i] = ' '
	}
	return string(b)
}

// TerminalSize returns width and height for a file descriptor.
func TerminalSize(fd int) (width int, height int, err error) {
	return term.GetSize(fd)
}
