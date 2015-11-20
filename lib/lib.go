package lib

import (
	"github.com/digitalocean/godo"
	"time"
	"golang.org/x/oauth2"
	"github.com/tucnak/store"
	"log"
	"os"
	"os/exec"
	"fmt"
)

type Settings struct {
	PersonnalApiToken string
	SSHFingerprints   []string
}

var settings Settings
var token string

func GetSettings() Settings {
	return settings
}

func LoadSettings() {
	err := store.Load("preferences.json", &settings)
	if err != nil {
		log.Printf("failed to load preferences: %s\n", err)
		os.Exit(1)
	}
}

func GetFingerprints() []godo.DropletCreateSSHKey {
	var fingerprints []godo.DropletCreateSSHKey
	for _, f := range GetSettings().SSHFingerprints {
		fingerprints = append(fingerprints, godo.DropletCreateSSHKey{
			Fingerprint: f,
		})
	}
	return fingerprints
}

func SetPersonnalApiToken(PersonnalApiToken string) {
	settings.PersonnalApiToken = PersonnalApiToken
}

func AddSshFingerprint(Fingerprint string) {
	settings.SSHFingerprints = append(settings.SSHFingerprints, Fingerprint)
}

func DelSshFingerprint(Fingerprint string) {
	var fingerprints []string
	for _, f := range GetSettings().SSHFingerprints {
		if f != Fingerprint {
			fingerprints = append(fingerprints, f)
		}
	}
	settings.SSHFingerprints = fingerprints
}

func SaveSettings() {
	err := store.Save("preferences.json", &settings)
	if err != nil {
		log.Printf("failed to save preferences: %s\n", err)
		os.Exit(1)
	}
}
func UseClientToken(Token string) {
	token = Token
}
func DropletByName(Name string) (*godo.Droplet, error) {
	opt := &godo.ListOptions{}
	for {
		droplets, resp, err := GetClient().Droplets.List(opt)
		if err != nil {
			return nil, err
		}

		// append the current page's droplets to our list
		for _, d := range droplets {
			if d.Name == Name {
				return &d, nil
			}
		}

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}
	return nil, nil
}

func ImageByName(Name string) (*godo.Image, error) {
	opt := &godo.ListOptions{}
	for {
		droplets, resp, err := GetClient().Images.List(opt)
		if err != nil {
			return nil, err
		}

		// append the current page's droplets to our list
		for _, d := range droplets {
			if d.Name == Name {
				return &d, nil
			}
		}

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}
	return nil, nil
}

func DropletsList() ([]godo.Droplet, error) {
	// create a list to hold our droplets
	list := []godo.Droplet{}

	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		droplets, resp, err := GetClient().Droplets.List(opt)
		if err != nil {
			return nil, err
		}

		// append the current page's droplets to our list
		for _, d := range droplets {
			list = append(list, d)
		}

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}
	return list, nil
}

func ImagesList() ([]godo.Image, error) {
	// create a list to hold our droplets
	list := []godo.Image{}

	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		images, resp, err := GetClient().Images.List(opt)
		if err != nil {
			return nil, err
		}

		// append the current page's droplets to our list
		for _, i := range images {
			list = append(list, i)
		}

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}
	return list, nil
}

func ActionsList() ([]godo.Action, error) {
	// create a list to hold our droplets
	list := []godo.Action{}

	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		actions, resp, err := GetClient().Actions.List(opt)
		if err != nil {
			return nil, err
		}

		// append the current page's droplets to our list
		for _, a := range actions {
			list = append(list, a)
		}

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}
	return list, nil
}

func WaitAction(actionID int) (error) {
	var err error
	var action *godo.Action
	for {
		if err != nil {
			return err
		}
		if action != nil {
			fmt.Println("...", action.Status)
			if action.Status == "completed" {
				return nil
			}
		}
		time.Sleep(5 * time.Second)
		action, _, err = GetClient().Actions.Get(actionID)
	}
}
func WaitStatus(dropletID int, Status string) {
	var d *godo.Droplet
	for {
		time.Sleep(1 * time.Second) // Arbitrary time, because Action completed ≠ Droplet down
		d, _, _ = GetClient().Droplets.Get(dropletID)
		if d.Status == Status {
			return
		}
		fmt.Println("...", d.Status)
	}
}

func GetClient() *godo.Client {
	token := &oauth2.Token{AccessToken: token}
	t := oauth2.StaticTokenSource(token)

	oauthClient := oauth2.NewClient(oauth2.NoContext, t)
	return godo.NewClient(oauthClient)
}

func SshDroplet(d *godo.Droplet) {
	binary, lookErr := exec.LookPath("ssh")
	if lookErr != nil {
		panic(lookErr)
	}
	ip := fmt.Sprintf("root@%s", d.Networks.V4[0].IPAddress)
	cmd := exec.Command(binary, ip)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type ByCost []godo.Size

func (s ByCost) Len() int {
	return len(s)
}
func (s ByCost) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByCost) Less(i, j int) bool {
	return s[i].PriceMonthly < s[i].PriceMonthly
}
