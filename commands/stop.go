package commands
import (
	"github.com/codegangsta/cli"
	"fmt"
	. "github.com/popcorp/doenv/lib"
)

var StopCommand = cli.Command{
	Name: "stop",
	Usage:  "Shutdown a droplet",
	ArgsUsage: "<droplet name>",
	Action: func(c *cli.Context) {
		var DropletName string
		if len(c.Args()) < 1 {
			fmt.Println("Meh. Where's the droplet name ?! :(")
			return
		}
		DropletName = c.Args().Get(0)

		stopDroplet(DropletName)
	},
}

func stopDroplet(DropletName string) {
	d, _ := DropletByName(DropletName)
	if d == nil {
		fmt.Printf("Could not find a droplet named '%s'", DropletName)
		return
	}
	// status -> "new", "active", "off", or "archive".
	if d.Status == "active" {
		fmt.Println("Droplet powered on, stopping ...")
		action, _, err := GetClient().DropletActions.Shutdown(d.ID)
		if err != nil {
			panic(err)
		}
		WaitAction(action.ID)
		WaitStatus(d.ID, "off")
		d, _, _ = GetClient().Droplets.Get(d.ID)
	}
	if d.Status == "off" {
		fmt.Println("Droplet is now off :)")
	}
}
