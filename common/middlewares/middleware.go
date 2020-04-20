package middlewares

import (
	"github.com/ducmeit1/kafka-client/common/transports"
	"github.com/urfave/negroni"
	"net/http"
)

type Middlewares struct {
	M *negroni.Negroni
}

func NewMiddlewares() *Middlewares {
	return &Middlewares{M: negroni.New()}
}

func (m *Middlewares) AddGlobalMiddlewares() {
	m.M.Use(negroni.NewRecovery())
	m.M.Use(negroni.NewLogger())
	m.M.Use(NewDefaultCors())
}

func (m *Middlewares) AddMiddlewares(handlers ...http.Handler) {
	for _, h := range handlers {
		m.M.UseHandler(h)
	}
}

func (m *Middlewares) AddRoutes(pathPrefix string, routes []transports.Route) {
	r := transports.NewRouter(pathPrefix, routes)
	m.M.UseHandler(r)
}

func (m *Middlewares) GetHandlers() http.Handler {
	return m.M
}
