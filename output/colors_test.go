package output

import (
	"github.com/jedib0t/go-pretty/text"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestColors(t *testing.T) {
	log.SetLevel(log.TraceLevel)
	Error("TestColors", "Error message")
	Debug("TestColors", "Debug message")
	Trace("TestColors", "Trace message")
	Green("Green message")
	GreenBold("Green Bold message")
	Blue("Blue message")
	BlueBold("Blue Bold message")
	Yellow("Yellow message")
	YellowBold("Yellow Bold message")
	Cyan("Cyan message")
	CyanBold("Cyan Bold message")
	HiWhite("HiWhite message")
	HiWhiteBold("HiWhiteBold Bold message")
	HiYellow("HiYellow message")
	HiYellowBold("HiYellowBold Bold message")
	HiMagenta("HiMagenta message")
	HiMagentaBold("HiMagentaBold Bold message")
	Message("Generic message", text.Colors{text.Bold, text.BgGreen})
	//ErrorExit("TestColors", "Error message bye bye...")
}
