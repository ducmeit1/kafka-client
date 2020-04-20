package transports

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"path"
)

func NewRouter(pathPrefix string, transports []Route) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		io.WriteString(w, `{"status": "ok"}`)
	})

	for _, t := range transports {
		r.Handle(path.Join(pathPrefix, t.Path), t.Handler).Methods(t.Method)
	}

	return r
}
