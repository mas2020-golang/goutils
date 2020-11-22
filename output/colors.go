package output

import (
	"fmt"
	"github.com/jedib0t/go-pretty/text"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(new(TextFormatter))
}

func Error(function string, message string) {
	log.Error(fmt.Sprintf("(%s) %s", function, message))
	fmt.Println(text.Colors{text.FgRed}.Sprintf("Houston, help: %s", message))
}

func ErrorExit(function string, message string) {
	Error(function, message)
	os.Exit(1)
}

func CheckError(function string, message string, err error) {
	if err != nil {
		ErrorExit(function, message)
	}
}

func Debug(function string, message string) {
	log.Debug(fmt.Sprintf("%s - %s", function, message))
}

func Trace(function string, message string) {
	log.Trace(fmt.Sprintf("%s - %s", function, message))
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
