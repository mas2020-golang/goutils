package output

import (
	"testing"
	"github.com/jedib0t/go-pretty/text"
	log "github.com/sirupsen/logrus"
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
	White("White message")
	WhiteBold("White Bold message")
	Message("Generic message", text.Colors{text.Bold, text.BgGreen})
	//ErrorExit("TestColors", "Error message bye bye...")
}