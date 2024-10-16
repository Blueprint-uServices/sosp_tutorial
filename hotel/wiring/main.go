package main

import (
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/workflow/workflowspec"
	"github.com/blueprint-uservices/sosp_tutorial/hotel/wiring/specs"
)

func main() {
	// Configure the location of the workflow spec
	workflowspec.AddModule("github.com/blueprint-uservices/blueprint/examples/dsb_hotel/workflow")

	name := "hotel_reservation"
	cmdbuilder.MakeAndExecute(
		name,
		specs.Default,
	)
}
