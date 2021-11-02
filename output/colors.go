package output

import (
	"fmt"
)

const (
	Reset        = "\033[0m"
	RED          = "\033[31m"
	LIGHT_RED    = "\033[91m"
	GREEN        = "\033[32m"
	LIGHT_GREEN  = "\033[92m"
	YELLOW       = "\033[33m"
	LIGHT_YELLOW = "\033[93m"
	BLUE         = "\033[34m"
	LIGHT_BLUE   = "\033[94m"
	PURPLE       = "\033[35m"
	LIGHT_PURPLE = "\033[95m"
	CYAN         = "\033[36m"
	LIGHT_CYAN   = "\033[96m"
	GRAY         = "\033[37m"
	WHITE        = "\033[97m"
	DEFAULT      = "\033[39m"
	DARK_GRAY    = "\033[90m"
	BOLD         = "\033[1m"
	Orange       = "\033[38;5;167m"
	Cmd          = "\033[1;92m➜\033[0m"
	SubCmd       = "\033[38;5;167m⌁\033[0m"
	OkFlag       = "\u001B[92m✔\u001B[0m"
	ErrorFlag    = "\033[91m✘\033[0m"
	ErrorCmd     = "[\033[91mERROR\033[0m]"
)

// GreenOut prints onto the standard output
func GreenOut(message string) {
	fmt.Printf("%s%s%s\n", GREEN, message, Reset)
}

// GreenOutBold prints green-bold onto the standard output
func GreenOutBold(message string) {
	fmt.Printf("%s%s%s%s\n", GREEN, BOLD, message, Reset)
}

func GreenS(message string) string {
	return fmt.Sprintf("%s%s%s", GREEN, message, Reset)
}

func GreenBoldS(message string) string {
	return fmt.Sprintf("%s%s%s%s", GREEN, BOLD, message, Reset)
}

func OrangeS(message string) string {
	return fmt.Sprintf("%s%s%s", Orange, message, Reset)
}

// Prints onto the standard output
func BlueOut(message string) {
	fmt.Printf("%s%s%s\n", BLUE, message, Reset)
}

func BlueOutBold(message string) {
	fmt.Printf("%s%s%s%s\n", BLUE, BOLD, message, Reset)
}

func BlueS(message string) string {
	return fmt.Sprintf("%s%s%s", BLUE, message, Reset)
}

func BlueBoldS(message string) string {
	return fmt.Sprintf("%s%s%s%s", BLUE, BOLD, message, Reset)
}

func YellowOut(message string) {
	fmt.Printf("%s%s%s\n", YELLOW, message, Reset)
}

func YellowOutBold(message string) {
	fmt.Printf("%s%s%s%s\n", YELLOW, BOLD, message, Reset)
}

func YellowBoldS(message string) string {
	return fmt.Sprintf("%s%s%s%s", YELLOW, BOLD, message, Reset)
}

func YellowS(message string) string {
	return fmt.Sprintf("%s%s%s", YELLOW, message, Reset)
}

func CyanOut(message string) {
	fmt.Printf("%s%s%s\n", CYAN, message, Reset)
}

func CyanOutBold(message string) {
	fmt.Printf("%s%s%s%s\n", CYAN, BOLD, message, Reset)
}

func WhiteOut(message string) {
	fmt.Printf("%s%s%s\n", WHITE, message, Reset)
}

func WhiteS(message string) string {
	return fmt.Sprintf("%s%s%s", WHITE, message, Reset)
}

func WhiteOutBold(message string) {
	fmt.Printf("%s%s%s%s\n", WHITE, BOLD, message, Reset)
}

func WhiteBoldS(message string) string {
	return fmt.Sprintf("%s%s%s%s", WHITE, BOLD, message, Reset)
}

func MagentaOut(message string) {
	fmt.Printf("%s%s%s\n", PURPLE, message, Reset)
}

func MagentaOutBold(message string) {
	fmt.Printf("%s%s%s%s\n", PURPLE, BOLD, message, Reset)
}

func RedOut(message string) {
	fmt.Printf("%s%s%s\n", RED, message, Reset)
}

func RedBoldS(message string) string {
	return fmt.Sprintf("%s%s%s%s", RED, BOLD, message, Reset)
}

func RedS(message string) string {
	return fmt.Sprintf("%s%s%s", RED, message, Reset)
}

func RedOutBold(message string) {
	fmt.Printf("%s%s%s%s\n", RED, BOLD, message, Reset)
}

func Bold(message string) {
	fmt.Printf("%s%s%s\n", BOLD, message, Reset)
}

func BoldS(message string) string {
	return fmt.Sprintf("%s%s%s", BOLD, message, Reset)
}
