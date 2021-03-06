package digitalocean

import (
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"log"
)

type stepProvision struct{}

func (*stepProvision) Run(state map[string]interface{}) multistep.StepAction {
	comm := state["communicator"].(packer.Communicator)
	hook := state["hook"].(packer.Hook)
	ui := state["ui"].(packer.Ui)

	log.Println("Running the provision hook")
	if err := hook.Run(packer.HookProvision, ui, comm, nil); err != nil {
		state["error"] = err
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (*stepProvision) Cleanup(map[string]interface{}) {}
