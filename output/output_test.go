package output

import (
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	log.SetLevel(log.TraceLevel)
	Error("TestOutput", "Error message")
	Error("", "Error message without function name")
	ErrorLog("TestOutput", "Error message")
	ErrorLog("", "Error message without function name")
	Debug("TestOutput", "Debug message")
	Debug("", "Debug message without function name")
	DebugLog("TestOutput", "Debug message")
	DebugLog("", "Debug message without function name")
	Trace("TestOutput", "Trace message")
	Trace("", "Trace message without function name")
	TraceLog("TestOutput", "Trace message")
	TraceLog("", "Trace message without function name")
	Warn("TestOutput", "Warn message")
	Warn("", "Warn message without function name")
	Warning("", "Warning message without function name")
	OK("This is an OK message")
	fmt.Printf("%s - %s\n", "WarnS output", WarnS("WarnS message without function name"))
	WarningLog("TestOutput", "Warning message")
	WarningLog("", "Warning message without function name")
	Info("TestOutput", "Info message")
	Info("", "Info message without function name")
	InfoBox("this is a new box info")
	InfoLog("TestOutput", "Info message")
	InfoLog("", "Info message without function name")
	Activity("This is a new activity starting...")
	SubActivity("this is a new sub activity starting...")
	ActivityBox("this is a new activity box message")
	fmt.Printf("%s ok flag, %s error flag, %s error cmd\n", OkFlag, ErrorFlag, ErrorCmd)

	//err := fmt.Errorf("this is an error to trace")
	//CheckErrorAndExitLog("", "Error message:", err)

	// Use color in the log output
	log.SetFormatter(new(TextColorFormatter))
	ErrorLog("", "Error message without function name")
	InfoLog("", "Info message without function name")
	WarningLog("", "Warning message without function name")
	TraceLog("TestOutput", "Trace message and color")
}
