# Adding a new Plugin

In the bonus part of the tutorial, we will be adding a new plugin that instruments every method for a service (both on the client-side and the server-side) to print a `Hello, SOSP` message to standard output every time a method is executed. 

We have provided the plugin implementation for instrumenting the server side. Your goal is to complete the implementation for instrumenting the client side. Once you have completed the plugin, we will then test out the newly completed plugin with the leaf application.

## Instrumenting the Client Side

To implement the client side portion of the plugin, you need to complete 3 steps: (i) Implementing a client-side template which will represent the modified code that will contain the instrumentation; (ii) Implementing the client IR node so that Blueprint's compiler can generate the modified code from the IR node; and (iii) Implementing the wiring API so that this plugin can be used in wiring specifications.

### Implementing the client-side template

To implement the client-side template, complete the template attributed to the variable `clientInstrumentTemplate` in [templates.go](./sosp_plugins/hellososp/templates.go) file.

The template is generated using the `templateArgs` struct defined on lines 9-17 of the `templates.go` file. 

> Hint: Use the serverInstrumentTemplate as a guide for implementing the various steps in the clientInstrumentTemplate.

### Implementing the IR client node

To implement the client-side IR node, complete the TODO list in the [ir_client.go](./sosp_plugins/hellososp/ir_client.go) file.

We suggest first completing the implementation of the function `newHelloInstrumentClientWrapper` that constructs the IR Node as well as the implementation of the struct `HelloInstrumentClientWrapper` before finishing other TODOs.

> Hint: Use the HelloInstrumentServerWrapper defined in [ir_server.go](./sosp_plugins/hellososp/ir_server.go) as a guide for implementing various steps in the serverInstrumentTemplate.

### Implementing the IR wiring

To implement the IR wiring, complete the `AddHelloSOSPStatement` function in the [wiring.go](./sosp_plugins/hellososp/wiring.go) file.

To complete this, you will need add the client-side wrapper to the client-side node chain of the desired service. Wrappers can be added to the node chains of the desired service as follows:

```go
import "github.com/blueprint-uservices/blueprint/blueprint/pkg/coreplugins/pointer"

func addToChain(serviceName string, serverWrapper string, clientWrapper string) {
    // Get the pointer for the serviceName that has access
    ptr := pointer.GetPointer(spec, serviceName)
    serverNext := ptr.AddDstModifier(spec, serverWrapper) // Add serverWrapper to the server-side node chain
    clientNext := ptr.AddSrcModifier(spec, clientWrapper) // Add clientWrapper to the client-side node chain

    // Use serverNext to instantiate the IR node.

    // Use clientNext to instantiate the IR node.
}
```

Once you have added the client-side wrapper to the node chain, you must then use the return value `clientNext` which is the IR node that will be wrapped by the new IR Node. The following snippet shows how to add a new IR node to the wiring spec for the server side.

```go
spec.Define(wrapper_name, &HelloInstrumentServerWrapper{}, func(ns wiring.Namespace) (ir.IRNode, error) {
		// Get the IRNode that will be wrapped by HelloWrapper
		var server golang.Service
		if err := ns.Get(serverNext, &server); err != nil {
			return nil, blueprint.Errorf("Tutorial Plugin %s expected %s to be a golang.Service, but encountered %s", wrapper_name, serverNext, err)
		}

		// Instantiate the IRNode
		return newHelloInstrumentServerWrapper(wrapper_name, server)
	})
```

## Running Leaf application

To run the leaf application with our new plugin, no extra code needs to be written. We have already provided a wiring specification that generates the implementation with our plugin.

### Generating the application

To generate the application, execute the following steps:

```bash
cd leaf/wiring
go run main.go -w bonus -o build
```

### Building the application

We will use `docker compose` to build the containers and then eventually deploy them.

To build the application, execute the following steps:

```bash
# Assuming you are in the leaf/wiring folder
cd build/docker # Switch to the docker directory where the containers are
cp ../../.env . # Copy the provided .env file from the wiring folder. Note that we are not using the generated .env file for simplicity.
docker compose build # Build the containers
```

> Note: executing docker commands may require sudo access

### Running the application

To launch the built containers, execute the following command:

```bash
docker compose up -d
```

### Testing the application

To test the application, you can execute the following `curl` command:

```bash
curl http://localhost:12348/Hello?a=5
```

The Hello API simply returns a counter value indicating the number of previous requests processed.

Thus, you should see the following output:

```bash
{"Ret0":0}
```

On subsequent executions of the curl command, the counter will increase monotonically.

### Checking logs

To check if the message was printed, open the docker logs for the non-leaf container. To do so, execute the following command:

```bash
docker logs nonleaf_ctr-1
```

### Stopping the application

To stop the launched containers, execute the following command:

```bash
docker compose down
```

## Conclusion

Congratulations on finishing all parts of the SOSP tutorial! You are now officially a semi-pro in Blueprint fundamentals!