package TaskCli
import (
	"fmt"
	"os"
)

func UpdateTask() (string, string) {
	command := os.Args[1]
	currentTask := os.Args[2]
	newTask := os.Args[3]
	if command == "update" {
		fmt.Println("Updating the task...")
		return currentTask, newTask
	}
	return "", ""
}