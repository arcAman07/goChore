package generalHelp
import (
	"fmt"
	"os"
	"github.com/fatih/color"
)

var i int = 1
func GeneralHelp() {
	color.Red("I am a god")
	command := os.Args[1]
	if command == "help" {
		fmt.Printf("%d\t%+v\n",i,"help")
	}
}
