package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

func main() {

	modeSelect()

}

func modeSelect() {
	var mode = ""

	fmt.Println("select app mode, options are: pricer, funny")
	fmt.Scanln(&mode)

	switch {
	case mode == "pricer":
		pricer()
	case mode == "funny":
		funny()
	}
}

func pricer() {

	var price = 0.0
	var discount = 0.0
	var tax_version = ""
	var tax = 0.0

	fmt.Println("price")
	fmt.Scanln(&price)

	fmt.Println("discount")
	fmt.Scanln(&discount)

	fmt.Println("apply tax ?")
	fmt.Scanln(&tax_version)

	switch {
	case tax_version == "poland":
		tax = 0.23
	default:
		tax = 0.0
	}

	result := math.Round(price*(1-(discount/100))*(1+tax)*100) / 100

	fmt.Println(result)

}

func funny() {
	type JokeResponse struct {
		ID     string `json:"id"`
		Joke   string `json:"joke"`
		Status int    `json:"status"`
	}

	client := http.Client{}

	request, err := http.NewRequest("GET", "https://icanhazdadjoke.com", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	request.Header.Set("Accept", "application/json")

	request.Header.Set("User-Agent", "My Library (https://github.com/kociumba/ktool)")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	defer response.Body.Close()

	var joke JokeResponse
	err = json.NewDecoder(response.Body).Decode(&joke)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	fmt.Println(joke.Joke)

}
