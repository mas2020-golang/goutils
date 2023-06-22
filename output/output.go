package output

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
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
	case "│ Error:":
		fmt.Printf("%s %s\n", RED+BOLD+level+Reset, t)
	case "│ Warning:":
		fmt.Printf("%s %s\n", YELLOW+BOLD+level+Reset, t)
	case "│ Warn:":
		fmt.Printf("%s %s\n", YELLOW+BOLD+level+Reset, t)	
	case "│ Info:":
		fmt.Printf("%s %s\n", GREEN+BOLD+level+Reset, t)
	case "Debug:":
		fmt.Printf("%s %s\n", GRAY+BOLD+level+Reset, t)
	case "Trace:":
		fmt.Printf("%s %s\n", DARK_GRAY+BOLD+level+Reset, t)
	}
}

func Error(function string, message string) {
	printMessage("│ Error:", function, message)
}

func Debug(function string, message string) {
	printMessage("Debug:", function, message)
}

func Trace(function string, message string) {
	printMessage("Trace:", function, message)
}

func Warning(function string, message string) {
	printMessage("│ Warning:", function, message)
}

func Warn(function string, message string) {
	printMessage("│ Warn:", function, message)
}

func WarnS(message string) string {
	return fmt.Sprintf("%s %s", YELLOW+BOLD+"│ Warn:"+Reset, message)
}

func Info(function string, message string) {
	printMessage("│ Info:", function, message)
}

func OK(message string) {
	fmt.Printf("%s%s%s\n", GREEN, message, Reset)
}

func InfoBox(message string) {
	fmt.Printf("%s--%s %s\n", YELLOW, Reset, message)
}

// Activity prints to the standard output the common pattern for an activity
func Activity(message string) {
	fmt.Printf("%s%s%s\n", BOLD, message, Reset)
}

// SubActivity prints to the standard output the common pattern for a sub activity
func SubActivity(message string) {
	fmt.Printf("%s %s\n", Cmd, message)
}

// Success prints to the standard output the common pattern for a successful message
func Success(message string) {
	fmt.Printf("%s %s\n", OkFlag, message)
}

// SuccessBold prints to the standard output the common pattern for a successful bold message
func SuccessBold(message string) {
	fmt.Printf("%s %s\n", OkFlag, fmt.Sprintf(BOLD+message+Reset))
}

func printLog(level, function, message string) {
	var t string
	if len(function) != 0 {
		t = fmt.Sprintf("(%s) %s", function, message)
	} else {
		t = fmt.Sprintf("%s", message)
	}
	switch level {
	case "Error":
		log.Error(fmt.Sprintf("%s", t))
	case "Warning":
		log.Warn(fmt.Sprintf("%s", t))
	case "Info":
		log.Info(fmt.Sprintf("%s", t))
	case "Debug":
		log.Debug(fmt.Sprintf("%s", t))
	case "Trace":
		log.Trace(fmt.Sprintf("%s", t))
	}
}

func ErrorLog(function string, message string) {
	printLog("Error", function, message)
}

func WarningLog(function string, message string) {
	printLog("Warning", function, message)
}

func WarnLog(function string, message string) {
	printLog("Warning", function, message)
}

func InfoLog(function string, message string) {
	printLog("Info", function, message)
}

func DebugLog(function string, message string) {
	printLog("Debug", function, message)
}

func TraceLog(function string, message string) {
	printLog("Trace", function, message)
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
		ErrorExit(function, message+" "+err.Error())
	}
}

// CheckErrorAndExit checks any error and in case of any shows the message using the ErrorLog func and then exit,
// otherwise nothing happens.
func CheckErrorAndExitLog(function string, message string, err error) {
	if err != nil {
		ErrorLog(function, message+" "+err.Error())
		os.Exit(1)
	}
}
