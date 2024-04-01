package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"regexp"

	"github.com/AlecAivazis/survey/v2"
)

// note actions handler
func notes() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	homeDir := fmt.Sprintf(u.HomeDir + `/ktool/notes.txt`)
	err = os.MkdirAll(u.HomeDir+"/ktool", 0700)
	if err != nil {
		log.Fatal(err)
	}

	noteAction := ""

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
		addNote(homeDir)
	case noteAction == "read notes":
		readNotes(homeDir)
	case noteAction == "delete notes":
		deleteNotes(homeDir)
	}

}

// add notes and write to disk
func addNote(homeDir string) {
	var newNote string

	promptNote := &survey.Input{
		Renderer: survey.Renderer{},
		Message:  "enter your note:",
		Default:  "",
		Help:     "enter a note here and it will be saved for later use",
	}

	survey.AskOne(promptNote, &newNote, survey.WithValidator(survey.Required))

	// targetFile := "notes.txt"
	// absPath := targetFile

	// if runtime.GOOS == "windows" {
	// 	root := "/Program Files (x86)/fuck you inc/ktool"
	// 	absPath = filepath.Join(root, targetFile)
	// }

	_, err := os.Stat(homeDir)
	if os.IsNotExist(err) {
		// File not found, create it
		_, err = os.Create(homeDir)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		// fmt.Println(ctc.ForegroundYellow, "The note save file has been created. This should only happen once if you are on windows", ctc.Reset)
	}

	f, err := os.OpenFile(homeDir, os.O_APPEND|os.O_WRONLY, 0600)
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

	// notes()
}

// read notes from disk
func readNotes(homedir string) {

	// if runtime.GOOS == "windows" {
	// 	f, err := os.Open("/Program Files (x86)/fuck you inc/ktool/notes.txt")
	// 	if err != nil {
	// 		fmt.Println("Error opening file:", err)
	// 		return
	// 	}

	// 	defer f.Close()

	// 	scanner := bufio.NewScanner(f)
	// 	counter := 1

	// 	for scanner.Scan() {
	// 		fmt.Println(counter, scanner.Text())
	// 		counter++
	// 	}
	// } else {
	// 	f, err := os.Open("notes.txt")

	// 	if err != nil {
	// 		fmt.Println("Error opening file:", err)
	// 		return
	// 	}

	// 	defer f.Close()

	// 	scanner := bufio.NewScanner(f)
	// 	counter := 1

	// 	for scanner.Scan() {
	// 		fmt.Println(counter, scanner.Text())
	// 		counter++
	// 	}
	// }

	f, err := os.Open(homedir)
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

	// notes()
}

// handler for deleting notes
func deleteNotes(homedir string) {
	// if runtime.GOOS == "windows" {
	f, err := os.Open(homedir)
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

	deleter(notesToDelete, homedir)
	// }

	// notes()
}

// deletes notes from disk
func deleter(notesToDelete []string, homedir string) error {
	// if runtime.GOOS == "windows" {

	file, err := os.Open(homedir)
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

	err = os.Remove(homedir)
	if err != nil {
		fmt.Printf("error removing original file: %v", err)
	}

	err = os.Rename(tmpFile.Name(), homedir)
	if err != nil {
		fmt.Printf("error renaming temporary file: %v", err)
	}

	// }

	return nil
}
