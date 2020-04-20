package transports

import "net/http"

type Route struct {
	Path    string
	Method  string
	Handler http.Handler
}

type Transport struct {
	PathPrefix string
	Routes     []Route
}
