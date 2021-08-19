package output

import (
	"testing"
)

func TestColors(t *testing.T) {
	Red("RED message")
	RedBold("RED bold message")
	Green("GREEN message")
	GreenBold("GREEN bold message")
	Blue("BLUE message")
	BlueBold("BLUE bold message")
	Yellow("YELLOW message")
	YellowBold("YELLOW bold message")
	Cyan("CYAN message")
	CyanBold("CYAN bold message")
	White("WHITE message")
	WhiteBold("WHITE bold message")
	Magenta("Magenta message")
	MagentaBold("Magenta bold message")
}
