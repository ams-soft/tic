package tic

type Theme struct {
	Name       string
	Foreground Color
	Background Color
	Accent     Color
	Success    Color
	Warning    Color
	Error      Color
	Info       Color
}

func ThemeC64() Theme {
	return Theme{
		Name:       "C64",
		Foreground: ColorC64FG,
		Background: ColorC64BG,
		Accent:     ColorLightPurple,
		Success:    ColorLightGreen,
		Warning:    ColorYellow,
		Error:      ColorLightRed,
		Info:       ColorCyan,
	}
}

func ThemeAtari800() Theme {
	return Theme{
		Name:       "Atari800",
		Foreground: ColorA800FG,
		Background: ColorA800BG,
		Accent:     ColorCyan,
		Success:    ColorLightGreen,
		Warning:    ColorYellow,
		Error:      ColorLightRed,
		Info:       ColorLightBlue,
	}
}

func ThemeCRT() Theme {
	return Theme{
		Name:       "CRT",
		Foreground: ColorGreen,
		Background: ColorBlack,
		Accent:     ColorLightGreen,
		Success:    ColorLightGreen,
		Warning:    ColorYellow,
		Error:      ColorLightRed,
		Info:       ColorCyan,
	}
}

func (t Theme) Title(text string) string {
	return Style().Fg(t.Foreground).Bg(t.Background).Bold().Sprint(text)
}

func (t Theme) SuccessText(text string) string {
	return Style().Fg(t.Success).Bg(t.Background).Bold().Sprint(text)
}

func (t Theme) WarningText(text string) string {
	return Style().Fg(t.Warning).Bg(t.Background).Bold().Sprint(text)
}

func (t Theme) ErrorText(text string) string {
	return Style().Fg(t.Error).Bg(t.Background).Bold().Sprint(text)
}

func (t Theme) InfoText(text string) string {
	return Style().Fg(t.Info).Bg(t.Background).Sprint(text)
}
