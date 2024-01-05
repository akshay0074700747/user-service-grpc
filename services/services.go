package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/akshay0074700747/proto-files-for-microservices/pb"
	"github.com/akshay0074700747/user-service/adapters/adapterinterfaces"
	"github.com/akshay0074700747/user-service/entities"
	"github.com/akshay0074700747/user-service/helpers"
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

func (user *UserServiceServer) SignupUser(ctx context.Context, req *pb.SignupUserRequest) (*pb.UserResponce, error) {

	if req.Name == "" {
		return nil, errors.New("the name cannot be empty")
	}
	if req.Email == "" {
		return nil, errors.New("the email cannot be empty")
	}
	if req.Password == "" {
		return nil, errors.New("the password cannot be empty")
	}
	if req.Mobile == "" {
		return nil, errors.New("the mobie cannot be empty")
	}

	password, err := helpers.Hash_pass(req.Password)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	reqq := entities.Clients{
		Name:     req.Name,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Password: password,
	}

	userRes, err := user.Adapter.Signup(reqq)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &pb.UserResponce{
		Id:     uint32(userRes.Id),
		Name:   userRes.Name,
		Email:  userRes.Email,
		Mobile: userRes.Mobile,
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
		Id:     uint32(userRes.Id),
		Name:   userRes.Name,
		Email:  userRes.Email,
		Mobile: userRes.Mobile,
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
		userResponce = append(userResponce, &pb.UserResponce{Id: uint32(userRes.Id), Name: userRes.Name, Email: userRes.Email, Mobile: userRes.Mobile})
	}

	return &pb.AllUsersResponce{
		Users: userResponce,
	}, nil
}

func (user *UserServiceServer) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.UserResponce, error) {

	pass, err := user.Adapter.GetPassByEmail(req.Email, req.IsAdmin, req.IsSuAdmin)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println(req.Password)
	fmt.Println(pass)

	if !req.IsSuAdmin && helpers.VerifyPassword(pass, req.Password) != nil {
		fmt.Println("passwords doesnt match")
		return nil, errors.New("passwors doesnt match")
	}
	if req.IsAdmin {

		userRes, err := user.Adapter.LoginAdmin(req.Email, pass)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		return &pb.UserResponce{
			Id:      uint32(userRes.Id),
			Name:    userRes.Name,
			Email:   userRes.Email,
			Mobile:  userRes.Mobile,
			IsAdmin: true,
		}, nil
	} else if req.IsSuAdmin {

		userRes, err := user.Adapter.LoginSuAdmin(req.Email, pass)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		return &pb.UserResponce{
			Id:        uint32(userRes.Id),
			Name:      userRes.Name,
			Email:     userRes.Email,
			Mobile:    userRes.Mobile,
			IsSuAdmin: true,
		}, nil
	}
	userRes, err := user.Adapter.LoginUser(req.Email, pass)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &pb.UserResponce{
		Id:     uint32(userRes.Id),
		Name:   userRes.Name,
		Email:  userRes.Email,
		Mobile: userRes.Mobile,
	}, nil
}

func (user *UserServiceServer) GetAdmin(ctx context.Context, req *pb.UserRequest) (*pb.UserResponce, error) {

	if req.Id < 1 {
		return nil, errors.New("the id cannot be less than zero")
	}

	userRes, err := user.Adapter.GetAdmin(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.UserResponce{
		Id:      uint32(userRes.Id),
		Name:    userRes.Name,
		Email:   userRes.Email,
		Mobile:  userRes.Mobile,
		IsAdmin: true,
	}, nil
}

func (user *UserServiceServer) GetSuAdmin(ctx context.Context, req *pb.UserRequest) (*pb.UserResponce, error) {

	if req.Id < 1 {
		return nil, errors.New("the id cannot be less than zero")
	}

	userRes, err := user.Adapter.GetSuAdmin(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.UserResponce{
		Id:        uint32(userRes.Id),
		Name:      userRes.Name,
		Email:     userRes.Email,
		Mobile:    userRes.Mobile,
		IsSuAdmin: true,
	}, nil
}

func (user *UserServiceServer) GetAllAdminsResponce(context.Context, *empty.Empty) (*pb.AllUsersResponce, error) {

	users, err := user.Adapter.GetAllAdmins()
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("there are no admins yet")
	}

	var userResponce []*pb.UserResponce
	for _, userRes := range users {
		userResponce = append(userResponce, &pb.UserResponce{Id: uint32(userRes.Id), Name: userRes.Name, Email: userRes.Email, Mobile: userRes.Mobile, IsAdmin: true})
	}

	return &pb.AllUsersResponce{
		Users: userResponce,
	}, nil
}

func (user *UserServiceServer) AddAdmin(ctx context.Context, req *pb.SignupUserRequest) (*pb.UserResponce, error) {

	if req.Name == "" {
		return nil, errors.New("the name cannot be empty")
	}
	if req.Email == "" {
		return nil, errors.New("the email cannot be empty")
	}
	if req.Password == "" {
		return nil, errors.New("the password cannot be empty")
	}
	if req.Mobile == "" {
		return nil, errors.New("the mobie cannot be empty")
	}

	password, err := helpers.Hash_pass(req.Password)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	reqq := entities.Admins{
		Name:     req.Name,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Password: password,
	}

	userRes, err := user.Adapter.AddAdmin(reqq)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponce{
		Id:     uint32(userRes.Id),
		Name:   userRes.Name,
		Email:  userRes.Email,
		Mobile: userRes.Mobile,
	}, nil
}
