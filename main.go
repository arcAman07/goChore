package main

// Designing cli apps
import (
	"Go-Tasks/actions"
	"fmt"
	"log"
	"os"
	"Go-Tasks/models"
)

var app = actions.App

func main() {
	// fmt.Println("Hello World")
	x := models.Date
	fmt.Println(x)
	actions.Actions()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
