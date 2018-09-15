package account

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"

	pb "github.com/asa-taka/practice-grpc-blog/out/go/account"
	"github.com/asa-taka/practice-grpc-blog/server/services/util"
)

func RegisterNewServer(grpcServer *grpc.Server) {
	pb.RegisterAccountServiceServer(grpcServer, newServer())
}

type accountServiceServer struct {
	mu    sync.Mutex
	users []*pb.User
}

func newServer() *accountServiceServer {
	s := &accountServiceServer{}
	s.loadUsers("server/data/users.json")
	return s
}

func (s *accountServiceServer) loadUsers(filePath string) {
	decoder := util.CreateArrayJSONDecoder(filePath)
	for decoder.More() {
		u := pb.User{}
		if err := jsonpb.UnmarshalNext(decoder, &u); err != nil {
			log.Fatal(err)
		}
		s.users = append(s.users, &u)
	}
}

// rpc Handlers
// ------------

func (s *accountServiceServer) QueryUsers(ctx context.Context, in *pb.QueryUsersRequest) (*pb.QueryUsersResponse, error) {
	res := &pb.QueryUsersResponse{Users: s.users}
	return res, nil
}

func (s *accountServiceServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	u, err := s.getUserByID(in.Id)
	return &pb.GetUserResponse{User: u}, err
}

func (s *accountServiceServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res := &pb.CreateUserResponse{User: in.User} // mock
	return res, nil
}

func (s *accountServiceServer) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	res := &pb.UpdateUserResponse{User: in.User} // mock
	return res, nil
}

func (s *accountServiceServer) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	u, err := s.getUserByID(in.Id) // mock
	return &pb.DeleteUserResponse{User: u}, err
}

// Utilities

func (s *accountServiceServer) getUserByID(id int32) (*pb.User, error) {
	for _, u := range s.users {
		if u.Id == id {
			return u, nil
		}
	}
	return nil, fmt.Errorf("User not found: %v", id)
}
