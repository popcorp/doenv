package commands

import (
	"fmt"
	"os"
	"github.com/codegangsta/cli"
	. "github.com/popcorp/doenv/lib"
)

var SetTokenCommand = cli.Command{
	Name: "set-token",
	Usage:  "Set your Personnal API Token (PAT)",
	Action:
	func(c *cli.Context) {
		if len(c.Args()) < 1 {
			fmt.Println("Sorry, but I can't read what's in your mind.")
			fmt.Println("Just give me your Personnal App Token and we'll be friends. :)")
		}else {
			SetPersonnalApiToken(c.Args().Get(0))
			fmt.Println("Great! Now we can work together.")
			fmt.Println("How about creating a small project ? Would you ? Just run `doenv init my-project` :)")
			SaveSettings()
			os.Exit(0)
		}
	},
}
