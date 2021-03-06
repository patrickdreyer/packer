package digitalocean

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
	"github.com/hashicorp/packer/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer/packer-plugin-sdk/packer"
)

type stepDropletInfo struct{}

func (s *stepDropletInfo) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	client := state.Get("client").(*godo.Client)
	ui := state.Get("ui").(packersdk.Ui)
	c := state.Get("config").(*Config)
	dropletID := state.Get("droplet_id").(int)

	ui.Say("Waiting for droplet to become active...")

	err := waitForDropletState("active", dropletID, client, c.StateTimeout)
	if err != nil {
		err := fmt.Errorf("Error waiting for droplet to become active: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	// Set the IP on the state for later
	droplet, _, err := client.Droplets.Get(context.TODO(), dropletID)
	if err != nil {
		err := fmt.Errorf("Error retrieving droplet: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	// Verify we have an IPv4 address
	invalid := droplet.Networks == nil ||
		len(droplet.Networks.V4) == 0
	if invalid {
		err := fmt.Errorf("IPv4 address not found for droplet")
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	// Find the ip address which will be used by communicator
	foundNetwork := false
	for _, network := range droplet.Networks.V4 {
		if (c.ConnectWithPrivateIP && network.Type == "private") ||
			(!(c.ConnectWithPrivateIP) && network.Type == "public") {
			state.Put("droplet_ip", network.IPAddress)
			foundNetwork = true
			break
		}
	}
	if !foundNetwork {
		err := fmt.Errorf("Count not find a public IPv4 address for this droplet")
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *stepDropletInfo) Cleanup(state multistep.StateBag) {
	// no cleanup
}
