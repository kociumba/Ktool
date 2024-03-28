package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/wzshiming/ctc"
)

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

		// the actual connection and messaging
		// fuck

		DMscanner := bufio.NewScanner(f)
		for DMscanner.Scan() {
			fmt.Println(DMscanner.Text())
		}
	}
}
