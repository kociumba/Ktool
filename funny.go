package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wzshiming/ctc"
)

// fetch dad jokes from api
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

	fmt.Println("<----------GO DRINK SOMETHING---------->")
	fmt.Println(ctc.Negative, joke.Joke, ctc.Reset)
	fmt.Println("<-------------------------------------->")

	main()
}
