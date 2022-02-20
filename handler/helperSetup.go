package handler

import (
	"goChore/generalHelp"
	"os"
)

func Helper() {
	command := os.Args[1]
	if command == "help" {
		generalHelp.GeneralHelp()
		return
	}
}
