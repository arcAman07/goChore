package generalHelp
import (
	"fmt"
	"os"
	"github.com/fatih/color"
)

// Commands => 

// 1) register "Username" "Password"
// 2) login "Username" "Password"
// 3) add "Username" "TaskName"
// 4) delete "Username" "TaskName"
// 5) list "Username"
// 6) finish "Username" "TaskName"
// 7) remain "Username"
// 8) update "Username" "TaskName" "NewTaskName"

// 9) logout "Username"
// 10) exit

var i int = 1
func GeneralHelp() {
	color.Red("I am a god")
	command := os.Args[1]
	if command == "help" {
		fmt.Printf("%d\t%+v\n",i,"help")
	}
}
