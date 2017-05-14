package web

import (
	"fmt"
	"log"

	"github.com/knq/envcfg"

	"github.com/synoday/gateway/client"
	"github.com/synoday/gateway/web/domain"
	"github.com/synoday/gateway/web/domain/task"
	"github.com/synoday/gateway/web/domain/user"
	"github.com/synoday/gateway/web/router"
)

// Gateway contains information needed to start synoday web api gateway.
type Gateway struct {
	config *envcfg.Envcfg
	router *router.Router
	client *client.ServiceClient
}

// Run starts synoday gateway web api gateway server.
func (g *Gateway) Run() {
	g.bootstrap()

	addr := fmt.Sprintf(":%s", g.config.GetString("app.port"))
	log.Printf("Synoday API Gateways started on: %s \n", addr)
	g.router.Run(addr)
}

// bootstrap prepare and do preprocess works before actually running synoday web api gateway.
func (g *Gateway) bootstrap() {
	var err error

	// setup app config.
	g.config, err = envcfg.New()
	if err != nil {
		log.Fatal(err)
	}

	// setup app routes.
	g.router = router.New()

	// setup grpc client connections.
	g.client = client.New(
		client.UserService(g.config.GetString("usersvc.host"), g.config.GetString("usersvc.port")),
		client.TaskService(g.config.GetString("tasksvc.host"), g.config.GetString("tasksvc.port")),
	)

	// register domain services.
	g.plug(
		user.Domain,
		task.Domain,
	)
}

// plug registers all synodays service domains.
func (g *Gateway) plug(domains ...domain.Domain) {
	for _, domain := range domains {
		domain.PlugRoute(g.router)
		domain.PlugClient(g.client)
	}
}
