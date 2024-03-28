package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/AlecAivazis/survey/v2"
	"github.com/wzshiming/ctc"
)

func currencyConvert() {

	amount := 0.0
	from := ""
	to := ""

	promptAmount := &survey.Input{
		Message: "select the amount:",
		Help:    "amaout of money to convert, currency nonspecific",
	}

	survey.AskOne(promptAmount, &amount, survey.WithValidator(survey.Required))

	promptFrom := &survey.Input{
		Message: "convert from:",
		Help:    "use an international currency code like 'USD', if it doesn't work it means the api i'm using doesn't support this currency",
	}

	survey.AskOne(promptFrom, &from, survey.WithValidator(survey.Required))

	promptTo := &survey.Input{
		Message: "convert to:",
		Help:    "use an international currency code like 'USD', if it doesn't work it means the api i'm using doesn't support this currency",
	}

	survey.AskOne(promptTo, &to, survey.WithValidator(survey.Required))

	client := http.Client{}

	url := fmt.Sprintf("https://api.exchangerate.host/convert?from=%v&to=%v", from, to)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	rate := result["info"].(map[string]interface{})["rate"].(float64)
	fmt.Println("<----------CONVERSION---------->")
	fmt.Println(ctc.ForegroundBright, "current exchange from", from, "to", to, "is", rate, ctc.Reset)

	exchangedAmount := math.Round((amount*rate)*100) / 100

	fmt.Println(ctc.ForegroundBrightGreen, amount, from, "is", exchangedAmount, to, ctc.Reset)
	fmt.Println("<------------------------------>")

	main()
}
