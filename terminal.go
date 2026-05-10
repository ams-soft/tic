package tic

func SetFg(c Color) {
	DefaultRenderer().SetFg(c)
}

func SetBg(c Color) {
	DefaultRenderer().SetBg(c)
}

// Reset resets all terminal attributes.
func Reset() {
	DefaultRenderer().Reset()
}

// ResetAll is a semantic alias for Reset.
func ResetAll() {
	Reset()
}

func ResetFg() {
	DefaultRenderer().ResetFg()
}

func ResetBg() {
	DefaultRenderer().ResetBg()
}

func ResetScreenBg() {
	DefaultRenderer().ResetScreenBg()
}

func ClearScreenWithBg(c Color) {
	DefaultRenderer().ClearScreenWithBg(c)
}

func ClearScreen() {
	DefaultRenderer().ClearScreen()
}

func PrintLine(line string, fg Color, bg Color) {
	DefaultRenderer().PrintLine(line, fg, bg)
}

func PrintANSI256Table() {
	DefaultRenderer().PrintANSI256Table()
}
