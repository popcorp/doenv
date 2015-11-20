package commands
import (
	"fmt"
	"github.com/codegangsta/cli"
	. "github.com/popcorp/doenv/lib"
)

var EraseCommand = cli.Command{
	Name: "erase",
	Usage:  "Let's just drop this droplet.",
	ArgsUsage: "<droplet name>",
	Action:  func(c *cli.Context) {
		var DropletName string
		if len(c.Args()) < 1 {
			fmt.Println("Meh. Where's the droplet name ?! :(")
			return
		}
		DropletName = c.Args().Get(0)

		eraseDroplet(DropletName)
	},
}

func eraseDroplet(DropletName string) {
	d, _ := DropletByName(DropletName)
	if d == nil {
		fmt.Printf("Could not find a droplet named '%s'", DropletName)
		return
	}
	//"new", "active", "off", or "archive".
	//fmt.Println(d.Status)
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
	// println(d.Status)
	if d.Status == "off" {
		fmt.Println("Droplet is off :)")
		_, err := GetClient().Droplets.Delete(d.ID)
		if err != nil {
			panic(err)
		}
		fmt.Println("Droplet is gone :(")
	}
}
