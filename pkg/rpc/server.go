package rpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"

	pb "go-ops/pkg/rpc/inf"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:50051"
)

type Auth struct {
	AppKey    string
	AppSecret string
}

var OpenTls = true

type service struct {
	auth *Auth
}

func (auth *Auth) Authenticate(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "token auth fail")
	}
	var (
		appKey    string
		appSecret string
	)
	if value, ok := md["app_key"]; ok {
		appKey = value[0]
	}
	if value, ok := md["app_secret"]; ok {
		appSecret = value[0]
	}
	if appKey == auth.AppKey && appSecret == auth.AppSecret {
		return nil
	} else {
		return status.Errorf(codes.Unauthenticated, "wrong token")
	}
}

func (s service) GetUser(ctx context.Context, in *pb.UserRq) (*pb.UserProfile, error) {
	log.Printf("Received: %v", in.Id)
	err := s.auth.Authenticate(ctx)
	if err != nil { // wrong auth
		return nil, err
	}
	return &pb.UserProfile{Name: "Hello ", Email: "hl"}, nil
}

func RunServer() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var grpcServer *grpc.Server

	// open TLS
	if OpenTls {
		creds, err := credentials.NewClientTLSFromFile("config/server/server.pem", "config/server/server.key")
		if err != nil {
			grpclog.Fatalf("Failed to generate credentials %v", err)
		}
		grpcServer = grpc.NewServer(grpc.Creds(creds))
	} else {
		grpcServer = grpc.NewServer()
	}

	// register service
	var auth Auth = Auth{AppKey: "testkey", AppSecret: "test"}
	pb.RegisterDataServer(grpcServer, service{auth: &auth})

	grpcServer.Serve(listen)
}
