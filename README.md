![Logo](logo.png) 

# AMS TIC

```txt
┌─┴─┴─┴─┴─┴─┴─┴─┴─┐
🭲○  ♦ AMS TIC ♦   🭲
└─┬─┬─┬─┬─┬─┬─┬─┬─┘
```

**TIC** is a Go library for retro-inspired terminal styling. It provides a complete color palette inspired by the Commodore 64/VIC-II and Atari 800, ANSI 16/256-color output, a fluent style builder, a testable renderer API, themes, box components, and text utility helpers.

## Installation

```bash
go get github.com/ams-soft/tic
```

---

## Colors

TIC defines 22 named colors as `Color` constants. All style and renderer functions accept these constants.

| Constant          | Description                          |
|-------------------|--------------------------------------|
| `ColorBlack`      | Black                                |
| `ColorWhite`      | Bright white                         |
| `ColorRed`        | Standard red                         |
| `ColorCyan`       | Cyan                                 |
| `ColorPurple`     | Purple (ANSI 256: 93)                |
| `ColorGreen`      | Standard green                       |
| `ColorBlue`       | Standard blue                        |
| `ColorYellow`     | Yellow/amber                         |
| `ColorOrange`     | Orange (ANSI 256: 208)               |
| `ColorBrown`      | Brown (ANSI 256: 130)                |
| `ColorLightRed`   | Bright red                           |
| `ColorDarkGray`   | Dark gray                            |
| `ColorMediumGray` | Medium gray                          |
| `ColorLightGreen` | Bright green                         |
| `ColorLightBlue`  | Bright blue                          |
| `ColorLightGray`  | Light gray                           |
| `ColorLightPurple`| Light purple (ANSI 256: 135)         |
| `ColorDarkPurple` | Dark purple (ANSI 256: 54)           |
| `ColorC64BG`      | Commodore 64 background (ANSI 256: 18) |
| `ColorC64FG`      | Commodore 64 foreground (ANSI 256: 63) |
| `ColorA800BG`     | Atari 800 background (ANSI 256: 23)  |
| `ColorA800FG`     | Atari 800 foreground (ANSI 256: 43)  |

### Color methods

```go
c := tic.ColorCyan
c.Valid()   // true — c is a built-in TIC color
c.String()  // "Cyan"
```

Invalid `Color` values silently fall back to the terminal default (ANSI reset code).

### ColorList

`ColorList` is a `[]ColorEntry` that exposes the full palette in stable display order. Each entry has `Color` and `Name` fields.

```go
for _, entry := range tic.ColorList {
    fmt.Printf("%s\n", tic.ColorText(entry.Name, entry.Color))
}
```

---

## Style helpers

Stateless functions that wrap text in ANSI sequences and reset. All respect the global `NoColor` flag.

```go
tic.Bold("text")       // bold
tic.Dim("text")        // dim/faint
tic.Italic("text")     // italic
tic.Underline("text")  // underline
tic.Blink("text")      // blink
tic.Reverse("text")    // reverse video (swap fg/bg)
tic.Strike("text")     // strikethrough

tic.ColorText("text", tic.ColorCyan)                   // foreground color
tic.ColorBg("text", tic.ColorYellow)                   // background color
tic.WithColors("text", tic.ColorWhite, tic.ColorBlue)  // fg + bg
tic.Colorize("text", tic.ColorWhite, tic.ColorBlue)    // alias for WithColors

tic.ResetStyle()  // returns "\x1b[0m" string (or "" when NoColor)
```

---

## Style Builder

`StyleBuilder` is a fluent API for composing multiple styles into a single call. Useful when you need to combine colors and attributes without nesting function calls.

### Construction

```go
s := tic.Style()     // preferred alias
s := tic.NewStyle()  // identical
```

### Chaining methods

| Method          | Effect                     |
|-----------------|----------------------------|
| `.Fg(c Color)`  | Set foreground color       |
| `.Bg(c Color)`  | Set background color       |
| `.Bold()`       | Bold                       |
| `.Dim()`        | Dim/faint                  |
| `.Italic()`     | Italic                     |
| `.Underline()`  | Underline                  |
| `.Blink()`      | Blink                      |
| `.Reverse()`    | Reverse video              |
| `.Strike()`     | Strikethrough              |

### Rendering

```go
text := tic.Style().
    Fg(tic.ColorLightGreen).
    Bg(tic.ColorBlack).
    Bold().
    Sprint("SUCCESS: operation completed")

fmt.Println(text)

// With newline
fmt.Print(tic.Style().Fg(tic.ColorRed).Sprintln("ERROR"))
```

`Sprint` returns the styled string. `Sprintln` appends `\n`. Both return plain text when `NoColor` is active.

### Inspecting codes

```go
codes := tic.Style().Fg(tic.ColorCyan).Bold().Codes()
// ["36", "1"]
```

---

## Renderer

`Renderer` writes output to any `io.Writer`. Prefer it over package-level functions in applications, tests, and libraries.

### Creating a renderer

```go
r := tic.NewRenderer(os.Stdout)
r := tic.NewRenderer(os.Stdout, tic.WithNoColor(true))
r := tic.NewRenderer(&buf)          // bytes.Buffer for tests
r := tic.NewRenderer(nil)           // nil is safe: writes are discarded
```

### Renderer options

| Option                      | Effect                                   |
|-----------------------------|------------------------------------------|
| `WithNoColor(bool)`         | Disable/enable ANSI output for this renderer |
| `WithOutput(io.Writer)`     | Set output writer via option (nil → discard) |

### Output control

```go
r.SetOutput(w io.Writer)  // change writer after construction
r.Output() io.Writer      // retrieve current writer
r.SetNoColor(bool)        // toggle no-color mode at runtime
r.NoColor() bool          // query current no-color state
```

### Writing text

```go
r.Print(a ...any)
r.Println(a ...any)
r.Printf(format string, a ...any)
```

### Color control

```go
r.SetFg(c Color)   // set foreground and keep it for subsequent writes
r.SetBg(c Color)   // set background and keep it
r.Reset()          // reset all attributes (\x1b[0m)
r.ResetFg()        // reset foreground only (\x1b[39m)
r.ResetBg()        // reset background only (\x1b[49m)
```

### Screen control

```go
r.ClearScreen()            // clear screen and move cursor to home
r.ClearScreenWithBg(c)     // set background color then clear screen
r.ResetScreenBg()          // reset background then clear screen
```

### Convenience

```go
r.PrintLine(line string, fg, bg Color)  // set colors, print line, reset
r.PrintANSI256Table()                   // print all 256 ANSI color samples
```

### Example: testable output

```go
var buf bytes.Buffer
r := tic.NewRenderer(&buf)
r.PrintLine("OK", tic.ColorLightGreen, tic.ColorBlack)
// buf.String() contains the styled output
```

---

## Package-level functions

Thin delegates to the internal `DefaultRenderer` (writes to `os.Stdout`). Kept for V1 compatibility and quick scripts.

```go
// Color control
tic.SetFg(c Color)
tic.SetBg(c Color)
tic.Reset()
tic.ResetAll()    // alias for Reset
tic.ResetFg()
tic.ResetBg()
tic.ResetScreenBg()

// Screen
tic.ClearScreen()
tic.ClearScreenWithBg(c Color)

// Output
tic.PrintLine(line string, fg, bg Color)
tic.PrintANSI256Table()

// Default renderer output target
tic.SetOutput(w io.Writer)  // redirect default renderer (nil → os.Stdout)
tic.ResetOutput()           // restore to os.Stdout

// No-color
tic.DisableColor()
tic.EnableColor()
tic.SetNoColor(bool)

// Access the default renderer directly
r := tic.DefaultRenderer()
```

---

## Themes

A `Theme` bundles a palette of semantic colors. Three built-in themes are provided.

### Built-in themes

| Constructor          | Name       | Inspired by         |
|----------------------|------------|---------------------|
| `ThemeC64()`         | C64        | Commodore 64/VIC-II |
| `ThemeAtari800()`    | Atari800   | Atari 800           |
| `ThemeCRT()`         | CRT        | Green phosphor CRT  |

### Theme fields

```go
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
```

### Theme methods

All methods return a styled string (bold where noted):

```go
theme := tic.ThemeC64()

fmt.Println(theme.Title("AMS BANKING CORE"))           // fg+bg, bold
fmt.Println(theme.SuccessText("READY."))               // success color, bold
fmt.Println(theme.WarningText("LOW MEMORY"))           // warning color, bold
fmt.Println(theme.ErrorText("DEVICE NOT PRESENT"))     // error color, bold
fmt.Println(theme.InfoText("64K RAM SYSTEM  38911 BASIC BYTES FREE"))  // info color
```

### Custom theme

```go
myTheme := tic.Theme{
    Name:       "Custom",
    Foreground: tic.ColorLightGray,
    Background: tic.ColorBlack,
    Accent:     tic.ColorCyan,
    Success:    tic.ColorLightGreen,
    Warning:    tic.ColorYellow,
    Error:      tic.ColorLightRed,
    Info:       tic.ColorLightBlue,
}
```

---

## Components

### Box

Renders a bordered box with a title and a multi-line body.

```go
output := tic.Box(title string, body string, style BoxStyle) string
```

The width is computed from the longest visible line (ANSI-aware). Styled strings in `title` or `body` work correctly.

```go
fmt.Println(tic.Box("STATUS", "SYSTEM ONLINE\nREADY.", tic.BoxRetro))
```

Output:
```
┌───────────────┐
│ STATUS        │
│───────────────│
│ SYSTEM ONLINE │
│ READY.        │
└───────────────┘
```

### Header

Shorthand for a `Box` with `"AMS TIC"` as the body.

```go
fmt.Println(tic.Header("AMS BANKING CORE"))
```

### Box styles

Three built-in `BoxStyle` values:

| Variable    | Corners / lines         |
|-------------|-------------------------|
| `BoxSingle` | `┌ ┐ └ ┘ ─ │`          |
| `BoxDouble` | `╔ ╗ ╚ ╝ ═ ║`          |
| `BoxRetro`  | same as `BoxSingle`     |

Custom box style:

```go
custom := tic.BoxStyle{
    TopLeft: "+", TopRight: "+",
    BottomLeft: "+", BottomRight: "+",
    Horizontal: "-", Vertical: "|",
}
fmt.Println(tic.Box("TITLE", "body", custom))
```

---

## Text utilities

### ANSI stripping

```go
clean := tic.StripANSI(s string) string  // remove all ANSI/VT escape sequences
tic.HasANSI(s string) bool               // true if s contains any ANSI sequence
```

### Visible length

```go
tic.VisibleLen(s string) int  // rune count excluding ANSI sequences
```

Always use `VisibleLen` instead of `len` when measuring styled text, e.g. for padding or alignment inside components.

### Padding and alignment

```go
// Pad s to width with trailing spaces (ANSI-aware)
tic.PadRightVisible(s string, width int) string

// Center s within width with leading spaces (ANSI-aware)
tic.CenterVisible(s string, width int) string
```

### Terminal size

```go
width, height, err := tic.TerminalSize(fd int)
// Wraps golang.org/x/term.GetSize
// Typical usage: tic.TerminalSize(int(os.Stdout.Fd()))
```

---

## ANSI 256 utilities

```go
// Returns a printable sample line for a given ANSI 256 color index (0–255)
line := tic.ANSI256Sample(i int) string
// Example output: "♦  42  rgb(  0,215,135)   38;5;42"

// Print all 256 colors to the default renderer
tic.PrintANSI256Table()

// Or via a specific renderer
r.PrintANSI256Table()
```

---

## No-color mode

TIC respects a global no-color flag that strips all ANSI output. Useful for CI, log files, or piped output.

```go
// Global (affects package-level functions and DefaultRenderer)
tic.DisableColor()
fmt.Println(tic.ColorText("plain text", tic.ColorRed))  // prints: plain text
tic.EnableColor()

// Per-renderer
r := tic.NewRenderer(os.Stdout, tic.WithNoColor(true))

// Toggle at runtime
r.SetNoColor(true)
```

`StyleBuilder.Sprint` also checks the `DefaultRenderer` no-color flag automatically.

---

## Development

```bash
make dev    # run demo: go run ./cmd/tic-demo/main.go
make test   # go test ./...
make fmt    # gofmt -w .
make vet    # go vet ./...
```

Run a single test:

```bash
go test -run TestName ./...
```

---

## License

MIT © AMS Tecnologia e Serviços Financeiros LTDA, AKA AMS Soft
