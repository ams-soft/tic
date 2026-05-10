package tic

import (
	"bytes"
	"testing"
)

func TestRendererUsesWriter(t *testing.T) {
	var buf bytes.Buffer

	r := NewRenderer(&buf)
	r.ClearScreen()

	if got := buf.String(); got != "\x1b[2J\x1b[H" {
		t.Fatalf("expected clear screen sequence, got %q", got)
	}
}

func TestPackageClearScreenUsesConfiguredWriter(t *testing.T) {
	var buf bytes.Buffer

	SetOutput(&buf)
	defer ResetOutput()

	ClearScreen()

	if got := buf.String(); got != "\x1b[2J\x1b[H" {
		t.Fatalf("expected clear screen sequence, got %q", got)
	}
}

func TestRendererNoColorPrintLine(t *testing.T) {
	var buf bytes.Buffer

	r := NewRenderer(&buf, WithNoColor(true))
	r.PrintLine("OK", ColorGreen, ColorBlack)

	if got := buf.String(); got != "OK\n" {
		t.Fatalf("expected plain line, got %q", got)
	}
}
