package main

import (
	"github.com/codegangsta/cli"
	"os"
	"fmt"
	"github.com/tucnak/store"
	. "github.com/popcorp/doenv/lib"
	. "github.com/popcorp/doenv/commands"
)

func main() {
	store.SetApplicationName("doenv")
	LoadSettings()
	var settings = GetSettings()

	app := cli.NewApp()
	app.Name = "doenv"
	app.Usage = "It's like virtualenv, but powered by DigitalOcean"
	app.Version = "1.0.0"
	app.Author = "PunKeel <punkeel@me.com>"

	if settings.PersonnalApiToken == "" {
		app.Action = func(c *cli.Context) {
			fmt.Println("Well, looks like you haven't defined your Personnal API token yet.")
			fmt.Println("To grab one, simply go to https://cloud.digitalocean.com/settings/applications and generate a new token")
			fmt.Println("Then, asusming the token is '421gtp', simply run `doenv set-token 421gtp`")
			os.Exit(1)
		}

		app.Commands = []cli.Command{SetTokenCommand}
		app.Run(os.Args)
		os.Exit(1)
	}else {
		UseClientToken(settings.PersonnalApiToken)
	}


	app.Action = func(c *cli.Context) {
		if len(c.Args()) == 1 {
			GuessCommand(c.Args().First())
		}
		cli.ShowAppHelp(c)
	}

	app.Commands = []cli.Command{
		ListCommand,
		CreateCommand,
		EraseCommand,

		RestoreCommand,
		FreezeCommand,
		SnapshotCommand,

		SshCommand,
		KeysCommand,
		SetTokenCommand,

		/**
		 * Disabled commands: they don't fit into this project.
		 * Kept because they're still cool & might be re-enabled lated
		 */
		// StatusCommand,
		// StartCommand,
		// StopCommand,
	}

	app.Run(os.Args)
}
