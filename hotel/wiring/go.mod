module github.com/blueprint-uservices/sosp_tutorial/hotel/wiring

go 1.22.1

require (
	github.com/blueprint-uservices/blueprint/blueprint v0.0.0-20241015110303-ca8bcf724c6d
	github.com/blueprint-uservices/blueprint/plugins v0.0.0-20241015110303-ca8bcf724c6d
)

require github.com/blueprint-uservices/sosp_tutorial/hotel/workload v0.0.0

replace github.com/blueprint-uservices/sosp_tutorial/hotel/workload => ../workload

require (
	github.com/blueprint-uservices/blueprint/examples/dsb_hotel/workflow v0.0.0-20241015110303-ca8bcf724c6d // indirect
	github.com/blueprint-uservices/blueprint/runtime v0.0.0-20240619221802-d064c5861c1e // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/hailocab/go-geoindex v0.0.0-20160127134810-64631bfe9711 // indirect
	github.com/otiai10/copy v1.14.0 // indirect
	go.mongodb.org/mongo-driver v1.15.0 // indirect
	go.opentelemetry.io/otel v1.26.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.26.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.26.0 // indirect
	go.opentelemetry.io/otel/metric v1.26.0 // indirect
	go.opentelemetry.io/otel/sdk v1.26.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.26.0 // indirect
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/tools v0.20.0 // indirect
)
