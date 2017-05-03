package task

import (
	"github.com/synoday/gateway/client"
	"github.com/synoday/gateway/web/router"
)

var synodayClient *client.ServiceClient

// Domain expose task domain implementation.
var Domain domain

// domain is task domain implementation.
type domain struct{}

// PlugRoute registers tasl domain routes.
func (d domain) PlugRoute(route *router.Router) {
	for _, r := range routes {
		route.R.
			Methods(r.Method).
			Path(r.Path).
			HandlerFunc(r.Handler)
	}
}

// PlugClient attach gRPC client service to user domain.
func (d domain) PlugClient(k *client.ServiceClient) {
	synodayClient = k
}
