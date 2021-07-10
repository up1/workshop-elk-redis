module service01

go 1.16

require (
	common v0.0.1
	github.com/gin-gonic/gin v1.7.2
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.21.0
	go.opentelemetry.io/otel v1.0.0-RC1
	go.opentelemetry.io/otel/exporters/zipkin v1.0.0-RC1
	go.opentelemetry.io/otel/sdk v1.0.0-RC1
)

replace common v0.0.1 => ../common
