package hellososp

import (
	"log/slog"

	"github.com/blueprint-uservices/blueprint/blueprint/pkg/blueprint"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/coreplugins/pointer"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/ir"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/plugins/golang"
)

// [AddHelloSOSPStatement] can be called from the wiring specification to add a "Hello, SOSP attendees!" logging statement to every method in both the server and client side code of Service with `serviceName`.
func AddHelloSOSPStatement(spec wiring.WiringSpec, serviceName string) {
	// Define the names for the wrapper nodes we are adding to the Blueprint IR
	wrapper_name := serviceName + ".hello.sosp.server"
	client_wrapper_name := serviceName + ".hello.client.server"

	// Get the pointer for the serviceName to ensure that the newly defined wrapper IR node will be attached to the node chain of the desired service
	ptr := pointer.GetPointer(spec, serviceName)
	if ptr == nil {
		slog.Error("Unable to add instrument " + serviceName + " as it is not a pointer. Did you forget to define " + serviceName + "? You can define serviceName using `workflow.Service`")
		return
	}

	// Attach the Hello wrapper node to the server-side node chain of the desired service
	serverNext := ptr.AddDstModifier(spec, wrapper_name)

	// Define the IRNode for the wrapper node and add it to the wiring specification
	spec.Define(wrapper_name, &HelloInstrumentServerWrapper{}, func(ns wiring.Namespace) (ir.IRNode, error) {
		// Get the IRNode that will be wrapped by HelloWrapper
		var server golang.Service
		if err := ns.Get(serverNext, &server); err != nil {
			return nil, blueprint.Errorf("Tutorial Plugin %s expected %s to be a golang.Service, but encountered %s", wrapper_name, serverNext, err)
		}

		// Instantiate the IRNode
		return newHelloInstrumentServerWrapper(wrapper_name, server)
	})

	// Step 1: Attach the Hello wrapper node to the client-side node chain of the desired service

	// Step 2: Define the IRNode for the client wrapper node and add it to the wiring specification

	// Step 3: Once the above two steps are completed, remove the following statement
	_ = client_wrapper_name

}
