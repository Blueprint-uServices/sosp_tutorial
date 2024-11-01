package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/leaf/workflow/leaf"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/goproc"
	"github.com/blueprint-uservices/blueprint/plugins/healthchecker"
	"github.com/blueprint-uservices/blueprint/plugins/http"
	"github.com/blueprint-uservices/blueprint/plugins/linuxcontainer"
	"github.com/blueprint-uservices/blueprint/plugins/memcached"
	"github.com/blueprint-uservices/blueprint/plugins/mongodb"
	"github.com/blueprint-uservices/blueprint/plugins/workflow"
	"github.com/blueprint-uservices/sosp_tutorial/sosp_plugins/hellososp"
)

var Bonus = cmdbuilder.SpecOption{
	Name:        "bonus",
	Description: "Deploys each service in a separate container with http, uses mongodb as NoSQL database backends, and applies a number of modifiers",
	Build:       makeBonusSpec,
}

func makeBonusSpec(spec wiring.WiringSpec) ([]string, error) {
	applyBonusDefaults := func(spec wiring.WiringSpec, serviceName string) string {
		hellososp.AddHelloSOSPStatement(spec, serviceName)
		healthchecker.AddHealthCheckAPI(spec, serviceName)
		http.Deploy(spec, serviceName)
		goproc.Deploy(spec, serviceName)
		return linuxcontainer.Deploy(spec, serviceName)
	}

	leaf_db := mongodb.Container(spec, "leaf_db")
	leaf_cache := memcached.Container(spec, "leaf_cache")
	leaf_service := workflow.Service[*leaf.LeafServiceImpl](spec, "leaf_service", leaf_cache, leaf_db)
	leaf_ctr := applyBonusDefaults(spec, leaf_service)

	nonleaf_service := workflow.Service[leaf.NonLeafService](spec, "nonleaf_service", leaf_service)
	nonleaf_ctr := applyBonusDefaults(spec, nonleaf_service)

	return []string{leaf_ctr, nonleaf_ctr}, nil
}