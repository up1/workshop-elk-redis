package main

import (
	"common"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("service01")
var logger = common.NewLogger("service01")

func main(){
	traceProvider := initTracer()
	defer func() {
		if err := traceProvider.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	// Start server
	r := gin.New()
	r.Use(otelgin.Middleware("service01"))
	r.GET("/users/:id", func(c *gin.Context) {

		span := oteltrace.SpanFromContext(c.Request.Context())
		logger.WithTracing(span.SpanContext().TraceID().String())

		id := c.Param("id")
		name := getUser(c, id)
		c.JSON(200, gin.H{
			"name": name,
			"id":   id,
		})
	})
	_ = r.Run(":8080")
}


func initTracer() *trace.TracerProvider {
	exporter, _ := zipkin.New(
		"http://localhost:9411/api/v2/spans",
	)

	traceProvider := trace.NewTracerProvider(
		trace.WithSampler(trace.TraceIDRatioBased(1)), // All (0-1)
		trace.WithBatcher(exporter,
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
			trace.WithBatchTimeout(trace.DefaultBatchTimeout),
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
		),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("service01"),
				attribute.String("exporter", "zipkin"),
			),
		),
	)

	otel.SetTracerProvider(traceProvider)
	return traceProvider
}

func getUser(c *gin.Context, id string) string {
	_, span := tracer.Start(c.Request.Context(), "getUser", oteltrace.WithAttributes(attribute.String("id", id)))
	defer span.End()
	if id == "123" {
		return "name for tester"
	}
	return "unknown"
}
