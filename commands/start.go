package commands
import (
	"github.com/codegangsta/cli"
	"fmt"
	. "github.com/popcorp/doenv/lib"
)

var StartCommand = cli.Command{
	Name: "start",
	Usage:  "Start a droplet (PowerOn)",
	ArgsUsage: "<droplet name>",
	Action: func(c *cli.Context) {
		var DropletName string
		if len(c.Args()) < 1 {
			fmt.Println("Meh. Where's the droplet name ?! :(")
			return
		}
		DropletName = c.Args().Get(0)

		startDroplet(DropletName)
	},
}

func startDroplet(DropletName string) {
	d, _ := DropletByName(DropletName)
	if d == nil {
		fmt.Printf("Could not find a droplet named '%s'", DropletName)
		return
	}
	// status -> "new", "active", "off", or "archive".
	// fmt.Println(d.Status)
	if d.Status == "off" {
		fmt.Println("Droplet powered off, starting ...")
		action, _, err := GetClient().DropletActions.PowerOn(d.ID)
		if err != nil {
			panic(err)
		}
		WaitAction(action.ID)
		WaitStatus(d.ID, "active")
		d, _, _ = GetClient().Droplets.Get(d.ID)
	}
	// fmt.Println(d.Status)
	if d.Status == "active" {
		fmt.Println("Droplet is now active :)")
	}
}
