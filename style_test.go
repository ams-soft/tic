package tic

import "testing"

func TestStyleBuilder(t *testing.T) {
	EnableColor()

	got := Style().Fg(ColorLightGreen).Bg(ColorBlack).Bold().Sprint("OK")
	want := "\x1b[92;40;1mOK\x1b[0m"

	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestNoColor(t *testing.T) {
	DisableColor()
	defer EnableColor()

	got := Style().Fg(ColorRed).Bold().Sprint("OK")
	if got != "OK" {
		t.Fatalf("expected plain text, got %q", got)
	}
}
