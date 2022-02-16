package main

// Designing cli apps
import (
	"Go-Tasks/actions"
	"fmt"
	"log"
	"os"
)

var app = actions.App

func main() {
	fmt.Println("Hello World")
	actions.Actions()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
