package main

import (
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/workflow/workflowspec"
	"github.com/blueprint-uservices/sosp_tutorial/leaf/wiring/specs"
)

func main() {
	// Configure the location of the workflow spec
	workflowspec.AddModule("github.com/blueprint-uservices/blueprint/examples/leaf/workflow")

	name := "leafapp"
	cmdbuilder.MakeAndExecute(
		name,
		specs.Docker,
	)
}
