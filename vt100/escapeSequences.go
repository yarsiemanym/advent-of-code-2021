package vt100

import (
	"fmt"
	"strings"
)

const ResetAttributes = "0"

const BrightAttribute = "1"
const DimAttribute = "2"
const UndescoreAttribute = "4"
const BlinkAttribute = "5"
const ReverseAttribute = "7"
const HiddenAttribute = "8"

const BlackForegroundAttribute = "30"
const RedForegroundAttribute = "31"
const GreenForegroundAttribute = "32"
const YellowForegroundAttribute = "33"
const BlueForegroundAttribute = "34"
const MagentaForegroundAttribute = "35"
const CyanForegroundAttribute = "36"
const WhiteForegroundAttribute = "37"

const BlackBackgroundAttribute = "40"
const RedBackgroundAttribute = "41"
const GreenBackgroundAttribute = "42"
const YellowBackgroundAttribute = "43"
const BlueBackgroundAttribute = "44"
const MagentaBackgroundAttribute = "45"
const CyanBackgroundAttribute = "46"
const WhiteBackgroundAttribute = "47"

func Print(text string, attributes ...string) {
	attributeString := strings.Join(attributes, ";")
	fmt.Printf("\x1b[%sm%s\x1b[0m", attributeString, text)
}

func Println(text string, attributes ...string) {
	text += "\n"
	Print(text, attributes...)
}
