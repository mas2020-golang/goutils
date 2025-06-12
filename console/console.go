package console

import (
	"fmt"
	"strings"
)

// ANSI color codes
const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
	Dim   = "\033[2m"

	// Colors
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Gray    = "\033[90m"

	// Background colors
	BgRed    = "\033[41m"
	BgGreen  = "\033[42m"
	BgYellow = "\033[43m"
	BgBlue   = "\033[44m"
	BgCyan   = "\033[46m"
)

// Icons using Unicode symbols
const (
	IconInfo        = Yellow+"--"+Reset
	IconWarn        = "❕"
	IconOK          = "👍"
	IconError       = "❌"
	IconActivity    = Cyan+">>"+Reset
	IconSubActivity = " ↳"
	IconBox         = "📦"
	IconDone         = "🚀"
)

// Message types
type MessageType int

const (
	TypeInfo MessageType = iota
	TypeWarn
	TypeOK
	TypeError
	TypeActivity
	TypeSubActivity
	TypeDone
)

// Config holds the console configuration
type Config struct {
	UseColors bool
	UseIcons  bool
	Prefix    string
}

// Default configuration
var defaultConfig = Config{
	UseColors: true,
	UseIcons:  true,
	Prefix:    "",
}

// SetConfig allows customization of the console output
func SetConfig(config Config) {
	defaultConfig = config
}

// colorize applies color to text if colors are enabled
func colorize(text, color string) string {
	if !defaultConfig.UseColors {
		return text
	}
	return color + text + Reset
}

// formatMessage formats a message with icon and styling
func formatMessage(msgType MessageType, message string, isBox bool) string {
	//var icon, color, prefix string
	var icon, color string

	switch msgType {
	case TypeInfo:
		icon = IconInfo
		color = Blue
		// prefix = ""
	case TypeWarn:
		icon = IconWarn
		color = Yellow
		// prefix = ""
	case TypeOK:
		icon = IconOK
		color = Green
		// prefix = ""
	case TypeError:
		icon = IconError
		color = Red
		// prefix = ""
	case TypeActivity:
		icon = IconActivity
		color = Cyan
		// prefix = ""
	case TypeSubActivity:
		icon = IconSubActivity
		color = Gray
		// prefix = ""
	case TypeDone:
		icon = IconDone
		color = Gray
		// prefix = ""
	}

	// Build the formatted message
	var parts []string

	// Add custom prefix if configured
	if defaultConfig.Prefix != "" {
		parts = append(parts, colorize(fmt.Sprintf("[%s]", defaultConfig.Prefix), Gray))
	}

	// Add icon if enabled
	if defaultConfig.UseIcons {
		if isBox {
			parts = append(parts, IconBox+" "+icon)
		} else {
			parts = append(parts, icon)
		}
	}

	// Add colored prefix
	//parts = append(parts, colorize(Bold+prefix+Reset, color))

	// Add the message
	parts = append(parts, colorize(message, color))

	return strings.Join(parts, " ")
}

// formatBox creates a box around the message
func formatBox(message string, color string) string {
	if !defaultConfig.UseColors {
		color = ""
	}

	lines := strings.Split(message, "\n")
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	// Minimum box width
	if maxLen < 20 {
		maxLen = 20
	}

	boxWidth := maxLen + 4
	topBottom := strings.Repeat("─", boxWidth-2)
	
	var result strings.Builder	
	result.WriteString(color + "┌" + topBottom + "┐" + Reset + "\n")

	for _, line := range lines {
		padding := boxWidth - len(line)
		//fmt.Println("boxWidth", boxWidth, "len line", len(line), "padding", padding)
		result.WriteString(color + "│ " + Reset + line + strings.Repeat(" ", padding) + color + " │" + Reset + "\n")
	}

	result.WriteString(color + "└" + topBottom + "┘" + Reset)
	return result.String()
}

// Info displays an info message
func Info(message string) {
	fmt.Println(formatMessage(TypeInfo, message, false))
}

// InfoBox displays an info message in a box
func InfoBox(message string) {
	boxed := formatBox(message, Blue)
	fmt.Println(boxed)
}

// Warn displays a warning message
func Warn(message string, colorize bool) {
	if colorize {
		message = Yellow+message+Reset
	}
	fmt.Println(formatMessage(TypeWarn, message, false))
}

// OK displays a success message
func OK(message string) {
	fmt.Println(formatMessage(TypeOK, message, false))
}

// Error displays an error message. Giving colorize ensures the color message,
// despite the Config settings.
func Error(message string, colorize bool) {
	if colorize {
		message = Red+message+Reset
	}
	fmt.Println(formatMessage(TypeError, message, false))
}

// Error displays an error message
func ActivityDone(message string) {
	fmt.Println(formatMessage(TypeDone, message, false))
}

// Activity displays an activity message
func Activity(message string) {
	fmt.Println(formatMessage(TypeActivity, message, false))
}

// SubActivity displays a sub-activity message
func SubActivity(message string) {
	fmt.Println(formatMessage(TypeSubActivity, message, false))
}

// ActivityBox displays an activity message in a box
func ActivityBox(message string) {
	formatted := formatMessage(TypeActivity, message, true)
	boxed := formatBox(formatted, Cyan)
	fmt.Println(boxed)
}

// Spinner represents a loading spinner
type Spinner struct {
	message string
	active  bool
	frames  []string
	current int
}

// NewSpinner creates a new spinner for activities
func NewSpinner(message string) *Spinner {
	return &Spinner{
		message: message,
		frames:  []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	}
}

// Start begins the spinner animation
func (s *Spinner) Start() {
	s.active = true
	// Implementation would require goroutines for actual spinning
	// This is a simplified version
	fmt.Printf("%s %s %s\n",
		colorize(s.frames[0], Cyan),
		colorize("LOADING", Cyan),
		s.message)
}

// Stop ends the spinner animation
func (s *Spinner) Stop(success bool) {
	s.active = false
	if success {
		OK(s.message + " completed")
	} else {
		Error(s.message + " failed", true)
	}
}
