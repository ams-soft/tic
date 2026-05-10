package tic

func styleCode(code string, text string) string {
	if DefaultRenderer().NoColor() {
		return text
	}

	return "\x1b[" + code + "m" + text + "\x1b[0m"
}

// Helpers de estilo ANSI simples.
func Bold(t string) string      { return styleCode("1", t) }
func Dim(t string) string       { return styleCode("2", t) }
func Italic(t string) string    { return styleCode("3", t) }
func Underline(t string) string { return styleCode("4", t) }
func Blink(t string) string     { return styleCode("5", t) }
func Reverse(t string) string   { return styleCode("7", t) }
func Strike(t string) string    { return styleCode("9", t) }

// ResetStyle returns the ANSI reset sequence unless color is disabled.
func ResetStyle() string {
	if DefaultRenderer().NoColor() {
		return ""
	}

	return "\x1b[0m"
}

func WithColors(text string, fg Color, bg Color) string {
	if DefaultRenderer().NoColor() {
		return text
	}

	return ansiSeq(ansiFgSeq(fg), ansiBgSeq(bg)) + text + "\x1b[0m"
}

func ColorText(text string, fg Color) string {
	if DefaultRenderer().NoColor() {
		return text
	}

	return ansiSeq(ansiFgSeq(fg)) + text + "\x1b[0m"
}

func ColorBg(text string, bg Color) string {
	if DefaultRenderer().NoColor() {
		return text
	}

	return ansiSeq(ansiBgSeq(bg)) + text + "\x1b[0m"
}

func Colorize(text string, fg Color, bg Color) string {
	return WithColors(text, fg, bg)
}
