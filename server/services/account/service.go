package account

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"google.golang.org/grpc"

	pb "github.com/asa-taka/practice-grpc-blog/out/go/account"
)

func RegisterNewServer(grpcServer *grpc.Server) {
	pb.RegisterAccountServiceServer(grpcServer, newServer())
}

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

// rpc Handlers
// ------------

func (s *accountServiceServer) QueryUsers(ctx context.Context, in *pb.QueryUsersRequest) (*pb.QueryUsersResponse, error) {
	res := &pb.QueryUsersResponse{Users: s.users}
	return res, nil
}

func (s *accountServiceServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res := &pb.GetUserResponse{User: s.users[0]}
	return res, nil
}

func (s *accountServiceServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res := &pb.CreateUserResponse{}
	return res, nil
}

func (s *accountServiceServer) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	res := &pb.UpdateUserResponse{}
	return res, nil
}

func (s *accountServiceServer) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	res := &pb.DeleteUserResponse{}
	return res, nil
}
