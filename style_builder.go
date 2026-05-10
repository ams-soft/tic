package tic

import "strings"

// StyleBuilder permite compor estilos (cores + atributos ANSI) de forma fluente.
type StyleBuilder struct {
	fg *Color
	bg *Color

	bold      bool
	dim       bool
	italic    bool
	underline bool
	blink     bool
	reverse   bool
	strike    bool
}

// Style creates a new fluent style builder.
func Style() *StyleBuilder { return &StyleBuilder{} }

// NewStyle creates a new fluent style builder.
func NewStyle() *StyleBuilder { return &StyleBuilder{} }

// Fg define a cor de foreground.
func (s *StyleBuilder) Fg(c Color) *StyleBuilder {
	s.fg = &c
	return s
}

// Bg define a cor de background.
func (s *StyleBuilder) Bg(c Color) *StyleBuilder {
	s.bg = &c
	return s
}

func (s *StyleBuilder) Bold() *StyleBuilder {
	s.bold = true
	return s
}

func (s *StyleBuilder) Dim() *StyleBuilder {
	s.dim = true
	return s
}

func (s *StyleBuilder) Italic() *StyleBuilder {
	s.italic = true
	return s
}

func (s *StyleBuilder) Underline() *StyleBuilder {
	s.underline = true
	return s
}

func (s *StyleBuilder) Blink() *StyleBuilder {
	s.blink = true
	return s
}

func (s *StyleBuilder) Reverse() *StyleBuilder {
	s.reverse = true
	return s
}

func (s *StyleBuilder) Strike() *StyleBuilder {
	s.strike = true
	return s
}

func (s *StyleBuilder) Codes() []string {
	codes := make([]string, 0, 9)

	if s.fg != nil {
		codes = append(codes, ansiFgSeq(*s.fg))
	}
	if s.bg != nil {
		codes = append(codes, ansiBgSeq(*s.bg))
	}
	if s.bold {
		codes = append(codes, "1")
	}
	if s.dim {
		codes = append(codes, "2")
	}
	if s.italic {
		codes = append(codes, "3")
	}
	if s.underline {
		codes = append(codes, "4")
	}
	if s.blink {
		codes = append(codes, "5")
	}
	if s.reverse {
		codes = append(codes, "7")
	}
	if s.strike {
		codes = append(codes, "9")
	}

	return codes
}

// Sprint aplica o estilo ao texto e retorna ANSI + reset.
func (s *StyleBuilder) Sprint(text string) string {
	if DefaultRenderer().NoColor() {
		return text
	}

	codes := s.Codes()
	if len(codes) == 0 {
		return text
	}

	return "\x1b[" + strings.Join(codes, ";") + "m" + text + "\x1b[0m"
}

func (s *StyleBuilder) Sprintln(text string) string {
	return s.Sprint(text) + "\n"
}
