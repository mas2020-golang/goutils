package output

import (
	"fmt"
	"os"
	"time"
	"github.com/jedib0t/go-pretty/text"
	log "github.com/sirupsen/logrus"
)

func Error(function string, message string) {
	log.Error(function + " " + time.Now().Format("2006-01-02T15:04:05.000 ") + message)
	fmt.Println(text.Colors{text.FgRed}.Sprint("Houston, we have a problem: " + message))
}

func ErrorExit(function string, message string) {
	Error(function,message)
	os.Exit(1)
}

func CheckError(function string, message string, err error) {
	if err != nil {
		ErrorExit(function, message)
	}
}

func Debug(function string, message string) {
	log.Debug(function + " " + time.Now().Format("2006-01-02T15:04:05.000 ") + message)
}

func Trace(function string, message string) {
	log.Trace(function + " " + time.Now().Format("2006-01-02T15:04:05.000 ") + message)
}

func Green(message string) {
	fmt.Println(text.Colors{text.FgGreen}.Sprint(message))
}

func GreenBold(message string) {
	fmt.Println(text.Colors{text.FgGreen, text.Bold}.Sprint(message))
}

func Blue(message string) {
	fmt.Println(text.Colors{text.FgBlue}.Sprint(message))
}

func BlueBold(message string) {
	fmt.Println(text.Colors{text.FgBlue, text.Bold}.Sprint(message))
}

func Yellow(message string) {
	fmt.Println(text.Colors{text.FgYellow}.Sprint(message))
}

func YellowBold(message string) {
	fmt.Println(text.Colors{text.FgYellow, text.Bold}.Sprint(message))
}

func Cyan(message string) {
	fmt.Println(text.Colors{text.FgCyan}.Sprint(message))
}

func CyanBold(message string) {
	fmt.Println(text.Colors{text.FgCyan, text.Bold}.Sprint(message))
}

// Generic message with a custom color
func Message(message string, colors text.Colors) {
	fmt.Println(colors.Sprint(message))
}
