package main

import (
	"os"
	"slices"

	kserver "Ktool/Kserver"
	"Ktool/ui"

	"github.com/charmbracelet/log"
)

var (
	mode      string
	noRepeat  = false
	repeat    = true
	firstPass = true
	args      = os.Args
)

func main() {

	if !firstPass {
		if len(os.Args) > 1 {
			args = slices.Delete(args, 1, len(args))
		}
	}

	if len(os.Args) == 1 {
		modeSelect()
	} else {
		switch os.Args[1] {
		case "-help":
			help()
		case "-notes":
			notes()
		case "-currency", "-currencyconvert":
			CurrencyConvert(noRepeat)
		case "-ksorter":
			ksorter_integration()
		case "-ls", "-listfromdirectory":
			listFromDirectory()
		case "-funny":
			funny(noRepeat)
		case "-pricer":
			pricer(noRepeat)
		case "-fibonacci":
			fibonacciLuncher(noRepeat)
		case "-kserver":
			kserver.StartKserver()
		// case "-time", "-timezoneconverter":
		// 	timeZoneConvert()
		default:
			log.Warn("Invalid argument provided, opening normally...")
			modeSelect()
		}
	}

	if firstPass {
		firstPass = false
	}

}

func modeSelect() {

	ui.ModePrompt.Value(&mode).Run()

	_ = luncher(mode) // returns appMode | not needed rn

}

func test() { // used only for new feature testing

}
