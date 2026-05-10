package tic

import (
	"strings"
	"testing"
)

func TestBox(t *testing.T) {
	got := Box("TITLE", "OK", BoxRetro)

	if !strings.Contains(got, "TITLE") {
		t.Fatalf("expected title in box, got %q", got)
	}

	if !strings.Contains(got, "OK") {
		t.Fatalf("expected body in box, got %q", got)
	}
}
