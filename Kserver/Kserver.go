package kserver

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mingrammer/cfmt"
)

// StartKserver generates a server to host files from a specified directory and may automatically forward the port.
func StartKserver() {
	dir := ""
	forward := ""

	dirInput := &survey.Input{
		Renderer: survey.Renderer{},
		Message:  "provide the directory to host:",
		Default:  "./",
		Help:     "provide an absolute path to the directory you want to open on the file server",
	}

	forwardInput := &survey.Select{
		Renderer: survey.Renderer{},
		Message:  "automatically forward the port to the server? (requires ngrok installed!!!)",
		Options:  []string{"yes", "no"},
		Default:  "no",
		Help:     "should the server forward the port automatically to make the server accessible from the internet",
	}

	survey.AskOne(dirInput, &dir)

	survey.AskOne(forwardInput, &forward)

	fileHandler := http.FileServer(http.Dir(dir))

	server := &http.Server{
		Addr:                         ":6969",
		Handler:                      fileHandler,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    &tls.Config{},
		ReadTimeout:                  60 * time.Second,
		ReadHeaderTimeout:            60 * time.Second,
		WriteTimeout:                 60 * time.Second,
		IdleTimeout:                  0,
		MaxHeaderBytes:               0,
		ErrorLog:                     &log.Logger{},
	}

	cfmt.Infoln("Server starting on port ':6969'...")
	go server.ListenAndServe()

	// fmt.Println(forward)

	if forward == "yes" {
		forwardPort()
	} else {
		cfmt.Successln("Server live on http://localhost:6969/")
	}
}

// forwardPort is a Go function that forwards a port using ngrok.
func forwardPort() {

	cfmt.Infoln("forwarding port :6969 through ngrok...")
	// forward with ngrok
	cmd := exec.Command("ngrok", "http", "6969")
	cmd.Stdin = os.Stdin

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cfmt.Successln("check your ngrok account for the host details")

	err := cmd.Run()
	if err != nil {
		// Handle error
		os.Stderr.WriteString(fmt.Sprintf("Error executing PowerShell script: %s\n", err.Error()))
		return
	}

	os.Stdout.WriteString(fmt.Sprintf("PowerShell Script Output:\n%s\n", stdout.String()))

}
