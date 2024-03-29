package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	switch args[1] {
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
	default:
		modeSelect()
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

func timeZoneConvert() { //TODO make it work

	client := http.Client{}

	requestFrom, err := http.NewRequest("GET", "https://worldtimeapi.org/api/timezone", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	resp, err := client.Do(requestFrom)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println(result)

}

func test() { // used only for new feature testing

}
