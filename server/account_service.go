package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	mu    sync.Mutex
	users []*pb.User
}

func newServer() *accountServiceServer {
	return &accountServiceServer{
		users: loadUsers(),
	}
}

func loadUsers() []*pb.User {
	jsonFile, err := os.Open("server/data/users.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users []*pb.User
	json.Unmarshal(byteValue, &users)

	return users
}

func (s *accountServiceServer) QueryUsers(ctx context.Context, in *pb.QueryUsersRequest) (*pb.QueryUsersResponse, error) {
	res := &pb.QueryUsersResponse{Users: s.users}
	log.Printf("QueryUsers: %v > %v", in, res)
	return res, nil
}

func (s *accountServiceServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res := &pb.GetUserResponse{User: s.users[0]}
	log.Printf("GetUser: %v > %v", in, res)
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
