package client

import (
	"flag"
	"os"
	"path/filepath"

	taskpb "github.com/synoday/golang/protogen/task"
	userpb "github.com/synoday/golang/protogen/user"
)

var (
	serverName = flag.String("server-name", "dev.synoday.com", "Server Name")
	caCert     = flag.String("ca-cert", envDir("ca.pem"), "Trusted CA certificate.")
	tlsCert    = flag.String("tls-cert", envDir("client.pem"), "TLS server certificate.")
	tlsKey     = flag.String("tls-key", envDir("client-key.pem"), "TLS server key.")
)

func envDir(path string) string {
	return filepath.Join(os.Getenv("GOPATH"), "src", "github.com/synoday/gateway", "creds/certs", path)
}

// ServiceClient is combined synoday gRPC service client.
type ServiceClient struct {
	userpb.UserServiceClient
	taskpb.TaskServiceClient
}

// New creates new synoday service client.
func New(opts ...Option) *ServiceClient {
	flag.Parse()

	client := new(ServiceClient)

	for _, o := range opts {
		o(client)
	}
	return client
}
