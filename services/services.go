package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/akshay0074700747/proto-files-for-microservices/pb"
	"github.com/akshay0074700747/user-service/adapters/adapterinterfaces"
	"github.com/akshay0074700747/user-service/entities"
	"github.com/golang/protobuf/ptypes/empty"
)

type UserServiceServer struct {
	Adapter adapterinterfaces.AdapterInterface
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(adapter adapterinterfaces.AdapterInterface) *UserServiceServer {
	return &UserServiceServer{
		Adapter: adapter,
	}
}

func (user *UserServiceServer) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.UserResponce, error) {

	if req.Name == "" {
		return nil, errors.New("the name cannot be empty")
	}

	reqq := entities.Users{
		Name:    req.Name,
		IsAdmin: req.IsAdmin,
	}
	userRes, err := user.Adapter.Adduser(reqq)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponce{
		Id:      uint32(userRes.Id),
		Name:    userRes.Name,
		IsAdmin: userRes.IsAdmin,
	}, nil
}

func (user *UserServiceServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponce, error) {

	if req.Id < 1 {
		return nil, errors.New("the id cannot be less than zero")
	}

	userRes, err := user.Adapter.GetUser(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.UserResponce{
		Id:      uint32(userRes.Id),
		Name:    userRes.Name,
		IsAdmin: userRes.IsAdmin,
	}, nil
}

func (user *UserServiceServer) GetAllUsersResponce(context.Context, *empty.Empty) (*pb.AllUsersResponce, error) {

	users, err := user.Adapter.GetAllUsers()
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("there are no users yet")
	}

	var userResponce []*pb.UserResponce
	for _, userRes := range users {
		userResponce = append(userResponce, &pb.UserResponce{Id: uint32(userRes.Id), Name: userRes.Name, IsAdmin: userRes.IsAdmin})
	}

	return &pb.AllUsersResponce{
		Users: userResponce,
	}, nil
}
