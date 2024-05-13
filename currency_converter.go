package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
)

var (
	amount    float64
	amountStr string
	from      string
	to        string
)

type CurrencyResponse struct {
	Amount float64            `json:"amount"`
	Base   string             `json:"base"`
	Date   string             `json:"date"`
	Rates  map[string]float64 `json:"rates"`
}

// fetches the latest exchange rate from an API and performs the currency conversion.
func CurrencyConvert(repeat bool) {

	huh.NewInput().
		Placeholder("100").
		Description("amount of money to convert, currency nonspecific").
		Value(&amountStr).
		Run()

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		log.Fatal(err)
	}

	huh.NewInput().
		Placeholder("USD").
		Description("convert from:").
		Value(&from).
		Run()

	huh.NewInput().
		Placeholder("EUR").
		Description("convert to:").
		Value(&to).
		Run()

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
	defer resp.Body.Close()

	var response CurrencyResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		panic(err)
	}

	rate, ok := response.Rates[to]
	if !ok {
		// Try converting 'to' to uppercase in case of mismatch
		to = strings.ToUpper(to)
		rate, ok = response.Rates[to]
	}

	var result string

	if !ok {
		result = "Exchange rate not available for the selected currency pair"
	} else {
		exchangedAmount := math.Round((amount*rate)*100) / 100
		result = fmt.Sprintf("%v %v is %v %v", amount, from, exchangedAmount, to)
	}

	output := fmt.Sprintf("<----------CONVERSION---------->\n\ncurrent exchange from %v to %v\n%v\n\n<----------------------------->", from, to, result)

	huh.NewText().Lines(6).Value(&output).Run()

	if repeat {
		main()
	}
}
