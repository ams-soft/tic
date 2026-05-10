package tic

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var (
	defaultRendererMu sync.RWMutex
	defaultRenderer   = NewRenderer(os.Stdout)
)

type RendererOption func(*Renderer)

// Renderer writes TIC terminal output to a specific io.Writer.
//
// Prefer Renderer for applications, tests, libraries, and servers.
// Package-level functions remain available for quick scripts and V1 compatibility.
type Renderer struct {
	out     io.Writer
	noColor bool
}

func NewRenderer(w io.Writer, opts ...RendererOption) *Renderer {
	if w == nil {
		w = io.Discard
	}

	r := &Renderer{out: w}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func WithNoColor(noColor bool) RendererOption {
	return func(r *Renderer) {
		r.noColor = noColor
	}
}

func WithOutput(w io.Writer) RendererOption {
	return func(r *Renderer) {
		if w == nil {
			r.out = io.Discard
			return
		}
		r.out = w
	}
}

func (r *Renderer) SetOutput(w io.Writer) {
	if w == nil {
		r.out = io.Discard
		return
	}
	r.out = w
}

func (r *Renderer) Output() io.Writer {
	return r.out
}

func (r *Renderer) SetNoColor(noColor bool) {
	r.noColor = noColor
}

func (r *Renderer) NoColor() bool {
	return r.noColor
}

func (r *Renderer) Print(a ...any) (int, error) {
	return fmt.Fprint(r.out, a...)
}

func (r *Renderer) Println(a ...any) (int, error) {
	return fmt.Fprintln(r.out, a...)
}

func (r *Renderer) Printf(format string, a ...any) (int, error) {
	return fmt.Fprintf(r.out, format, a...)
}

func (r *Renderer) SetFg(c Color) {
	if r.noColor {
		return
	}
	_, _ = fmt.Fprintf(r.out, "\x1b[%sm", ansiFgSeq(c))
}

func (r *Renderer) SetBg(c Color) {
	if r.noColor {
		return
	}
	_, _ = fmt.Fprintf(r.out, "\x1b[%sm", ansiBgSeq(c))
}

func (r *Renderer) Reset() {
	if r.noColor {
		return
	}
	_, _ = fmt.Fprint(r.out, "\x1b[0m")
}

func (r *Renderer) ResetFg() {
	if r.noColor {
		return
	}
	_, _ = fmt.Fprint(r.out, "\x1b[39m")
}

func (r *Renderer) ResetBg() {
	if r.noColor {
		return
	}
	_, _ = fmt.Fprint(r.out, "\x1b[49m")
}

func (r *Renderer) ResetScreenBg() {
	if r.noColor {
		_, _ = fmt.Fprint(r.out, "\x1b[2J\x1b[H")
		return
	}
	_, _ = fmt.Fprint(r.out, "\x1b[49m\x1b[2J\x1b[H")
}

func (r *Renderer) ClearScreen() {
	_, _ = fmt.Fprint(r.out, "\x1b[2J\x1b[H")
}

func (r *Renderer) ClearScreenWithBg(c Color) {
	if !r.noColor {
		r.SetBg(c)
	}
	r.ClearScreen()
}

func (r *Renderer) PrintLine(line string, fg Color, bg Color) {
	if r.noColor {
		_, _ = fmt.Fprintln(r.out, line)
		return
	}

	r.SetBg(bg)
	r.SetFg(fg)
	_, _ = fmt.Fprintln(r.out, line)
	r.Reset()
}

func (r *Renderer) PrintANSI256Table() {
	for i := 0; i < 256; i++ {
		if r.noColor {
			red, green, blue := ansi256ToRGB(i)
			_, _ = fmt.Fprintf(r.out, "   %3d  rgb(%3d,%3d,%3d)   38;5;%d\n", i, red, green, blue, i)
			continue
		}

		_, _ = fmt.Fprintln(r.out, ANSI256Sample(i))
	}
}

func SetOutput(w io.Writer) {
	if w == nil {
		w = os.Stdout
	}

	defaultRendererMu.Lock()
	defer defaultRendererMu.Unlock()

	defaultRenderer.SetOutput(w)
}

func ResetOutput() {
	SetOutput(os.Stdout)
}

func SetNoColor(noColor bool) {
	defaultRendererMu.Lock()
	defer defaultRendererMu.Unlock()

	defaultRenderer.SetNoColor(noColor)
}

func DisableColor() {
	SetNoColor(true)
}

func EnableColor() {
	SetNoColor(false)
}

func DefaultRenderer() *Renderer {
	defaultRendererMu.RLock()
	defer defaultRendererMu.RUnlock()

	return defaultRenderer
}
