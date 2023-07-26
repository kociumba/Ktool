package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/AlecAivazis/survey/v2"

	"github.com/mingrammer/cfmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"

	"github.com/wzshiming/ctc"
)

func main() {

	modeSelect()

}

func modeSelect() {

	mode := ""

	prompt := &survey.Select{
		Message: "app mode:",
		Options: []string{"pricer", "sys info", "funny", "exit"},
	}

	survey.AskOne(prompt, &mode, survey.WithValidator(survey.Required))

	// fmt.Println("select app mode, options are: pricer, funny")
	// fmt.Scanln(&mode)

	switch {
	case mode == "pricer":
		pricer()
	case mode == "funny":
		funny()
	case mode == "sys info":
		sysInfo()
	case mode == "exit":
		cfmt.Errorln("exiting...")
		return
	}
}

func pricer() {

	var price = 0.0
	var discount = 0.0
	var tax_version = ""
	var tax = 0.0

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

	cfmt.Successln(result)

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

	fmt.Println(ctc.Negative, joke.Joke, ctc.Reset)

}

func sysInfo() {
	v, _ := mem.VirtualMemory()
	o, _ := host.Info()
	c, _ := cpu.Info()

	memoryInfo := fmt.Sprintf("Total: %v, Free:%v, UsedPercent:%f%%", v.Total, v.Free, v.UsedPercent)
	osInfo := fmt.Sprintf("OS: %v, Uptime: %v", o.OS, o.Uptime)
	cpuInfo := fmt.Sprintf("Vendor: %v, Model: %v", c[0].VendorID, c[0].ModelName)

	fmt.Println("<----------SYS INFO---------->")
	fmt.Println(ctc.ForegroundYellow, "System:", osInfo, ctc.Reset)
	fmt.Println(ctc.ForegroundBrightCyan, "Cpu:", cpuInfo, ctc.Reset)
	fmt.Println(ctc.ForegroundBrightGreenBackgroundBlack, "Memory:", memoryInfo, ctc.Reset)
	fmt.Println("<---------------------------->")
}
