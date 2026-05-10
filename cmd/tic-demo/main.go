package main

import (
	"fmt"
	"os"

	"github.com/ams-soft/tic"
)

func AMSLogo() string {

	symbol := "⫽"
	return fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		tic.Style().Bg(tic.ColorC64BG).Fg(tic.ColorLightRed).Italic().Bold().Sprint(symbol),
		tic.Style().Bg(tic.ColorC64BG).Fg(tic.ColorOrange).Italic().Bold().Sprint(symbol),
		tic.Style().Bg(tic.ColorC64BG).Fg(tic.ColorYellow).Italic().Bold().Sprint(symbol),
		tic.Style().Bg(tic.ColorC64BG).Fg(tic.ColorLightGreen).Italic().Bold().Sprint(symbol),
		tic.Style().Bg(tic.ColorC64BG).Fg(tic.ColorLightBlue).Italic().Bold().Sprint(symbol),
		tic.Style().Bg(tic.ColorC64BG).Fg(tic.ColorWhite).Italic().Bold().Sprint("A"),
		tic.Style().Bg(tic.ColorC64BG).Fg(tic.ColorWhite).Italic().Bold().Sprint("MS"))
}

func main() {
	theme := tic.ThemeC64()

	r := tic.NewRenderer(os.Stdout)
	r.ClearScreenWithBg(theme.Background)

	fmt.Println(fmt.Sprintf("%s%s", AMSLogo(), tic.Style().Fg(tic.ColorWhite).Bg(tic.ColorC64BG).Bold().Sprint(" TIC – RETRO TERMINAL TOOLKIT ")))
	fmt.Println()
	fmt.Println(tic.Box(
		"♦ AMS TIC ♦",
		theme.InfoText("READY.")+"\n"+
			theme.SuccessText("SYSTEM ONLINE")+"\n"+
			theme.WarningText("VIC-II PALETTE ACTIVE"),
		tic.BoxRetro,
	))
	fmt.Println()

	for _, ce := range tic.ColorList {

		fmt.Printf("%s %-14s %s\n",

			tic.ColorText("■", ce.Color),
			ce.Name,
			tic.ColorText("sample", ce.Color),
		)
	}

	fmt.Println()

	fmt.Println(tic.Style().Fg(tic.ColorLightGreen).Bg(tic.ColorBlack).Bold().Sprint("SUCCESS: demo completed"))
	tic.ResetAll()
}
