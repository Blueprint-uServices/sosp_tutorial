package hellososp

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"reflect"

	"github.com/blueprint-uservices/blueprint/blueprint/pkg/blueprint"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/coreplugins/service"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/ir"
	"github.com/blueprint-uservices/blueprint/plugins/golang"
	"github.com/blueprint-uservices/blueprint/plugins/golang/gocode"
	"github.com/blueprint-uservices/blueprint/plugins/golang/gogen"
)

// Blueprint IRNode for representing the wrapper node that instruments every server-side method in the node that gets wrapped
type HelloInstrumentServerWrapper struct {
	golang.Service
	golang.GeneratesFuncs
	golang.Instantiable

	InstanceName string
	Wrapped      golang.Service

	outputPackage string
}

// Implements ir.IRNode
func (node *HelloInstrumentServerWrapper) ImplementsGolangNode() {}

// Implements ir.IRNode
func (node *HelloInstrumentServerWrapper) Name() string {
	return node.InstanceName
}

// Implements ir.IRNode
func (node *HelloInstrumentServerWrapper) String() string {
	return node.Name() + " = HelloInstrumentServerWrapper(" + node.Wrapped.Name() + ")"
}

// Implements golang.ProvidesInterface
func (node *HelloInstrumentServerWrapper) AddInterfaces(builder golang.ModuleBuilder) error {
	return node.Wrapped.AddInterfaces(builder)
}

func newHelloInstrumentServerWrapper(name string, server ir.IRNode) (*HelloInstrumentServerWrapper, error) {
	serverNode, ok := server.(golang.Service)
	if !ok {
		return nil, blueprint.Errorf("tutorial server wrapper requires %s to be a golang service but got %s", server.Name(), reflect.TypeOf(server).String())
	}

	node := &HelloInstrumentServerWrapper{}
	node.InstanceName = name
	node.Wrapped = serverNode
	node.outputPackage = "sosp"

	return node, nil
}

// Implements service.ServiceNode
func (node *HelloInstrumentServerWrapper) GetInterface(ctx ir.BuildContext) (service.ServiceInterface, error) {
	return node.Wrapped.GetInterface(ctx)
}

// Implements golang.GeneratesFuncs
func (node *HelloInstrumentServerWrapper) GenerateFuncs(builder golang.ModuleBuilder) error {
	iface, err := golang.GetGoInterface(builder, node)
	if err != nil {
		return err
	}
	return generateServerInstrumentHandler(builder, iface, node.outputPackage)
}

// Implements golang.Instantiable
func (node *HelloInstrumentServerWrapper) AddInstantiation(builder golang.NamespaceBuilder) error {
	if builder.Visited(node.InstanceName) {
		return nil
	}

	iface, err := golang.GetGoInterface(builder, node.Wrapped)
	if err != nil {
		return err
	}

	constructor := &gocode.Constructor{
		Package: builder.Module().Info().Name + "/" + node.outputPackage,
		Func: gocode.Func{
			Name: fmt.Sprintf("New_%v_TutorialInstrumentServerWrapper", iface.BaseName),
			Arguments: []gocode.Variable{
				{Name: "ctx", Type: &gocode.UserType{Package: "context", Name: "Context"}},
				{Name: "service", Type: iface},
			},
		},
	}

	return builder.DeclareConstructor(node.InstanceName, constructor, []ir.IRNode{node.Wrapped})
}

func generateServerInstrumentHandler(builder golang.ModuleBuilder, wrapped *gocode.ServiceInterface, outputPackage string) error {
	pkg, err := builder.CreatePackage(outputPackage)
	if err != nil {
		return err
	}

	server := &templateArgs{
		Package:   pkg,
		Service:   wrapped,
		Iface:     wrapped,
		Name:      wrapped.BaseName + "_TutorialInstrumentServerWrapper",
		IfaceName: wrapped.Name,
		Imports:   gogen.NewImports(pkg.Name),
	}

	server.Imports.AddPackages("context", "log")

	slog.Info(fmt.Sprintf("Generating %v/%v", server.Package.PackageName, wrapped.BaseName+"_TutorialInstrumentServerWrapper"))
	outputFile := filepath.Join(server.Package.Path, wrapped.BaseName+"_TutorialInstrumentServerWrapper.go")
	return gogen.ExecuteTemplateToFile("Tutorial", serverInstrumentTemplate, server, outputFile)
}
