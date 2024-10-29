package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/dsb_hotel/workflow/hotelreservation"
	"github.com/blueprint-uservices/blueprint/plugins/memcached"
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
	recommendations_db := mongodb.Container(spec, "recomd_db")

	// Step 1c: Define reservations database
	reserv_db := mongodb.Container(spec, "reserv_db")

	// Step 1d: Define rate database
	rate_db := mongodb.Container(spec, "rate_db")

	// Step 1e: Define profile database
	profile_db := mongodb.Container(spec, "profile_db")

	// Step 1f: Define geo database
	geo_db := mongodb.Container(spec, "geo_db")

	// Step 2: Define backend caches
	// Step 2a: Define reservations cache
	reserv_cache := memcached.Container(spec, "reserv_cache")

	// Step 2b: Define rate cache
	rate_cache := memcached.Container(spec, "rate_cache")

	// Step 2c: Define profile cache
	profile_cache := memcached.Container(spec, "profile_cache")

	// Step 3: Define internal services
	// Step 3a: Define user service
	user_service := workflow.Service[hotelreservation.UserService](spec, "user_service", user_db)
	services = append(services, user_service)

	// Step 3b: Define recommendation service
	recomd_service := workflow.Service[hotelreservation.RecommendationService](spec, "recomd_service", recommendations_db)
	services = append(services, recomd_service)

	// Step 3c: Define reservation service
	reserv_service := workflow.Service[hotelreservation.ReservationService](spec, "reserv_service", reserv_cache, reserv_db)
	services = append(services, reserv_service)

	// Step 3d: Define geo service
	geo_service := workflow.Service[hotelreservation.GeoService](spec, "geo_service", geo_db)
	services = append(services, geo_service)

	// Step 3e: Define rate service
	rate_service := workflow.Service[hotelreservation.RateService](spec, "rate_service", rate_cache, rate_db)
	services = append(services, rate_service)

	// Step 3f: Define profile service
	profile_service := workflow.Service[hotelreservation.ProfileService](spec, "profile_service", profile_cache, profile_db)
	services = append(services, profile_service)

	// Step 3g: Define search service
	search_service := workflow.Service[hotelreservation.SearchService](spec, "search_service", geo_service, rate_service)
	services = append(services, search_service)

	// Step 3h: Define frontend service
	frontend_service := workflow.Service[hotelreservation.FrontEndService](spec, "frontend_service", search_service, profile_service, recomd_service, user_service, reserv_service)
	services = append(services, frontend_service)

	return services
}
