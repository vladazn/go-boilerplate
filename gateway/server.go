package gateway

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	frontoffice "github.com/vladazn/go-boilerplate/api/gateway"
	"github.com/vladazn/go-boilerplate/pkg/jwt"
	"github.com/vladazn/go-boilerplate/pkg/toolkit/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"net/http"
)

type GatewayServer struct {
	server *http.Server
}

func NewServer(config *GatewayServerConfig, jwtGenerator *jwt.JwtGenerator) (*GatewayServer, error) {
	mux := runtime.NewServeMux(
		//runtime.WithErrorHandler()
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			header := request.Header.Get("Authorization")
			if header != "" {
				params, _ := jwtGenerator.ParseAuthToken(header)
				if params != nil {
					if params.IsValid() {
						auth.WithUserId(ctx, params.UserId)
						return metadata.Pairs("userid", params.UserId.String())
					}
				}
			}

			return nil
		}),
	)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := frontoffice.RegisterAuthServiceHandlerFromEndpoint(
		context.TODO(),
		mux,
		config.Services.AuthService,
		opts,
	)
	if err != nil {
		return nil, err
	}

	err = frontoffice.RegisterUserServiceHandlerFromEndpoint(
		context.TODO(),
		mux,
		config.Services.UserService,
		opts,
	)
	if err != nil {
		return nil, err
	}

	handler := cors.AllowAll().Handler(mux)

	return &GatewayServer{
		server: &http.Server{
			Addr:    config.Address(),
			Handler: handler,
		},
	}, nil
}

func (s *GatewayServer) Start() error {
	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *GatewayServer) Stop() error {
	return s.server.Shutdown(context.TODO())
}
