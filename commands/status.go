package commands
import (
	"github.com/codegangsta/cli"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	. "github.com/popcorp/doenv/lib"
)

var StatusCommand = cli.Command{
	Name: "status",
	Usage:  "Fetch a droplet status",
	ArgsUsage: "<droplet name ...>",
	Action:  func(c *cli.Context) {
		statusDroplet(c.Args())
	},
}

func statusDroplet(Names []string) {
	fmt.Println("Fecthing droplets list...")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Status", "Name", "Created", "Plan", "IPv4"})
	table.SetBorder(false)
	list, _ := DropletsList()
	for _, d := range list {
		if Contains(Names, d.Name) {
			table.Append([]string{fmt.Sprintf("%d", d.ID), d.Status, d.Name, d.Created, d.SizeSlug, d.Networks.V4[0].IPAddress})
		}
	}
	fmt.Println("Droplets status.")
	table.Render()
}
