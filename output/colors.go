package output

import (
	"fmt"
)

const (
	RESET        = "\033[0m"
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
)

func Green(message string) {
	fmt.Printf("%s%s%s\n", GREEN, message, RESET)
}

func GreenBold(message string) {
	fmt.Printf("%s%s%s%s\n", GREEN, BOLD, message, RESET)
}

func Blue(message string) {
	fmt.Printf("%s%s%s\n", BLUE, message, RESET)
}

func BlueBold(message string) {
	fmt.Printf("%s%s%s%s\n", BLUE, BOLD, message, RESET)
}

func Yellow(message string) {
	fmt.Printf("%s%s%s\n", YELLOW, message, RESET)
}

func YellowBold(message string) {
	fmt.Printf("%s%s%s%s\n", YELLOW, BOLD, message, RESET)
}

func Cyan(message string) {
	fmt.Printf("%s%s%s\n", CYAN, message, RESET)
}

func CyanBold(message string) {
	fmt.Printf("%s%s%s%s\n", CYAN, BOLD, message, RESET)
}

func White(message string) {
	fmt.Printf("%s%s%s\n", WHITE, message, RESET)
}

func WhiteBold(message string) {
	fmt.Printf("%s%s%s%s\n", WHITE, BOLD, message, RESET)
}

func Magenta(message string) {
	fmt.Printf("%s%s%s\n", PURPLE, message, RESET)
}

func MagentaBold(message string) {
	fmt.Printf("%s%s%s%s\n", PURPLE, BOLD, message, RESET)
}

func Red(message string) {
	fmt.Printf("%s%s%s\n", RED, message, RESET)
}

func RedBold(message string) {
	fmt.Printf("%s%s%s%s\n", RED, BOLD, message, RESET)
}
