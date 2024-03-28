package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlecAivazis/survey/v2"

	"github.com/mingrammer/cfmt"
)

func main() {

	modeSelect()

}

func modeSelect() {

	mode := ""

	prompt := &survey.Select{
		Message: "app mode:",
		Options: []string{"sys info", "notes", "currency converter", "list from directory", "time zone converter", "pricer", "fibonacci", "funny", "exit"},
	}

	survey.AskOne(prompt, &mode, survey.WithValidator(survey.Required))

	// fmt.Println("select app mode, options are: pricer, funny")
	// fmt.Scanln(&mode)

	switch {
	case mode == "pricer":
		pricer()
	case mode == "notes":
		notes()
	case mode == "list from directory":
		listFromDirectory()
	case mode == "messenger":
		messenger()
	case mode == "funny":
		funny()
	case mode == "sys info":
		sysInfo()
	case mode == "time zone converter":
		timeZoneConvert()
	case mode == "currency convert":
		currencyConvert()
	case mode == "fibonacci":
		fibonacciLuncher()
	case mode == "test":
		test()
	case mode == "exit":
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
