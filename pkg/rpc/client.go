package rpc

import (
	"context"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"log"
	"strconv"
	"time"

	pb "go-ops/pkg/rpc/inf"
	"google.golang.org/grpc"
)

const (
	defaultName = "world"
	Address = "127.0.0.1:50051"
	AppKey       = "testkey1"
	AppSecret      = "test"
)

var openTls = false // 是否开启SSL认证

/* 实现认证接口
type PerRPCCredentials interface {
    GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
    RequireTransportSecurity() bool
}
*/

// customCredential 自定义认证
type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"app_key":  AppKey,
		"app_secret": AppSecret,
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	return openTls
}

func RunClient() {
	var opts []grpc.DialOption
	// open TLS
	if openTls {
		creds, err := credentials.NewClientTLSFromFile("config/server/server.pem", "go-ops")
		if err != nil {
			grpclog.Fatal("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	// add PerRPCCredentials
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	// init connection
	conn, err := grpc.Dial(Address, opts...)
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	// init client
	c := pb.NewDataClient(conn)
	var userId string = "5"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	id, _ := strconv.Atoi(userId)

	// call method
	r, err := c.GetUser(ctx, &pb.UserRq{Id: int32(id)})
	if err != nil {
		log.Fatalf("could not getuser: %v", err)
	}
	log.Printf("Get user: %s", r.Name, r.Email)
}
