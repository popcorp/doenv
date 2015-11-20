package commands
import (
	"fmt"
	"github.com/codegangsta/cli"
	"time"
	. "github.com/popcorp/doenv/lib"
)

var FreezeCommand = cli.Command{
	Name: "freeze",
	Usage:  "Freeze a droplet (Stop, Snapshot and Destroy)",
	ArgsUsage: "<droplet name>",
	Action:  func(c *cli.Context) {
		var DropletName string
		if len(c.Args()) < 1 {
			fmt.Println("Meh. Where's the droplet name ?! :(")
			return
		}
		DropletName = c.Args().Get(0)

		freezeDroplet(DropletName)
	},
}

func freezeDroplet(DropletName string) {
	d, _ := DropletByName(DropletName)
	if d == nil {
		fmt.Printf("Could not find '%s'", DropletName)
		return
	}
	//"new", "active", "off", or "archive".
	fmt.Println(d.Status)
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
		fmt.Println("Droplet is off :)")
		fmt.Printf("Creation of Snapshot with name %s\n", d.Name)
		action, _, _ := GetClient().DropletActions.Snapshot(d.ID, d.Name)
		time.Sleep(15 * time.Second)
		WaitAction(action.ID)
		fmt.Println("Snapshot: done")
		fmt.Println("Destroying droplet")
		_, err := GetClient().Droplets.Delete(d.ID)
		if err != nil {
			panic(err)
		}
		fmt.Println("Droplet destroyed ! :( That makes me sad ;(")
	}
}
