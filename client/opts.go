package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	taskpb "github.com/synoday/golang/protogen/task"
	userpb "github.com/synoday/golang/protogen/user"
)

// Option is synoday service client option.
type Option func(*ServiceClient)

// UserService is an option that creates new user service client connection.
func UserService(host, port string) Option {
	return func(s *ServiceClient) {
		s.UserServiceClient = userpb.NewUserServiceClient(mustDial(fmt.Sprintf("%s:%s", host, port)))
	}
}

// TaskService is an option that creates new task service client connection.
func TaskService(host, port string) Option {
	return func(s *ServiceClient) {
		s.TaskServiceClient = taskpb.NewTaskServiceClient(mustDial(fmt.Sprintf("%s:%s", host, port)))
	}
}

// mustDial ensures a tcp connection to specified address.
func mustDial(addr string) *grpc.ClientConn {
	certificate, err := tls.LoadX509KeyPair(*tlsCert, *tlsKey)
	if err != nil {
		log.Fatalf("could not load server key pair: %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(*caCert)
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}

	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append client certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   *serverName,
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	return conn
}
