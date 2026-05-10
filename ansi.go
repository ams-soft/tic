package tic

import (
	"fmt"
	"strconv"
	"strings"
)

// ansiFgSeq returns the ANSI foreground sequence body without "\x1b[" and "m".
func ansiFgSeq(c Color) string {
	switch c {
	case ColorBlack:
		return "30"
	case ColorWhite:
		return "97"
	case ColorRed:
		return "31"
	case ColorCyan:
		return "36"
	case ColorPurple:
		return "38;5;93"
	case ColorGreen:
		return "32"
	case ColorBlue:
		return "34"
	case ColorYellow:
		return "33"
	case ColorOrange:
		return "38;5;208"
	case ColorBrown:
		return "38;5;130"
	case ColorLightRed:
		return "91"
	case ColorDarkGray:
		return "90"
	case ColorMediumGray:
		return "37"
	case ColorLightGreen:
		return "92"
	case ColorLightBlue:
		return "94"
	case ColorLightGray:
		return "37"
	case ColorLightPurple:
		return "38;5;135"
	case ColorDarkPurple:
		return "38;5;54"
	case ColorC64BG:
		return "38;5;18"
	case ColorC64FG:
		return "38;5;63"
	case ColorA800BG:
		return "38;5;23"
	case ColorA800FG:
		return "38;5;43"
	default:
		return "39"
	}
}

// ansiBgSeq returns the ANSI background sequence body without "\x1b[" and "m".
func ansiBgSeq(c Color) string {
	fg := ansiFgSeq(c)

	if strings.HasPrefix(fg, "38;") {
		return strings.Replace(fg, "38;", "48;", 1)
	}

	n, err := strconv.Atoi(fg)
	if err != nil {
		return "49"
	}

	if n >= 30 && n <= 37 {
		return strconv.Itoa(40 + (n - 30))
	}
	if n >= 90 && n <= 97 {
		return strconv.Itoa(100 + (n - 90))
	}

	return "49"
}

func ansiSeq(codes ...string) string {
	filtered := make([]string, 0, len(codes))
	for _, code := range codes {
		if code != "" {
			filtered = append(filtered, code)
		}
	}

	if len(filtered) == 0 {
		return ""
	}

	return "\x1b[" + strings.Join(filtered, ";") + "m"
}

func ansi256ToRGB(n int) (int, int, int) {
	switch {
	case n < 0:
		return 0, 0, 0

	case n < 16:
		base := []struct{ r, g, b int }{
			{0, 0, 0},
			{128, 0, 0},
			{0, 128, 0},
			{128, 128, 0},
			{0, 0, 128},
			{128, 0, 128},
			{0, 128, 128},
			{192, 192, 192},
			{128, 128, 128},
			{255, 0, 0},
			{0, 255, 0},
			{255, 255, 0},
			{0, 0, 255},
			{255, 0, 255},
			{0, 255, 255},
			{255, 255, 255},
		}
		c := base[n]
		return c.r, c.g, c.b

	case n >= 16 && n <= 231:
		n -= 16
		r := n / 36
		g := (n % 36) / 6
		b := n % 6

		toVal := func(v int) int {
			if v == 0 {
				return 0
			}
			return 55 + 40*v
		}

		return toVal(r), toVal(g), toVal(b)

	case n >= 232 && n <= 255:
		gray := 8 + 10*(n-232)
		return gray, gray, gray

	default:
		return 255, 255, 255
	}
}

// ANSI256Sample returns a printable line describing an ANSI 256 color.
func ANSI256Sample(i int) string {
	r, g, b := ansi256ToRGB(i)
	code := fmt.Sprintf("38;5;%d", i)

	return fmt.Sprintf(
		"\x1b[%sm♦\x1b[0m  %3d  rgb(%3d,%3d,%3d)   %s",
		code, i, r, g, b, code,
	)
}
