package output

import (
	"testing"
)

func TestColors(t *testing.T) {
	RedOut("RED message")
	RedOutBold("RED bold message")
	GreenOut("GREEN message")
	GreenOutBold("GREEN bold message")
	BlueOut("BLUE message")
	BlueOutBold("BLUE bold message")
	YellowOut("YELLOW message")
	YellowOutBold("YELLOW bold message")
	CyanOut("CYAN message")
	CyanOutBold("CYAN bold message")
	WhiteOut("WHITE message")
	WhiteOutBold("WHITE bold message")
	MagentaOut("MagentaOut message")
	MagentaOutBold("MagentaOut bold message")
}
