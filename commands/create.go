package commands

import (
	"github.com/codegangsta/cli"
	"github.com/docker/docker/pkg/namesgenerator"
	"fmt"
	"time"
	"github.com/digitalocean/godo"
	. "github.com/popcorp/doenv/lib"
)

var CreateCommand = cli.Command{
	Name: "create",
	Aliases: []string{"c", "init", "setup"},
	Usage:  "Create a droplet (Initial setup)",
	ArgsUsage: "<droplet name> [image] [region] [size]",
	Action: func(c *cli.Context) {
		var DropletName string
		var Image string
		var Region string
		var Size string
		if len(c.Args()) < 1 {
			DropletName = namesgenerator.GetRandomName(1)
		}else {
			DropletName = c.Args().Get(0)
		}
		if len(c.Args()) < 2 {
			Image = "debian-8-x64"
		}else {
			Image = c.Args().Get(1)
		}
		if len(c.Args()) < 3 {
			Size = "512mb"
		}else {
			Size = c.Args().Get(2)
		}
		if len(c.Args()) < 4 {
			Region = "nyc2"
		}else {
			Region = c.Args().Get(3)
		}
		createDroplet(DropletName, Image, Size, Region)
	},
}

func createDroplet(DropletName string, Image string, Size string, Region string) {
	createRequest := &godo.DropletCreateRequest{
		Name:   DropletName,
		Region: Region,
		Size:   Size,
		Image: godo.DropletCreateImage{
			Slug: Image,
		},
		IPv6: true,
		SSHKeys: GetFingerprints(),
	}

	fmt.Printf("Creating a droplet named %s, Size: %s - Region: %s\n", DropletName, Size, Region)
	newDroplet, resp, err := GetClient().Droplets.Create(createRequest)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n\n", err)
		panic(err)
	}

	time.Sleep(10 * time.Second)
	WaitAction(resp.Links.Actions[0].ID)
	newDroplet, _, _ = GetClient().Droplets.Get(newDroplet.ID)
	fmt.Println("Hoora! Just setup your droplet :)")
	fmt.Println(newDroplet.String())
	fmt.Println("Your IP is ", newDroplet.Networks.V4[0].IPAddress)
}
