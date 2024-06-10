package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// this is a luncher for my Ksorter powershell script
func ksorter_integration() {

	resp, err := http.Get("https://raw.githubusercontent.com/kociumba/Ksorter/main/Ksorter.ps1")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("Ksorter.ps1")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = io.Copy(f, resp.Body); err != nil {
		log.Fatal(err)
	}
	f.Close()

	filepath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("powershell.exe", "-ExecutionPolicy", "Unrestricted", fmt.Sprintf(filepath+"/Ksorter.ps1"))
	// fmt.Println(cmd.Args)
	cmd.Stdin = os.Stdin

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		// Handle error
		os.Stderr.WriteString(fmt.Sprintf("Error executing PowerShell script: %s\n", err.Error()))
		return
	}

	// Print the output of the PowerShell script
	os.Stdout.WriteString(fmt.Sprintf("PowerShell Script Output:\n%s\n", stdout.String()))

	os.Remove("Ksorter.ps1")
}
