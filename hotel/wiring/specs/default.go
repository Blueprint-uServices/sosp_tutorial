package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
)

var Default = cmdbuilder.SpecOption{
	Name:        "default",
	Description: "Deploys each service in a separate container with http, uses mongodb as NoSQL database backends, and applies a number of modifiers",
	Build:       makeDefaultSpec,
}

func makeDefaultSpec(spec wiring.WiringSpec) ([]string, error) {
	// TODO: Complete me!
	return []string{}, nil
}
