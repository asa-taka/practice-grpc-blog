package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	"github.com/improbable-eng/grpc-web/go/grpcweb"

	pb "github.com/asa-taka/practice-grpc-blog/out/go/account"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type accountServiceServer struct {
	mu sync.Mutex
}

func newServer() *accountServiceServer {
	s := &accountServiceServer{}
	return s
}

func (s *accountServiceServer) QueryUsers(ctx context.Context, in *pb.QueryUsersRequest) (*pb.QueryUsersResponse, error) {
	log.Printf("QueryUsers: %v", in)
	res := &pb.QueryUsersResponse{}
	return res, nil
}

func (s *accountServiceServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("GetUser: %v", in)
	res := &pb.GetUserResponse{}
	return res, nil
}

func (s *accountServiceServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Printf("CreateUser: %v", in)
	res := &pb.CreateUserResponse{}
	return res, nil
}

func (s *accountServiceServer) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	log.Printf("UpdateUser: %v", in)
	res := &pb.UpdateUserResponse{}
	return res, nil
}

func (s *accountServiceServer) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	log.Printf("DeleteUser: %v", in)
	res := &pb.DeleteUserResponse{}
	return res, nil
}

func main() {
	flag.Parse()
	// lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	grpcServer := grpc.NewServer()

	pb.RegisterAccountServiceServer(grpcServer, newServer())
	reflection.Register(grpcServer)

	// grpcServer.Serve(lis)

	// HTTP Server for gRPC Web

	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: http.HandlerFunc(handler),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("failed starting http server: %v", err)
	}
}
