package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/openfaas/faas/gateway/requests"
)

var ()

func Deploy(logger hclog.Logger) http.HandlerFunc {
	logger.Named("DeployHandler")

	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		body, _ := ioutil.ReadAll(r.Body)

		req := requests.CreateFunctionRequest{}
		err := json.Unmarshal(body, &req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		logger.Info("Function create/deploy called", "name", req.Service)
		// write any code
	}
}
