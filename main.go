package main

//go:generate go-winres make

import (
	"log"
	"net"
	"os"

	"github.com/AlecAivazis/survey/v2"

	"github.com/mingrammer/cfmt"
)

func main() {

	getCmdArgs()

}

func getCmdArgs() {
	args := os.Args
	noRepeat := false

	if len(args) == 1 {
		modeSelect()
	} else {
		switch args[1] {
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
		// case "-time", "-timezoneconverter":
		// 	timeZoneConvert()
		default:
			cfmt.Errorln("Invalid argument provided, opening normally...")
			modeSelect()
		}
	}

}

func modeSelect() {
	mode := ""
	repeat := true
	prompt := &survey.Select{
		Message: "app mode:",
		Options: []string{
			"notes",
			"currency convert",
			"open Ksorter",
			"list from directory",
			"funny",
			"pricer",
			"fibonacci",
			"exit",
		},
	}

	survey.AskOne(prompt, &mode)

	switch mode {
	case "sys info":
		sysInfo()
	case "pricer":
		pricer(repeat)
	case "notes":
		notes()
	case "list from directory":
		listFromDirectory()
	case "messenger":
		messenger()
	case "funny":
		funny(repeat)
	case "currency convert":
		CurrencyConvert(repeat)
	case "fibonacci":
		fibonacciLuncher(repeat)
	case "time zone converter":
		timeZoneConvert()
	case "test":
		test()
	case "open Ksorter":
		ksorter_integration()
	case "exit":
		cfmt.Errorln("exiting...")
		return

	}
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func test() { // used only for new feature testing

}
