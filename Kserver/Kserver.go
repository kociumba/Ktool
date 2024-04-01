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

type serverConf struct {
	forward string
	dir     string
	port    string
}

// StartKserver generates a server to host files from a specified directory and may automatically forward the port.
func StartKserver() {
	config := serverConf{}

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

	portInput := &survey.Input{
		Renderer: survey.Renderer{},
		Message:  "what port to host on: (default: 6969)",
		Default:  "6969",
		Help:     "choose which port to host on and possibly forward with ngrok",
	}

	survey.AskOne(dirInput, &config.dir)

	survey.AskOne(forwardInput, &config.forward)

	survey.AskOne(portInput, &config.port)

	fileHandler := http.FileServer(http.Dir(config.dir))

	server := &http.Server{
		Addr:                         ":" + config.port,
		Handler:                      fileHandler,
		DisableGeneralOptionsHandler: false,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
		},
		ReadTimeout:       60 * time.Second,
		ReadHeaderTimeout: 60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		ErrorLog:          &log.Logger{},
	}

	cfmt.Infoln("Server starting on port ':" + config.port + "'...")

	// fmt.Println(forward)

	if config.forward == "yes" {
		go server.ListenAndServe()
		forwardPort(config.port)
	} else {
		cfmt.Successln("Server live on http://localhost:" + config.port + "/")
		server.ListenAndServe()
	}
}

// forwardPort is a Go function that forwards a port using ngrok.
func forwardPort(port string) {
	cfmt.Infoln("forwarding port :" + port + " through ngrok...")

	// forward with ngrok
	cmd := exec.Command("ngrok", "http", port)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		// Handle error
		os.Stderr.WriteString(fmt.Sprintf("Error executing ngrok command: %s\n", err.Error()))
		os.Stderr.WriteString("You most likely don't have ngrok installed\nTry insatlling it with `scoop install main/ngrok`")
	}

	// Print the output of the ngrok command
	// os.Stdout.WriteString(fmt.Sprintf("ngrok command output:\n%s\n", stdout.String()))
	// os.Stderr.WriteString(fmt.Sprintf("ngrok command error output:\n%s\n", stderr.String()))
}
