package handler

import (
	"os"
	"goChore/generalHelp"
)

func Helper() {
	command := os.Args[1]
	if command == "help" {
		generalHelp.GeneralHelp()
		return
	}
}