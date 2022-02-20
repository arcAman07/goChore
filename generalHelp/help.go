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
		fmt.Printf("%d  %+v\n",i,"")
		i++
		fmt.Printf("%d  %+v\n",i,"register")
		i++
		fmt.Printf("%d  \t%+v\n",i,"login")
		i++
		fmt.Printf("%d\t%+v\n",i,"add")
		i++
		fmt.Printf("%d\t%+v\n",i,"delete")
		i++
		fmt.Printf("%d\t%+v\n",i,"list")
		i++
		fmt.Printf("%d\t%+v\n",i,"finish")
		i++
		fmt.Printf("%d\t%+v\n",i,"remain")
		i++
		fmt.Printf("%d\t%+v\n",i,"update")
		i++
		fmt.Printf("%d\t%+v\n",i,"refresh")
		i++
		fmt.Printf("%d\t%+v\n",i,"stats")
		i++
		fmt.Printf("%d\t%+v\n",i,"undo")
		i++
		fmt.Printf("%d\t%+v\n",i,"logout")
		i++
		fmt.Printf("%d\t%+v\n",i,"exit")
		i++
		return
	}
}
