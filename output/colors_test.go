package output

import (
	"testing"
)

func TestColors(t *testing.T) {
	RedText("Red message")
	RedTextB("Red bold message")
	GreenText("Green message")
	GreenTextB("Green bold message")
	BlueText("Blue message")
	BlueTextB("Blue bold message")
	YellowText("Yellow message")
	YellowTextB("Yellow bold message")
	CyanText("Cyan message")
	CyanTextB("Cyan bold message")
	WhiteText("White message")
	WhiteTextB("White bold message")
	MagentaText("Magenta message")
	MagentaTextB("Magenta bold message")
}
