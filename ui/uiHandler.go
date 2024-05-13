package ui

import (
	"github.com/charmbracelet/huh"
)

var ModePrompt = huh.NewSelect[string]().
	Options(
		huh.NewOption("notes", "notes"),
		huh.NewOption("currency convert", "currency convert"),
		huh.NewOption("open Ksorter", "open Ksorter"),
		huh.NewOption("list from directory", "list from directory"),
		huh.NewOption("open Kserver", "open Kserver"),
		huh.NewOption("funny", "funny"),
		huh.NewOption("pricer", "pricer"),
		huh.NewOption("fibonacci", "fibonacci"),
		huh.NewOption("exit", "exit"),
	).
	Title("app mode:")
