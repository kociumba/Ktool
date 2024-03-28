package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/wzshiming/ctc"
)

func fibonacciCalc(n int, cache map[int]int) int {
	if val, ok := cache[n]; ok {
		return val
	}

	var result int
	if n == 0 {
		result = 0
	} else if n == 1 {
		result = 1
	} else {
		result = fibonacciCalc(n-1, cache) + fibonacciCalc(n-2, cache)
	}

	cache[n] = result
	return result

}

// calculate the fibonacci sequence
func fibonacciLuncher() {
	var input string

	promt := &survey.Input{
		Message: "enter a number: ",
		Help:    "enter a number to get the fibonacci number for that number (this calculation breaks somewhere around the 1000th place in the sequence couse i'm lazy)",
	}

	survey.AskOne(promt, &input, survey.WithValidator(survey.Required))

	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(ctc.ForegroundRed, "Invalid input. Please enter an integer.", ctc.Reset)
		os.Exit(1)
	}

	cache := make(map[int]int)

	for i := 0; i <= n; i++ {
		msg := fmt.Sprintf("Fibonacci(%d) = %d\n", i, fibonacciCalc(i, cache))

		if i == n {
			fmt.Println("<----------FIBONACCI SEQUENCE---------->")
			fmt.Println(" ")
			fmt.Println(ctc.ForegroundBrightGreen, msg, ctc.Reset)
			fmt.Println("<-------------------------------------->")
		}
	}

	main()
}
