package main

import (
	kserver "Ktool/Kserver"

	"github.com/charmbracelet/log"
)

type appMode int

const (
	helpR               appMode = iota
	notesR              appMode = iota
	currencyR           appMode = iota
	ksorterR            appMode = iota
	lsR                 appMode = iota
	funnyR              appMode = iota
	pricerR             appMode = iota
	fibonacciR          appMode = iota
	kserverR            appMode = iota
	exitR               appMode = iota
	depracatedComponent appMode = iota
)

func luncher(mode string) appMode {

	switch mode {
	case "sys info":
		sysInfo()
		return depracatedComponent
	case "pricer":
		pricer(repeat)
		return pricerR
	case "notes":
		notes()
		return notesR
	case "list from directory":
		listFromDirectory()
		return lsR
	// case "messenger":
	// 	messenger()
	case "funny":
		funny(repeat)
		return funnyR
	case "currency convert":
		CurrencyConvert(repeat)
		return currencyR
	case "fibonacci":
		fibonacciLuncher(repeat)
		return fibonacciR
	// case "time zone converter":
	// 	timeZoneConvert()
	case "test":
		test()
		return depracatedComponent
	case "open Ksorter":
		ksorter_integration()
		return ksorterR
	case "open Kserver":
		kserver.StartKserver()
		return kserverR
	case "exit":
		log.Info("exiting...")
		return exitR
	}

	return exitR
}
