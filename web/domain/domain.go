package domain

import (
	"github.com/synoday/gateway/client"
	"github.com/synoday/gateway/web/router"
)

// Domain is standard interface for all synoday service domains.
type Domain interface {
	// PlugRoute registers domain routes.
	PlugRoute(r *router.Router)

	// PlugClient attach gRPC client service to each domain.
	PlugClient(k *client.ServiceClient)
}
