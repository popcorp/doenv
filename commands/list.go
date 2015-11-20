package commands
import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
	"os"
	. "github.com/popcorp/doenv/lib"
)

var ListCommand = cli.Command{
	Name: "list",
	Aliases: []string{"ls"},
	Usage:  "List your droplets",
	Action: func(c *cli.Context) {
		// Display in this order so that you may read without scrolling :)
		listImages("both")
		fmt.Println()
		listDroplets()
	},
	Subcommands: []cli.Command{
		{
			Name: "droplets",
			Aliases: []string{"d"},
			Usage:  "List your droplets",
			Action: func(c *cli.Context) {
				listDroplets()
			},
		},
		{
			Name: "snapshots",
			Aliases: []string{"s"},
			Usage:  "List your snapshots",
			Action: func(c *cli.Context) {
				listImages("snapshots")
			},
		},
		{
			Name: "image",
			Aliases: []string{"i"},
			Usage:  "List the publicly available images",
			Action: func(c *cli.Context) {
				listImages("images")
			},
		},
	},
}

func listImages(Filter string) {
	fmt.Println("Fecthing images list...")
	snapshotList := [][]string{}
	publicList := [][]string{}

	list, _ := ImagesList()
	for _, i := range list {
		if i.Public {
			data := []string{i.Distribution, i.Slug, i.Name, i.Created, fmt.Sprintf("%d", i.MinDiskSize)}
			publicList = append(publicList, data)
		}else {
			data := []string{i.Distribution, i.Name, i.Created, fmt.Sprintf("%d", i.MinDiskSize)}
			snapshotList = append(snapshotList, data)
		}
	}
	if Filter == "Both" || Filter == "images" {
		// There can't be no public images. :)
		fmt.Println("Public images")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"OS", "Slug", "Name", "Creation date", "Min disk"})
		table.SetBorder(false)
		table.AppendBulk(publicList)
		table.Render()

	}

	if Filter == "Both" {
		fmt.Println()
	}
	if Filter == "Both" || Filter == "snapshots" {
		if len(snapshotList) == 0 {
			fmt.Println("You don't have any snapshots.")
		}else {
			fmt.Println("Your snapshots")

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"OS", "Name", "Creation date", "Min disk"})
			table.SetBorder(false)
			table.AppendBulk(snapshotList)
			table.Render()
		}
	}
}

func listDroplets() {
	fmt.Println("Fecthing droplets list...")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Status", "Name", "Created", "Plan", "IPv4"})
	table.SetBorder(false)
	list, _ := DropletsList()

	if len(list) == 0 {
		fmt.Println("You don't have any droplet.")
		return
	}

	for _, d := range list {
		table.Append([]string{fmt.Sprintf("%d", d.ID), d.Status, d.Name, d.Created, d.SizeSlug, d.Networks.V4[0].IPAddress})
	}
	fmt.Println("Your droplets")
	table.Render()
}
