package console

import (
	"testing"
)

func TestOutput(t *testing.T) {
	// Demonstrate configuration options
	SetConfig(Config{
		UseColors: false,
		UseIcons:  true,
		Prefix:    "",
	})

	// Basic usage with default configuration
	Info("Info message without function name")
	InfoBox("this is a new box info")
	Warn("Warn message without function name", true)
	OK("This is an OK message")
	Activity("This is a new activity starting...")
	SubActivity("this is a new sub activity starting...")
	ActivityBox("this is a new activity box message")
	Error("this is a new error message (with color)", true)
	ActivityDone("this activity is done")

	// Demonstrate configuration options
	SetConfig(Config{
		UseColors: true,
		UseIcons:  true,
		Prefix:    "MyApp",
	})

	Info("This message includes the app prefix")

	// Example of disabling colors for CI/CD environments
	SetConfig(Config{
		UseColors: true,
		UseIcons:  true,
		Prefix:    "",
	})

	Warn("This warning has no colors and keeps icons", false)

	// Demonstrate spinner (simplified version)
	// spinner := NewSpinner("Processing data")
	// spinner.Start()
	// time.Sleep(2 * time.Second) // Simulate work
	// spinner.Stop(true)

	// Reset to default
	SetConfig(Config{
		UseColors: true,
		UseIcons:  true,
		Prefix:    "",
	})

	// Multi-line box example
	InfoBox("This is a multi-line box\nwith several lines of text\nto demonstrate the box formatting")

}
