package output

import (
	"fmt"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	LightRed    = "\033[91m"
	Green  = "\033[32m"
	LightGreen  = "\033[92m"
	Yellow = "\033[33m"
	LightYellow = "\033[93m"
	Blue   = "\033[34m"
	LightBlue   = "\033[94m"
	Purple = "\033[35m"
	LightPurple = "\033[95m"
	Cyan   = "\033[36m"
	LightCyan   = "\033[96m"
	Gray   = "\033[37m"
	White  = "\033[97m"
	Default = "\033[39m"
	DarkGray = "\033[90m"
	Bold = "\033[1m"
)

func GreenText(message string) {
	fmt.Printf("%s%s%s\n",Green,message,Reset)
}

func GreenTextB(message string) {
	fmt.Printf("%s%s%s%s\n",Green,Bold,message,Reset)
}

func BlueText(message string) {
	fmt.Printf("%s%s%s\n",Blue,message,Reset)
}

func BlueTextB(message string) {
	fmt.Printf("%s%s%s%s\n",Blue,Bold,message,Reset)
}

func YellowText(message string) {
	fmt.Printf("%s%s%s\n",Yellow,message,Reset)
}

func YellowTextB(message string) {
	fmt.Printf("%s%s%s%s\n",Yellow,Bold,message,Reset)
}

func CyanText(message string) {
	fmt.Printf("%s%s%s\n",Cyan,message,Reset)
}

func CyanTextB(message string) {
	fmt.Printf("%s%s%s%s\n",Cyan,Bold,message,Reset)
}

func WhiteText(message string) {
	fmt.Printf("%s%s%s\n",White,message,Reset)
}

func WhiteTextB(message string) {
	fmt.Printf("%s%s%s%s\n",White,Bold,message,Reset)
}

func MagentaText(message string) {
	fmt.Printf("%s%s%s\n",Purple,message,Reset)
}

func MagentaTextB(message string) {
	fmt.Printf("%s%s%s%s\n",Purple,Bold,message,Reset)
}

func RedText(message string) {
	fmt.Printf("%s%s%s\n",Red,message,Reset)
}

func RedTextB(message string) {
	fmt.Printf("%s%s%s%s\n",Red,Bold,message,Reset)
}