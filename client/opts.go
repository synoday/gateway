package client

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	taskpb "github.com/synoday/golang/protogen/task"
	userpb "github.com/synoday/golang/protogen/user"
)

// Option is synoday service client option.
type Option func(*ServiceClient)

// UserService is an option that creates new user service client connection.
func UserService(host, port string) Option {
	return func(s *ServiceClient) {
		log.Println(fmt.Sprintf("%s:%s", host, port))
		s.UserServiceClient = userpb.NewUserServiceClient(mustDial(fmt.Sprintf("%s:%s", host, port)))
	}
}

// TaskService is an option that creates new task service client connection.
func TaskService(host, port string) Option {
	return func(s *ServiceClient) {
		log.Println(fmt.Sprintf("%s:%s", host, port))
		s.TaskServiceClient = taskpb.NewTaskServiceClient(mustDial(fmt.Sprintf("%s:%s", host, port)))
	}
}

// mustDial ensures a tcp connection to specified address.
func mustDial(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	return conn
}
