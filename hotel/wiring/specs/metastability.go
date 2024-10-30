package specs

import (
	"fmt"

	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/goproc"
	"github.com/blueprint-uservices/blueprint/plugins/http"
	"github.com/blueprint-uservices/blueprint/plugins/linuxcontainer"
	"github.com/blueprint-uservices/blueprint/plugins/workload"
	"github.com/blueprint-uservices/sosp_tutorial/hotel/workload/workloadgen"
)

var Metastability = cmdbuilder.SpecOption{
	Name:        "metastability",
	Description: "Deploys each service in a separate docker container with http, uses mongodb as NoSQL database backends, uses memcached as cache backends, and applies the retries and timeout plugins to each service",
	Build:       makeMetastabilitySpec,
}

// applyMetastabilityScaffolding applies five plugins to the service:
// (i) timeout plugin --- adds timeout to calls to this service
// (ii) retry plugin --- calls to the service will have retries
// (iii) http plugin --- deploys the service as an independent http server
// (iv) goproc plugin --- deploys the service as an independent OS process
// (v) linuxcontainer plugin --- creates a docker container for the previously defined process.
func applyMetastabilityScaffolding(spec wiring.WiringSpec, serviceName string) string {

	// Part 3: TODO ---> Complete the function to apply the retry and timeout plugins

	// Step 1: Apply the retry plugin

	// Step 2: Apply the timeout plugin

	procName := fmt.Sprintf("%s_process", serviceName)
	ctrName := fmt.Sprintf("%s_container", serviceName)
	http.Deploy(spec, serviceName)
	goproc.CreateProcess(spec, procName, serviceName)
	return linuxcontainer.CreateContainer(spec, ctrName, procName)
}

func makeMetastabilitySpec(spec wiring.WiringSpec) ([]string, error) {
	// Initialize services
	services := initServices(spec)

	cntrs := []string{}
	for _, service := range services {
		cntr := applyMetastabilityScaffolding(spec, service)
		cntrs = append(cntrs, cntr)
	}

	// Use the complex generator
	wlgen := workload.Generator[workloadgen.ComplexWorkload](spec, "wlgen", "frontend_service")
	cntrs = append(cntrs, wlgen)

	return cntrs, nil
}
