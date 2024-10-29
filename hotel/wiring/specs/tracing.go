package specs

import (
	"fmt"

	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/goproc"
	"github.com/blueprint-uservices/blueprint/plugins/http"
	"github.com/blueprint-uservices/blueprint/plugins/jaeger"
	"github.com/blueprint-uservices/blueprint/plugins/linuxcontainer"
	"github.com/blueprint-uservices/blueprint/plugins/opentelemetry"
	"github.com/blueprint-uservices/blueprint/plugins/workload"
	"github.com/blueprint-uservices/sosp_tutorial/hotel/workload/workloadgen"
)

var Tracing = cmdbuilder.SpecOption{
	Name:        "tracing",
	Description: "Deploys each service in a separate docker container with http, uses mongodb as NoSQL database backends, uses memcached as cache backends, and applies the opentelemetry plugin to each service",
	Build:       makeTracingSpec,
}

// applyDefaultScaffolding applies four plugins to the service:
// (i) opentelemetry plugin --- instruments each service to generate opentelemetry compatible spans
// (ii) http plugin --- deploys each service as an independent http server
// (iii) goproc plugin --- deploys each server as an independent OS process
// (iv) linuxcontainer plugin --- creates a docker container for each of the previously defined process
func applyTracingScaffolding(spec wiring.WiringSpec, serviceName string, collectorName string) string {
	// Part 2: TODO ---> Complete the function to apply the necessary plugins

	procName := fmt.Sprintf("%s_process", serviceName)
	ctrName := fmt.Sprintf("%s_container", serviceName)

	// Step 1: Apply the opentelemetry plugin
	opentelemetry.Instrument(spec, serviceName, collectorName)

	// Step 2: Apply the http plugin
	http.Deploy(spec, serviceName)

	// Step 3: Apply the goproc plugin
	goproc.CreateProcess(spec, procName, serviceName)

	// Step 4: Apply the linuxcontainer plugin
	// Step 5: Replace the serviceName with the containerName in the return value
	return linuxcontainer.CreateContainer(spec, ctrName, procName)
}

func makeTracingSpec(spec wiring.WiringSpec) ([]string, error) {
	// Initialize services
	services := initServices(spec)

	// Part 2: TODO ---> Initialize a jaeger collector that can be used for collecting traces
	collector := jaeger.Collector(spec, "jaeger") // TODO: Replace this with a call to the jaeger plugin

	cntrs := []string{}
	for _, service := range services {
		cntr := applyTracingScaffolding(spec, service, collector)
		cntrs = append(cntrs, cntr)
	}

	// Use a simple workload generator to test that this works. We have written a simple workload generator for you that you can use here.

	wlgen := workload.Generator[workloadgen.SimpleWorkload](spec, "wlgen", "frontend_service")
	cntrs = append(cntrs, wlgen)
	return cntrs, nil
}
