package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/dsb_hotel/workflow/hotelreservation"
	"github.com/blueprint-uservices/blueprint/plugins/mongodb"
	"github.com/blueprint-uservices/blueprint/plugins/workflow"
)

// initServices adds the basic backends such as caches and databases to the wiring spec. The function also adds the user-defined (aka internal) services to the wiring specification.
// The function returns the array containing the list of names of internal services.
func initServices(spec wiring.WiringSpec) []string {
	// PART 1: TODO --> Complete the initialization of all services
	var services []string

	// Step 1: Define backend databases
	// Step 1a: Define user database
	user_db := mongodb.Container(spec, "user_db")

	// Step 1b: Define recommendations database

	// Step 1c: Define reservations database

	// Step 1d: Define rate database

	// Step 1e: Define profile database

	// Step 2: Define backend caches
	// Step 2a: Define reservations cache

	// Step 2b: Define rate cache

	// Step 2c: Define profile cache

	// Step 3: Define internal services
	// Step 3a: Define user service
	user_service := workflow.Service[hotelreservation.UserService](spec, "user_service", user_db)
	services = append(services, user_service)

	// Step 3b: Define recommendation service

	// Step 3c: Define reservation service

	// Step 3d: Define geo service

	// Step 3e: Define rate service

	// Step 3f: Define profile service

	// Step 3g: Define search service

	// Step 3h: Define frontend service

	return services
}
