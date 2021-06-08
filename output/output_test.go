package output

import (
	log "github.com/sirupsen/logrus"
	"testing"
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
	Warn("TestOutput", "Warning message")
	Warn("", "Warning message without function name")
	WarningLog("TestOutput", "Warning message")
	WarningLog("", "Warning message without function name")
	Info("TestOutput", "Info message")
	Info("", "Info message without function name")
	InfoLog("TestOutput", "Info message")
	InfoLog("", "Info message without function name")
}