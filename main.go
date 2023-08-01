package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"time"

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
		Options: []string{"sys info", "notes", "currency convert", "pricer", "fibonacci", "funny", "exit"},
	}

	survey.AskOne(prompt, &mode, survey.WithValidator(survey.Required))

	// fmt.Println("select app mode, options are: pricer, funny")
	// fmt.Scanln(&mode)

	switch {
	case mode == "pricer":
		pricer()
	case mode == "notes":
		notes()
	case mode == "messenger":
		messenger()
	case mode == "funny":
		funny()
	case mode == "sys info":
		sysInfo()
	case mode == "currency convert":
		currencyConvert()
	case mode == "fibonacci":
		fibonacciLuncher()
	case mode == "test":
		test()
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

	main()
}

func messenger() {

	action := ""

	promt := &survey.Select{
		Renderer: survey.Renderer{},
		Message:  "send messages or manage contacts:",
		Options:  []string{"send/read messages", "add/remove contacts"},
		Default:  nil,
	}

	survey.AskOne(promt, &action, survey.WithValidator(survey.Required))

	switch {
	case action == "send/read messages":
		messages()
	case action == "add/remove contacts":
		//contacts()
	}

}

func messages() {

	listOfContacts := []string{}
	selectedContactIndex := -1

	_, err := os.Stat("/Program Files (x86)/fuck you inc/ktool/contacts.txt")
	if os.IsNotExist(err) {
		// File not found, create it
		_, err = os.Create("/Program Files (x86)/fuck you inc/ktool/contacts.txt")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		fmt.Println(ctc.ForegroundYellow, "The contacts save file has been created. This should only happen once if you are on windows", ctc.Reset)
	}

	f, err := os.Open("/Program Files (x86)/fuck you inc/ktool/contacts.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		listOfContacts = append(listOfContacts, scanner.Text())
	}

	prompt := &survey.Select{
		Renderer: survey.Renderer{},
		Message:  "",
		Options:  listOfContacts,
		Default:  nil,
		Help:     "",
	}

	survey.AskOne(prompt, &selectedContactIndex, survey.WithValidator(survey.Required))

	if selectedContactIndex >= 0 && selectedContactIndex < len(listOfContacts) {
		selectedContact := listOfContacts[selectedContactIndex]
		fPath := fmt.Sprintf("%s/%s", "/Program Files (x86)/fuck you inc/ktool", selectedContact+".txt")
		f, err := os.OpenFile(fPath, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Println(ctc.ForegroundYellow, "Error opening message file for contact:", selectedContact, ctc.Reset)
			return
		}
		defer f.Close()

		DMscanner := bufio.NewScanner(f)
		for DMscanner.Scan() {
			fmt.Println(DMscanner.Text())
		}
	}
}

func notes() {

	var noteAction = ""

	promptNote := &survey.Select{
		Renderer: survey.Renderer{},
		Message:  "What to do:",
		Options:  []string{"add a note", "read notes", "delete notes"},
		Default:  nil,
		Help:     "select what you want to do with your notes",
	}

	survey.AskOne(promptNote, &noteAction, survey.WithValidator(survey.Required))

	switch {
	case noteAction == "add a note":
		addNote()
	case noteAction == "read notes":
		readNotes()
	case noteAction == "delete notes":
		deleteNotes()
	}

}

func addNote() {
	var newNote string

	promptNote := &survey.Input{
		Renderer: survey.Renderer{},
		Message:  "enter your note:",
		Default:  "",
		Help:     "enter a note here and it will be saved for later use",
	}

	survey.AskOne(promptNote, &newNote, survey.WithValidator(survey.Required))

	targetFile := "notes.txt"
	absPath := targetFile

	if runtime.GOOS == "windows" {
		root := "/Program Files (x86)/fuck you inc/ktool"
		absPath = filepath.Join(root, targetFile)
	}

	_, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		// File not found, create it
		_, err = os.Create(absPath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		fmt.Println(ctc.ForegroundYellow, "The note save file has been created. This should only happen once if you are on windows", ctc.Reset)
	}

	f, err := os.OpenFile(absPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	noteAssembled := fmt.Sprintf("%s\n", newNote)

	_, err = f.WriteString(noteAssembled)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	notes()
}

func readNotes() {

	if runtime.GOOS == "windows" {
		f, err := os.Open("/Program Files (x86)/fuck you inc/ktool/notes.txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)
		counter := 1

		for scanner.Scan() {
			fmt.Println(counter, scanner.Text())
			counter++
		}
	} else {
		f, err := os.Open("notes.txt")

		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)
		counter := 1

		for scanner.Scan() {
			fmt.Println(counter, scanner.Text())
			counter++
		}
	}

	notes()
}

func deleteNotes() {

	if runtime.GOOS == "windows" {
		f, err := os.Open("/Program Files (x86)/fuck you inc/ktool/notes.txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}

		defer f.Close()

		counter := 1

		notes := []string{}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			notes = append(notes, scanner.Text())
			counter++
		}

		notesToDelete := []string{}

		prompt := &survey.MultiSelect{
			Renderer: survey.Renderer{},
			Message:  "choose which notes to delete",
			Options:  notes,
			Default:  nil,
			Help:     "changes made here are permanent",
		}

		survey.AskOne(prompt, &notesToDelete, survey.WithValidator(survey.Required))

		fmt.Println(notesToDelete)

		err = f.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
			return
		}

		deleter(notesToDelete)
	}

	notes()
}

func deleter(notesToDelete []string) error {
	if runtime.GOOS == "windows" {

		file, err := os.Open("/Program Files (x86)/fuck you inc/ktool/notes.txt")
		if err != nil {
			return fmt.Errorf("error opening file: %v", err)
		}

		tmpFile, err := os.CreateTemp("", "notes.txt")
		if err != nil {
			return fmt.Errorf("error creating temporary file: %v", err)
		}

		defer os.Remove(tmpFile.Name())

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			isMatched := false
			for _, note := range notesToDelete {
				matched, _ := regexp.MatchString("^"+note+"$", line)
				if matched {
					isMatched = true
					break
				}
			}
			if !isMatched {
				_, err := io.WriteString(tmpFile, line+"\n")

				if err != nil {
					fmt.Printf("error writing to temporary file: %v", err)
				}
			}
		}

		err = file.Close()
		if err != nil {
			return fmt.Errorf("error closing file: %w", err)
		}

		err = tmpFile.Close()
		if err != nil {
			return fmt.Errorf("error closing temporary file: %w", err)
		}

		err = os.Remove("/Program Files (x86)/fuck you inc/ktool/notes.txt")
		if err != nil {
			fmt.Printf("error removing original file: %v", err)
		}

		err = os.Rename(tmpFile.Name(), "/Program Files (x86)/fuck you inc/ktool/notes.txt")
		if err != nil {
			fmt.Printf("error renaming temporary file: %v", err)
		}

	}

	return nil
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

	fmt.Println("<----------GO DRINK SOMETHING---------->")
	fmt.Println(ctc.Negative, joke.Joke, ctc.Reset)
	fmt.Println("<-------------------------------------->")

	main()
}

func sysInfo() {

	for {
		v, _ := mem.VirtualMemory()
		o, _ := host.Info()
		c, _ := cpu.Info()

		memoryInfo := fmt.Sprintf("Total: %v, Free:%v, UsedPercent:%f%%", v.Total, v.Free, v.UsedPercent)
		osInfo := fmt.Sprintf("OS: %v, Uptime: %v, Procs: %v", o.OS, o.Uptime, o.Procs)
		cpuInfo := fmt.Sprintf("Vendor: %v, Cores: %v, Mhz: %v, Model: %v", c[0].VendorID, c[0].Cores, c[0].Mhz, c[0].ModelName)

		fmt.Print("\033[2J")
		fmt.Print("\033[H")

		fmt.Println("<----------SYS INFO---------->")
		fmt.Println(ctc.ForegroundYellow, "System:", osInfo, ctc.Reset)
		fmt.Println(ctc.ForegroundBrightCyan, "Cpu:", cpuInfo, ctc.Reset)
		fmt.Println(ctc.ForegroundBrightGreenBackgroundBlack, "Memory:", memoryInfo, ctc.Reset)

		fmt.Println("<---------------------------->")

		time.Sleep(500 * time.Millisecond)
	}

}

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

func test() {

}
