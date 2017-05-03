package client

import (
	taskpb "github.com/synoday/golang/protogen/task"
	userpb "github.com/synoday/golang/protogen/user"
)

// ServiceClient is combined synoday gRPC service client.
type ServiceClient struct {
	userpb.UserServiceClient
	taskpb.TaskServiceClient
}

// New creates new synoday service client.
func New(opts ...Option) *ServiceClient {
	client := new(ServiceClient)

	for _, o := range opts {
		o(client)
	}
	return client
}
