package commands
import (
	"github.com/codegangsta/cli"
	"fmt"
	"time"
	. "github.com/popcorp/doenv/lib"
)

var SnapshotCommand = cli.Command{
	Name: "snapshot",
	Aliases: []string{"snap", "shot"},
	Usage:  "Snapshot a droplet to reuse it later",
	ArgsUsage: "<droplet name> [snapshot name]",
	Action:  func(c *cli.Context) {
		var DropletName string
		var SnapshotName string
		if len(c.Args()) < 1 {
			fmt.Println("Meh. Where's the droplet name ?! :(")
			return
		}
		DropletName = c.Args().Get(0)
		if len(c.Args()) < 2 {
			SnapshotName = DropletName
		}else {
			SnapshotName = c.Args().Get(1)
		}
		snapshotDroplet(DropletName, SnapshotName)
	},
}

func snapshotDroplet(DropletName, SnapshotName string) {
	d, _ := DropletByName(DropletName)

	if d == nil {
		fmt.Printf("Could not find '%s'", DropletName)
		return
	}

	oldStatus := d.Status

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
	//println(d.Status)

	if d.Status == "off" {
		fmt.Println("Droplet is off")
		fmt.Printf("Creating a Snapshot named '%s'\n", d.Name)
		action, _, _ := GetClient().DropletActions.Snapshot(d.ID, SnapshotName)
		time.Sleep(15 * time.Second)
		WaitAction(action.ID)
		fmt.Printf("Snapshot created, name: %s\n", SnapshotName)
		fmt.Println("Destroying droplet")
		if oldStatus == "active" {
			fmt.Println("Starting the droplet")
			action, _, _ := GetClient().DropletActions.PowerOn(d.ID)
			time.Sleep(15 * time.Second)
			WaitAction(action.ID)
			fmt.Println("Started :D")
		}
	}
}
