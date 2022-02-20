package generalHelp
import (
	"fmt"
	"os"
	"github.com/fatih/color"
)

// Commands => 

// 1) register "Username" "Password" => 
// 2) login "Username" "Password" =>
// 3) add "Username" "TaskName" =>
// 4) delete "Username" "TaskName" =>
// 5) list "Username" =>
// 6) finish "Username" "TaskName" =>
// 7) remain "Username" =>
// 8) update "Username" "TaskName" "NewTaskName" =>
// 9) refresh "Username" =>
// 10) stats "Username" => "Total Tasks" "Total Tasks Done" "Total Tasks Not Done" =>
// 11) undo "Username" "TaskName" => Change its status
// 12) logout "Username" =>
// Adding Point system + leaderboard later if traffic increases

var i int = 1
var whilte = color.New(color.FgWhite)
var boldWhite = whilte.Add(color.Bold)
func GeneralHelp() {
	color.Red("I am a god")
	command := os.Args[1]
	if command == "help" {
		boldWhite.Println("Commands:")
		fmt.Printf("%d\t%+v\n",i,"register \"Username\" \"Password\"\t=>\tRegisters a new user")
		i++
		fmt.Println()
		fmt.Printf("%d\t%+v\n",i,"login \"Username\" \"Password\"\t=>\tLogs in the registered user")
		i++
		fmt.Println()
		fmt.Printf("%d\t%+v\n",i,"logout \"Username\"\t=>\t\tLogs out the registered user")
		i++
		fmt.Println()
		fmt.Printf("%d\t%+v\n",i,"add \"Username\" \"TaskName\"=>\tAdds a new task to the registered user")
		i++
		fmt.Println()
		fmt.Printf("%d\t%+v\n",i,"delete \"Username\" \"TaskName\"=>\tDeletes a task from the registered user")
		i++
		fmt.Println()
		fmt.Printf("%d\t%+v\n",i,"list \"Username\"=>\t\tLists all the tasks of the registered user")
		i++
		fmt.Println()
		fmt.Printf("%d\t%+v\n",i,"finish \"Username\" \"TaskName\"=>\tMarks a task as done")
		i++
		fmt.Println()
		fmt.Printf("%d\t%+v\n",i,"remain \"Username\"=>\t\tLists all the tasks that are not done")
		i++
		fmt.Println()
		fmt.Printf("%d\t%+v\n",i,"update \"Username\" \"TaskName\" \"NewTaskName\"=>\tUpdates a task")
		i++
		fmt.Println()
		fmt.Printf("%d\t%+v\n",i,"refresh \"Username\"=>\t\tClears the list of all tasks")
		i++
		fmt.Println()
		fmt.Printf("%d\t%+v\n",i,"stats \"Username\"=>\t\tLists the stats of the registered user")
		i++
		fmt.Println()
		fmt.Printf("%d\t%+v\n",i,"undo \"Username\" \"TaskName\"=>\tUndoes a task")
		i++
		fmt.Println()
		fmt.Printf("%d\t%+v\n",i,"help =>\tDisplays the list of available commands")
		i++
		fmt.Println()
		fmt.Printf("%d\t%+v\n",i,"exit\t=>\tTerminates the CLI application")
		i++
		fmt.Println()
		return
	}
}
