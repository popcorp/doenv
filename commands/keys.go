package commands
import (
	"fmt"
	"os"
	"github.com/codegangsta/cli"
	. "github.com/popcorp/doenv/lib"
	"github.com/olekukonko/tablewriter"
)

var KeysCommand = cli.Command{
	Name: "keys",
	Usage:  "Manage the SSH keys used",
	Subcommands: []cli.Command{
		AddSshFingerprintCommand,
		ListSshFingerprintsCommand,
		DelSshFingerprintCommand,
	},
	Action:  func(c *cli.Context) {
		cli.ShowSubcommandHelp(c)
	},
}


var AddSshFingerprintCommand = cli.Command{
	Name: "add",
	Usage:  "Add an SSH key to be used on Droplet creation",
	Aliases: []string{"+"},
	Action:
	func(c *cli.Context) {
		if len(c.Args()) < 1 {
			fmt.Println("To get the SSH fingerprint of your key `id_rsa`, simply type:")
			fmt.Println("ssh-keygen -lf id_rsa.pub")
			fmt.Println("You will get something like 'aa:bb:cc:dd'. Exactly what we need! Set it up using")
			fmt.Println("doenv ssh add 'aa:bb:cc:dd'")
			fmt.Println("!! The key must already be configured on DigitalOcean to be used !!")
			return
		}else {
			AddSshFingerprint(c.Args().First())
			fmt.Println("SSH Fingerprint added! It will be used for the next Droplet creations")
			SaveSettings()
			os.Exit(0)
		}
	},
}

var ListSshFingerprintsCommand = cli.Command{
	Name: "list",
	Usage:  "List SSH keys to be used on Droplet creation",
	Aliases: []string{"ls"},
	Action:
	func(c *cli.Context) {
		if len(GetSettings().SSHFingerprints) == 0 {
			fmt.Println("You don't have any SSH Fingerprints setup to be used.")
			return
		}
		fmt.Println("SSH Fingerprints")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Fingerprint"})
		table.SetBorder(false)
		for _, f := range GetSettings().SSHFingerprints {
			table.Append([]string{f})
		}
		table.Render()
	},
}
var DelSshFingerprintCommand = cli.Command{
	Name: "del",
	Aliases: []string{"delete", "rm", "remove"},
	Usage:  "Stop using this SSH keys on Droplet creation",
	Action:
	func(c *cli.Context) {
		if len(c.Args()) < 1 {
			fmt.Println("To get the SSH fingerprint of your key `id_rsa`, simply type:")
			fmt.Println("ssh-keygen -lf id_rsa.pub")
			fmt.Println("You will get something like 'aa:bb:cc:dd'. Exactly what we need! Set it up using")
			fmt.Println("doenv ssh rm 'aa:bb:cc:dd'")
			fmt.Println("You may also use `doenv ssh list` to see the used fingerprints")
			return
		}else {
			DelSshFingerprint(c.Args().First())
			fmt.Println("SSH Fingerprint removed! :(")
			SaveSettings()
			os.Exit(0)
		}
	},
}
