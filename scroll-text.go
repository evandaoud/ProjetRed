package red

import (
	"time"

	"github.com/fatih/color"
)

// ////////////////////////////////////////////SCROLL LE TEXTE//////////////////////////////////////
func SpeedMsg(message string, speed int, colorName string) {
	defaultColor := color.New(color.FgWhite)
	Green := color.New(color.FgGreen)
	Red := color.New(color.FgRed)
	Blue := color.New(color.FgBlue)
	Cyan := color.New(color.FgCyan)
	Yellow := color.New(color.FgYellow)

	var selectedColor *color.Color

	switch colorName {
	case "green":
		selectedColor = Green
	case "red":
		selectedColor = Red
	case "blue":
		selectedColor = Blue
	case "cyan":
		selectedColor = Cyan
	case "yellow":
		selectedColor = Yellow
	default:
		selectedColor = defaultColor
	}

	for _, char := range message {
		selectedColor.Print(string(char))
		time.Sleep(time.Duration(10) * time.Millisecond)
	}
}
