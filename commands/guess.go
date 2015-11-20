package commands
import "fmt"
import . "github.com/popcorp/doenv/lib"

func GuessCommand(DropletName string) {
	fmt.Printf("Unknown command %s .. Trying some magic trick\n", DropletName)
	d, _ := DropletByName(DropletName)
	if d == nil {
		restoreDroplet(DropletName, DropletName)
	}
	if d.Status == "off" {
		fmt.Println("Droplet %s is powered off, starting it ...\n", DropletName)
		action, _, err := GetClient().DropletActions.PowerOn(d.ID)
		if err != nil {
			panic(err)
		}
		WaitAction(action.ID)
		d, _, _ = GetClient().Droplets.Get(d.ID)
	}

	if d.Status == "active" {
		fmt.Println("Droplet is running, let's hack into it ! :D")
		SshDroplet(d)
	}
}
