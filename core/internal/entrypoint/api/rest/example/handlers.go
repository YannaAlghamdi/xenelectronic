package example

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/xenelectronic/core/internal/component/example"
)

func getNoAuthHandler(exampleService example.Service) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.Debugf("GET no auth handler called", mux.Vars(req))
		responses.WriteOK(res)
	}
}
