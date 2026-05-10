package tic

import "testing"

func TestColorString(t *testing.T) {
	if ColorC64BG.String() != "C64BG" {
		t.Fatalf("expected C64BG, got %q", ColorC64BG.String())
	}
}

func TestInvalidColorFallbacks(t *testing.T) {
	var c Color = 999

	if got := ansiFgSeq(c); got != "39" {
		t.Fatalf("expected default foreground 39, got %q", got)
	}

	if got := ansiBgSeq(c); got != "49" {
		t.Fatalf("expected default background 49, got %q", got)
	}
}

func TestColorText(t *testing.T) {
	EnableColor()

	got := ColorText("ok", ColorRed)
	want := "\x1b[31mok\x1b[0m"

	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}
