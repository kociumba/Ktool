package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/atotto/clipboard"
)

func listFromDirectory() {
	dir := ""
	fileList := ""

	prompt := &survey.Input{
		Message: "select the directory (you can use env variables):",
		Help:    "this is the directory that will be read",
	}

	survey.AskOne(prompt, &dir, survey.WithValidator(survey.Required))

	files, err := os.ReadDir(os.ExpandEnv(dir))
	if err != nil {
		log.Fatal(err)
	}

	outputFile, err := os.Create("fileList.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		outputFile.WriteString(file.Name() + "\n")
		fileList += file.Name() + "\n"
	}

	err = clipboard.WriteAll(fileList)
	if err != nil {
		fmt.Println("Error copying to clipboard:", err)
		return
	}

	log.Println("File names logged to fileList.txt and copied to clipboard")
}
