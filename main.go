package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/acornies/faas-provider-example/handlers"
	hclog "github.com/hashicorp/go-hclog"
	bootstrap "github.com/openfaas/faas-provider"
	"github.com/openfaas/faas-provider/types"
)

var (
	port        = flag.Int("port", 8081, "Port to bind the server to")
	loggerLevel = flag.String("logger_level", "INFO", "Log output level INFO | ERROR | DEBUG | TRACE")
)

func main() {

	flag.Parse()

	config := &types.FaaSConfig{
		TCPPort: port,
	}

	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "faas-empty-provider",
		Level:      hclog.LevelFromString(*loggerLevel),
		JSONFormat: false,
		Output:     os.Stdout,
	})

	handlers := &types.FaaSHandlers{

		FunctionReader: func(w http.ResponseWriter, r *http.Request) {

			log := logger.Named("FunctionReader")

			exampleResponse := []byte(`[{"name":"exampleFunc", "replicas": 3}]`)
			log.Info("Reader called", "canned response", string(exampleResponse))

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(exampleResponse)
		},

		DeployHandler: handlers.Deploy(logger),

		FunctionProxy: func(w http.ResponseWriter, r *http.Request) {
			// TODO: implement
		},
		ReplicaReader: func(w http.ResponseWriter, r *http.Request) {
			// TODO: implement
		},
		ReplicaUpdater: func(w http.ResponseWriter, r *http.Request) {
			// TODO: implement
		},
		DeleteHandler: func(w http.ResponseWriter, r *http.Request) {
			// TODO: implement
		},
	}

	logger.Info("Started example provider", "port", *config.TCPPort)

	bootstrap.Serve(handlers, config)
}
