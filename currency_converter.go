package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/wzshiming/ctc"
)

// fetches the latest exchange rate from an API and performs the currency conversion.
func CurrencyConvert() {

	var amount float64
	from := ""
	to := ""

	type CurrencyResponse struct {
		Amount float64            `json:"amount"`
		Base   string             `json:"base"`
		Date   string             `json:"date"`
		Rates  map[string]float64 `json:"rates"`
	}

	promptAmount := &survey.Input{
		Message: "select the amount:",
		Help:    "amaout of money to convert, currency nonspecific",
	}

	if err := survey.AskOne(promptAmount, &amount, survey.WithValidator(survey.Required)); err != nil {
		panic(err)
	}

	promptFrom := &survey.Input{
		Message: "convert from:",
		Help:    "use an international currency code like 'USD', if it doesn't work it means the api i'm using doesn't support this currency",
	}

	if err := survey.AskOne(promptFrom, &from, survey.WithValidator(survey.Required)); err != nil {
		panic(err)
	}

	promptTo := &survey.Input{
		Message: "convert to:",
		Help:    "use an international currency code like 'USD', if it doesn't work it means the api i'm using doesn't support this currency",
	}

	if err := survey.AskOne(promptTo, &to, survey.WithValidator(survey.Required)); err != nil {
		panic(err)
	}

	client := http.Client{}

	from = strings.ToUpper(from)
	to = strings.ToUpper(to)

	url := fmt.Sprintf("https://www.frankfurter.app/latest?from=%v&to=%v", from, to)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	var response CurrencyResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		panic(err)
	}

	fmt.Println("<----------CONVERSION---------->")
	fmt.Println(ctc.ForegroundBright, "current exchange from", from, "to", to, ctc.Reset)

	rate, ok := response.Rates[to]
	if !ok {
		// Try converting 'to' to uppercase in case of mismatch
		to = strings.ToUpper(to)
		rate, ok = response.Rates[to]
	}

	if !ok {
		fmt.Println("Exchange rate not available for the selected currency pair")
	} else {
		exchangedAmount := math.Round((amount*rate)*100) / 100
		fmt.Println(ctc.ForegroundBrightGreen, amount, from, "is", exchangedAmount, to, ctc.Reset)
	}

	fmt.Println("<------------------------------>")

	main()
}
