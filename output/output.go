package output

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(new(TextFormatter))
}

func printMessage(level, function, message string) {
	var t string
	if len(function) != 0 {
		t = fmt.Sprintf("(%s) %s", function, message)
	} else {
		t = fmt.Sprintf("%s", message)
	}

	switch level {
	case "ERROR":
		fmt.Printf("[%s] %s\n", Red+level+Reset, t)
	case "WARNING":
		fmt.Printf("[%s] %s\n", Yellow+level+Reset, t)
	case "INFO":
		fmt.Printf("[%s] %s\n", Green+level+Reset, t)
	case "DEBUG":
		fmt.Printf("[%s] %s\n", Gray+level+Reset, t)
	case "TRACE":
		fmt.Printf("[%s] %s\n", DarkGray+level+Reset, t)
	}
}

func Error(function string, message string) {
	printMessage("ERROR", function, message)
}


func Debug(function string, message string) {
	printMessage("DEBUG", function, message)
}

func Trace(function string, message string) {
	printMessage("TRACE", function, message)
}

func Warning(function string, message string) {
	printMessage("WARNING", function, message)
}

func Warn(function string, message string) {
	printMessage("WARNING", function, message)
}

func Info(function string, message string) {
	printMessage("INFO", function, message)
}



func printLog(level, function, message string) {
	var t string
	if len(function) != 0 {
		t = fmt.Sprintf("(%s) %s", function, message)
	} else {
		t = fmt.Sprintf("%s", message)
	}

	switch level {
	case "ERROR":
		log.Error(fmt.Sprintf("%s", t))
	case "WARNING":
		log.Warn(fmt.Sprintf("%s", t))
	case "INFO":
		log.Info(fmt.Sprintf("%s", t))
	case "DEBUG":
		log.Debug(fmt.Sprintf("%s", t))
	case "TRACE":
		log.Trace(fmt.Sprintf("%s", t))
	}
}

func ErrorLog(function string, message string) {
	printLog("ERROR", function, message)
}

func WarningLog(function string, message string) {
	printLog("WARNING", function, message)
}

func WarnLog(function string, message string) {
	printLog("WARNING", function, message)
}

func InfoLog(function string, message string) {
	printLog("INFO", function, message)
}

func DebugLog(function string, message string) {
	printLog("DEBUG", function, message)
}

func TraceLog(function string, message string) {
	printLog("TRACE", function, message)
}

// ErrorExit shows error and exit with error code = 1
func ErrorExit(function string, message string) {
	Error(function, message)
	os.Exit(1)
}

// CheckErrorAndExit checks any error and in case of any shows the message using the Error func and then exit,
// otherwise nothing happens.
func CheckErrorAndExit(function string, message string, err error) {
	if err != nil {
		ErrorExit(function, message)
	}
}
