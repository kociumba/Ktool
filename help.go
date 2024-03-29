package main

import (
	"github.com/mingrammer/cfmt"
)

func help() {
	cfmt.Infoln("<---ktool help--->")

	cfmt.Infoln("To use ktool, just run the binary or pass in an argument from the following list:")
	cfmt.Infoln("-currency \n-ksorter \n-ls \n-funny \n-pricer \n-fibonacci \n-notes \n-time \n-help")

	cfmt.Infoln("for more info visit:")
	cfmt.Successln("https://github.com/kociumba/ktool/wiki")
}
