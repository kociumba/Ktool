package main

import (
	"fmt"
	"math"

	"github.com/AlecAivazis/survey/v2"
	"github.com/wzshiming/ctc"
)

// implement pricer
func pricer(repeat bool) {

	var price = 0.0
	var discount = 0.0
	var tax_version = ""
	var tax = 0.0
	var isItScam = ""

	promptPrice := &survey.Input{
		Message: "select the price:",
		Help:    "the currency that's inputed here will also be the result",
	}

	survey.AskOne(promptPrice, &price, survey.WithValidator(survey.Required))

	promptDiscount := &survey.Input{
		Message: "select the discount:",
		Help:    "this is trated as a % so a bogus value like 78834 won't work",
	}

	survey.AskOne(promptDiscount, &discount, survey.WithValidator(survey.Required))

	promptTax := &survey.Select{
		Message: "use tax ?",
		Options: []string{"poland", "no tax"},
	}

	survey.AskOne(promptTax, &tax_version, survey.WithValidator(survey.Required))

	switch {
	case tax_version == "poland":
		tax = 0.23
	default:
		tax = 0.0
	}

	result := math.Round(price*(1-(discount/100))*(1+tax)*100) / 100

	fmt.Println("<----------DISCOUNTED PRICE---------->")
	fmt.Println(ctc.ForegroundBrightGreen, " the discounted price is ", result, ctc.Reset)
	if result > price {
		isItScam = "the government is scamming you"
		fmt.Println(ctc.ForegroundBright, isItScam, ctc.Reset)
	}
	fmt.Println("<------------------------------------>")

	if repeat {
		main()
	}
}
