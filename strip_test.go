package tic

import "testing"

func TestStripANSI(t *testing.T) {
	got := StripANSI("\x1b[31mred\x1b[0m")
	if got != "red" {
		t.Fatalf("expected red, got %q", got)
	}
}

func TestStripANSICursorSequence(t *testing.T) {
	got := StripANSI("\x1b[2J\x1b[Hhello")
	if got != "hello" {
		t.Fatalf("expected hello, got %q", got)
	}
}

func TestVisibleLen(t *testing.T) {
	got := VisibleLen("\x1b[31mred\x1b[0m")
	if got != 3 {
		t.Fatalf("expected 3, got %d", got)
	}
}

func TestCenterVisible(t *testing.T) {
	got := CenterVisible("\x1b[31mX\x1b[0m", 5)
	want := "  \x1b[31mX\x1b[0m"

	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}
