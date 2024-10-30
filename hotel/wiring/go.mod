module github.com/blueprint-uservices/sosp_tutorial/hotel/wiring

go 1.22.1

require (
	github.com/blueprint-uservices/blueprint/blueprint v0.0.0-20241029180336-5902a9377ed2
	github.com/blueprint-uservices/blueprint/plugins v0.0.0-20241029180336-5902a9377ed2
)

require (
	github.com/blueprint-uservices/blueprint/examples/dsb_hotel/workflow v0.0.0-20241029180336-5902a9377ed2
	github.com/blueprint-uservices/sosp_tutorial/hotel/workload v0.0.0
)

replace github.com/blueprint-uservices/sosp_tutorial/hotel/workload => ../workload

require (
	github.com/blueprint-uservices/blueprint/runtime v0.0.0-20241029180336-5902a9377ed2 // indirect
	github.com/bradfitz/gomemcache v0.0.0-20230905024940-24af94b03874 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hailocab/go-geoindex v0.0.0-20160127134810-64631bfe9711 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/otiai10/copy v1.14.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	go.mongodb.org/mongo-driver v1.17.1 // indirect
	go.opentelemetry.io/otel v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.31.0 // indirect
	go.opentelemetry.io/otel/metric v1.31.0 // indirect
	go.opentelemetry.io/otel/sdk v1.31.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.31.0 // indirect
	go.opentelemetry.io/otel/trace v1.31.0 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/exp v0.0.0-20241009180824-f66d83c29e7c // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	golang.org/x/tools v0.26.0 // indirect
)
