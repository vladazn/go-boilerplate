package server

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/sirupsen/logrus"
	frontoffice "github.com/vladazn/go-boilerplate/api/client"
	grpchandler "github.com/vladazn/go-boilerplate/app/handler/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	address    string
	grpcServer *grpc.Server
}

func NewServer(
	config *GrpcServerConfigs,
	authHandler *grpchandler.AuthHandler,
	userHandler *grpchandler.UserHandler,
) *Server {
	server := &Server{
		address: config.Address(),
	}

	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}

	server.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(
		logging.UnaryServerInterceptor(InterceptorLogger(logrus.New()), opts...),
	))

	server.grpcServer.RegisterService(&frontoffice.AuthService_ServiceDesc, authHandler)
	server.grpcServer.RegisterService(&frontoffice.UserService_ServiceDesc, userHandler)

	reflection.Register(server.grpcServer)

	return server
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	err = s.grpcServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}
