package commands
import (
	"fmt"
	"github.com/codegangsta/cli"
	"sort"
	"math/rand"
	. "github.com/popcorp/doenv/lib"
	"github.com/digitalocean/godo"
	"time"
)

var RestoreCommand = cli.Command{
	Name: "restore",
	Aliases: []string{"bring", "up"},
	Usage:  "Create a droplet based on a snapshot (Restore snapshot, start)",
	ArgsUsage: "<snapshot name> [droplet name]",
	Action: func(c *cli.Context) {
		var DropletName string
		var SnapshotName string
		if len(c.Args()) < 1 {
			fmt.Println("Meh. Where's the snapshot name ?! :(")
			return
		}
		SnapshotName = c.Args().Get(0)
		if len(c.Args()) < 2 {
			DropletName = SnapshotName
		}else {
			DropletName = c.Args().Get(1)
		}

		restoreDroplet(SnapshotName, DropletName)
	},
}

func restoreDroplet(SnapshotName string, DropletName string) {
	image, _ := ImageByName(SnapshotName)
	if image == nil {
		fmt.Printf("Could not find a snapshot named '%s'\n", SnapshotName)
		return
	}
	if image.Public {
		fmt.Println("Nope. You should only start your own droplets. And this looks like a public image :(")
		return
	}

	Region := image.Regions[rand.Intn(len(image.Regions))]
	var Size string
	sizes, _, _ := GetClient().Sizes.List(nil)
	sort.Sort(ByCost(sizes))
	for _, s := range sizes {
		if s.Disk >= image.MinDiskSize {
			Size = s.Slug
			break
		}
	}

	createRequest := &godo.DropletCreateRequest{
		Name:   DropletName,
		Region: Region,
		Size:   Size,
		Image: godo.DropletCreateImage{
			ID: image.ID,
		},
		IPv6: true,
		SSHKeys: GetFingerprints(),
	}

	fmt.Printf("Creation request done ...\nSize: %s - Region: %s - Hostname: %n", Size, Region, DropletName)
	d, resp, err := GetClient().Droplets.Create(createRequest)
	if err != nil {
		panic(err)
	}
	time.Sleep(15 * time.Second)
	WaitAction(resp.Links.Actions[0].ID)
	d, _, _ = GetClient().Droplets.Get(d.ID)
	fmt.Println("Hoora! Just setup your droplet :)")
	fmt.Printf("Your ip is %s\n", d.Networks.V4[0].IPAddress)
	return
}
