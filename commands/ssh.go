package commands
import (
	"github.com/codegangsta/cli"
	"fmt"
	. "github.com/popcorp/doenv/lib"
)

var SshCommand = cli.Command{
	Name: "ssh",
	Aliases: []string{"enter", "hack"},
	Usage:  "Hack into a droplet!",
	ArgsUsage: "<droplet name>",
	Action:  func(c *cli.Context) {
		var DropletName string
		if len(c.Args()) < 1 {
			fmt.Println("Meh. Where's the droplet name ?! :(")
			return
		}
		DropletName = c.Args().Get(0)

		sshDroplet(DropletName)
	},
}

func sshDroplet(DropletName string) {
	d, _ := DropletByName(DropletName)
	if d == nil {
		fmt.Printf("Could not find a droplet named '%s'\n", DropletName)
		return
	}
	if d.Status != "active" {
		fmt.Println("Droplet not started ... :(")
		return
	}
	fmt.Println("Found the droplet ! Let's ssh into it")
	SshDroplet(d)
}
