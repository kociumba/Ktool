package main

import (
	"fmt"
	"math"
)

func main() {

	var mode = ""

	fmt.Println("select app mode, options are: pricer")
	fmt.Scanln(&mode)

	switch {
	case mode == "pricer":
		pricer()
		fmt.Println(mode)
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
