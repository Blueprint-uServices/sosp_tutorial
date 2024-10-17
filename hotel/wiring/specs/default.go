package specs

import (
	"fmt"

	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/goproc"
	"github.com/blueprint-uservices/blueprint/plugins/http"
	"github.com/blueprint-uservices/blueprint/plugins/linuxcontainer"
)

var Default = cmdbuilder.SpecOption{
	Name:        "default",
	Description: "Deploys each service in a separate container with http, uses mongodb as NoSQL database backends, and does not apply any additional modifiers/plugins",
	Build:       makeDefaultSpec,
}

// applyDefaultScaffolding applies three plugins to the service:
// (i) http plugin --- deploys each service as an independent http server
// (ii) goproc plugin --- deploys each server as an independent OS process
// (iii) linuxcontainer plugin --- creates a docker container for each of the previously defined process
func applyDefaultScaffolding(spec wiring.WiringSpec, serviceName string) string {
	procName := fmt.Sprintf("%s_process", serviceName)
	ctrName := fmt.Sprintf("%s_container", serviceName)
	http.Deploy(spec, serviceName)
	goproc.CreateProcess(spec, procName, serviceName)
	return linuxcontainer.CreateContainer(spec, ctrName, procName)
}

func makeDefaultSpec(spec wiring.WiringSpec) ([]string, error) {
	// Initialize services
	services := initServices(spec)
	// Apply default plugins to the services -- deploy each service as a http server running as an independent linux process running in a separate container
	cntrs := []string{}
	for _, service := range services {
		cntr := applyDefaultScaffolding(spec, service)
		cntrs = append(cntrs, cntr)
	}
	return cntrs, nil
}
