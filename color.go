package tic

// ┌─┴─┴─┴─┴─┴─┴─┴─┴─┐
// 🭲○  ♦ AMS TIC ♦   🭲
// └─┬─┬─┬─┬─┬─┬─┬─┬─┘

// Color represents a TIC color mapped to ANSI 16/256-color sequences.
//
// The public constants keep V1 compatibility while V2 adds validation,
// stable names, and safer fallbacks for invalid values.
type Color int

const (
	ColorBlack Color = iota
	ColorWhite
	ColorRed
	ColorCyan
	ColorPurple
	ColorGreen
	ColorBlue
	ColorYellow
	ColorOrange
	ColorBrown
	ColorLightRed
	ColorDarkGray
	ColorMediumGray
	ColorLightGreen
	ColorLightBlue
	ColorLightGray
	ColorLightPurple
	ColorDarkPurple
	ColorC64BG
	ColorC64FG
	ColorA800BG
	ColorA800FG
)

type ColorEntry struct {
	Color Color
	Name  string
}

// ColorList exposes the built-in TIC palette in a stable display order.
var ColorList = []ColorEntry{
	{ColorBlack, "Black"},
	{ColorWhite, "White"},
	{ColorRed, "Red"},
	{ColorCyan, "Cyan"},
	{ColorPurple, "Purple"},
	{ColorGreen, "Green"},
	{ColorBlue, "Blue"},
	{ColorYellow, "Yellow"},
	{ColorOrange, "Orange"},
	{ColorBrown, "Brown"},
	{ColorLightRed, "LightRed"},
	{ColorDarkGray, "DarkGray"},
	{ColorMediumGray, "MediumGray"},
	{ColorLightGreen, "LightGreen"},
	{ColorLightBlue, "LightBlue"},
	{ColorLightGray, "LightGray"},
	{ColorLightPurple, "LightPurple"},
	{ColorDarkPurple, "DarkPurple"},
	{ColorC64BG, "C64BG"},
	{ColorC64FG, "C64FG"},
	{ColorA800BG, "A800BG"},
	{ColorA800FG, "A800FG"},
}

// Valid reports whether c is one of TIC's built-in colors.
func (c Color) Valid() bool {
	return c >= ColorBlack && c <= ColorA800FG
}

// String returns the public TIC color name.
func (c Color) String() string {
	for _, entry := range ColorList {
		if entry.Color == c {
			return entry.Name
		}
	}
	return "Invalid"
}
