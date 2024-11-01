package hellososp

import (
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/blueprint-uservices/blueprint/blueprint/pkg/coreplugins/service"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/ir"
	"github.com/blueprint-uservices/blueprint/plugins/golang"
	"github.com/blueprint-uservices/blueprint/plugins/golang/gocode"
	"github.com/blueprint-uservices/blueprint/plugins/golang/gogen"
)

// Blueprint IRNode for representing the wrapper node that instruments every client-side method in the node that gets wrapped
type HelloInstrumentClientWrapper struct {
	golang.Service
	golang.GeneratesFuncs
	golang.Instantiable

	// TODO: Add other relevant fields
}

// Implements ir.IRNode
func (node *HelloInstrumentClientWrapper) ImplementsGolangNode() {}

// Implements ir.IRNode
func (node *HelloInstrumentClientWrapper) Name() string {
	// TODO: Implement this function. Add any relevant fields to the HelloInstrumentClientWrapper as needed.
	return ""
}

// Implements ir.IRNode
func (node *HelloInstrumentClientWrapper) String() string {
	// TODO: Implement this function. Add any relevant fields to the HelloInstrumentClientWrapper as needed.
	return ""
}

// Implements golang.ProvidesInterface
func (node *HelloInstrumentClientWrapper) AddInterfaces(builder golang.ModuleBuilder) error {
	// TODO: Implement this function. Add any relevant fields to the HelloInstrumentClientWrapper as needed.
	return nil
}

func newHelloInstrumentClientWrapper(name string, client ir.IRNode) (*HelloInstrumentClientWrapper, error) {
	// TODO: Implement this function. Add any relevant fields to the HelloInstrumentClientWrapper as needed.
	return nil, nil
}

// Implements service.ServiceNode
func (node *HelloInstrumentClientWrapper) GetInterface(ctx ir.BuildContext) (service.ServiceInterface, error) {
	// TODO: Implement this function. Add any relevant fields to the HelloInstrumentClientWrapper as needed.
	return nil, nil
}

// Implements golang.GeneratesFuncs
// More info: https://github.com/Blueprint-uServices/blueprint/tree/main/plugins/golang#GeneratesFuncs
func (node *HelloInstrumentClientWrapper) GenerateFuncs(builder golang.ModuleBuilder) error {
	// TODO: Implement this function. Add any relevant fields to the HelloInstrumentClientWrapper as needed.
	return nil
}

// Implements golang.Instantiable
// More Info: https://github.com/Blueprint-uServices/blueprint/tree/main/plugins/golang#type-instantiable
func (node *HelloInstrumentClientWrapper) AddInstantiation(builder golang.NamespaceBuilder) error {
	// TODO: Implement this function. Add any relevant fields to the HelloInstrumentClientWrapper as needed.
	return nil
}

func generateClientInstrumentHandler(builder golang.ModuleBuilder, wrapped *gocode.ServiceInterface, outputPackage string) error {
	pkg, err := builder.CreatePackage(outputPackage)
	if err != nil {
		return err
	}

	server := &templateArgs{
		Package:   pkg,
		Service:   wrapped,
		Iface:     wrapped,
		Name:      wrapped.BaseName + "_TutorialInstrumentClientWrapper",
		IfaceName: wrapped.Name,
		Imports:   gogen.NewImports(pkg.Name),
	}

	server.Imports.AddPackages("context", "log")

	slog.Info(fmt.Sprintf("Generating %v/%v", server.Package.PackageName, wrapped.BaseName+"_TutorialInstrumentClientWrapper"))
	outputFile := filepath.Join(server.Package.Path, wrapped.BaseName+"_TutorialInstrumentClientWrapper.go")
	return gogen.ExecuteTemplateToFile("Tutorial", clientInstrumentTemplate, server, outputFile)
}
