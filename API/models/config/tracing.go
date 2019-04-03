package config

import (
	"accountBook/models/beans"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"sourcegraph.com/sourcegraph/appdash"
	appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"
	"sourcegraph.com/sourcegraph/appdash/traceapp"
)

func startAppdashServer(appdashPort int) {
	store := appdash.NewMemoryStore()

	// Listen on any available TCP port locally.
	l, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 0})
	if err != nil {
		log.Error("%s", err)
	}
	collectorPort := l.Addr().(*net.TCPAddr).Port

	// Start an Appdash collection server that will listen for spans and
	// annotations and add them to the local collector (stored in-memory).
	cs := appdash.NewServer(l, appdash.NewLocalCollector(store))
	go cs.Start()

	// Print the URL at which the web UI will be running.
	appdashURLStr := fmt.Sprintf("http://localhost:%d", appdashPort)
	appdashURL, err := url.Parse(appdashURLStr)
	if err != nil {
		log.Error("Error parsing %s: %s", appdashURLStr, err)
	}
	fmt.Printf("To see your traces, go to %s/traces\n", appdashURL)

	// Start the web UI in a separate goroutine.
	tapp, err := traceapp.New(nil, appdashURL)
	if err != nil {
		log.Error("Error creating traceapp: %v", err)
	}
	tapp.Store = store
	tapp.Queryer = store
	go func() {
		log.Error("%s", http.ListenAndServe(fmt.Sprintf(":%d", appdashPort), tapp))
	}()

	tracer := appdashot.NewTracer(appdash.NewRemoteCollector(fmt.Sprintf(":%d", collectorPort)))
	opentracing.InitGlobalTracer(tracer)
}

func startZipkinServer(zipkinURL string) {
	collector, err := zipkin.NewHTTPCollector(zipkinURL, zipkin.HTTPTimeout(1*time.Minute))
	if err != nil {
		log.Error("unable to create Zipkin HTTP collector: %+v\n", err)
		beans.Tracing = false
		return
	}

	// Create our recorder.
	recorder := zipkin.NewRecorder(collector, true, "0.0.0.0:0", "accountBook")

	// Create our tracer.
	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(true),
		zipkin.TraceID128Bit(true),
	)
	if err != nil {
		log.Error("unable to create Zipkin tracer: %+v\n", err)
		beans.Tracing = false
		return
	}
	// Explicitly set our tracer to be the default tracer.
	opentracing.InitGlobalTracer(tracer)
}
