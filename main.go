package main

import (
	"fmt"
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

	// if err := keyboard.Open(); err != nil {
	// 	panic(err)
	// }
	// defer func() {
	// 	_ = keyboard.Close()
	// }()

	// fmt.Println("Press ESC to quit")
	// for {
	// 	char, key, err := keyboard.GetKey()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("end: rune %q, key %X\r\n", char, key)
	// 	if key == keyboard.KeyEsc {
	// 		break
	// 	}
	// }

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

	result := price * (1 - (discount / 100)) * (1 + tax)

	fmt.Println(result)

}
