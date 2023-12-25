package test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/akshay0074700747/proto-files-for-microservices/pb"
	mock_adapterinterfaces "github.com/akshay0074700747/user-service/adapters/mock"
	"github.com/akshay0074700747/user-service/entities"
	"github.com/akshay0074700747/user-service/services"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	adapter := mock_adapterinterfaces.NewMockAdapterInterface(ctrl)

	tests := []struct {
		name      string
		mockFunc  func(user entities.Users) (entities.Users, error)
		request   *pb.AddUserRequest
		wantError bool
		wantUser  *pb.UserResponce
	}{
		{
			name: "Success",
			mockFunc: func(user entities.Users) (entities.Users, error) {
				return entities.Users{Id: 1, Name: user.Name, IsAdmin: user.IsAdmin}, nil
			},
			request:   &pb.AddUserRequest{Name: "John", IsAdmin: false},
			wantError: false,
			wantUser:  &pb.UserResponce{Id: 1, Name: "John", IsAdmin: false},
		},
		{
			name: "Adapter Error",
			mockFunc: func(user entities.Users) (entities.Users, error) {
				return entities.Users{}, errors.New("adapter error")
			},
			request:   &pb.AddUserRequest{Name: "Abcd", IsAdmin: true},
			wantError: true,
			wantUser:  nil,
		},
		{
			name: "Success",
			mockFunc: func(user entities.Users) (entities.Users, error) {
				return entities.Users{Id: 1, Name: user.Name, IsAdmin: user.IsAdmin}, nil
			},
			request:   &pb.AddUserRequest{Name: "Akshay", IsAdmin: true},
			wantError: false,
			wantUser:  &pb.UserResponce{Id: 1, Name: "Akshay", IsAdmin: true},
		},
	}

	for _, tt := range tests {

		adapter.EXPECT().Adduser(gomock.Any()).DoAndReturn(tt.mockFunc).AnyTimes().Times(1)
		usrService := services.NewUserServiceServer(adapter)

		user, err := usrService.AddUser(context.TODO(), tt.request)
		if tt.wantError {
			fmt.Println("addUser fail")
			assert.Error(t, err)
			assert.Nil(t, user)
		} else {
			fmt.Println("addUser success")
			assert.NoError(t, err)
			assert.NotNil(t, user)
			assert.Equal(t, tt.wantUser, user)
		}
	}
}

func TestGetUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	adapter := mock_adapterinterfaces.NewMockAdapterInterface(ctrl)
	tests := []struct {
		name      string
		mockFunc  func(id uint) (entities.Users, error)
		request   *pb.UserRequest
		wantError bool
		wantUser  *pb.UserResponce
	}{
		{
			name: "Success",
			mockFunc: func(id uint) (entities.Users, error) {
				return entities.Users{Id: id, Name: "Akshay", IsAdmin: true}, nil
			},
			request:   &pb.UserRequest{Id: 1},
			wantError: false,
			wantUser:  &pb.UserResponce{Id: 1, Name: "Akshay", IsAdmin: true},
		},
		{
			name: "Failure",
			mockFunc: func(id uint) (entities.Users, error) {
				return entities.Users{}, fmt.Errorf("this testCase is a failure")
			},
			request:   &pb.UserRequest{Id: 100},
			wantError: true,
			wantUser:  nil,
		},
		{
			name: "Success",
			mockFunc: func(id uint) (entities.Users, error) {
				return entities.Users{Id: id, Name: "Frank", IsAdmin: false}, nil
			},
			request:   &pb.UserRequest{Id: 2},
			wantError: false,
			wantUser:  &pb.UserResponce{Id: 2, Name: "Frank", IsAdmin: false},
		},
	}

	for _, test := range tests {

		adapter.EXPECT().GetUser(gomock.Any()).DoAndReturn(test.mockFunc).AnyTimes().Times(1)
		usrService := services.NewUserServiceServer(adapter)

		user, err := usrService.GetUser(context.TODO(), test.request)
		if test.wantError {
			fmt.Println("getUser fail")
			assert.Error(t, err)
			assert.Nil(t, user)
		} else {
			fmt.Println("getUser success")
			assert.NoError(t, err)
			assert.NotNil(t, user)
			assert.Equal(t, test.request.Id, user.Id)
		}
	}
}

func TestGetAllUsersResponce(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	adapter := mock_adapterinterfaces.NewMockAdapterInterface(ctrl)
	tests := []struct {
		name      string
		request   *empty.Empty
		mockfunc  func() ([]entities.Users, error)
		expected  *pb.AllUsersResponce
		wantError bool
	}{
		{
			name:    "Success",
			request: &empty.Empty{},
			mockfunc: func() ([]entities.Users, error) {
				return []entities.Users{
					{Id: 1, Name: "Akshay", IsAdmin: true},
					{Id: 2, Name: "Frank", IsAdmin: false},
				}, nil
			},
			expected: &pb.AllUsersResponce{Users: []*pb.UserResponce{
				{Id: 1, Name: "Akshay", IsAdmin: true},
				{Id: 2, Name: "Frank", IsAdmin: false},
			}},
			wantError: false,
		},
		{
			name:    "Failure",
			request: &empty.Empty{},
			mockfunc: func() ([]entities.Users, error) {
				return []entities.Users{}, fmt.Errorf("this is a failed case")
			},
			expected:  nil,
			wantError: true,
		},
	}

	for _, test := range tests {

		adapter.EXPECT().GetAllUsers().DoAndReturn(test.mockfunc).AnyTimes().Times(1)
		usrService := services.NewUserServiceServer(adapter)

		users, err := usrService.GetAllUsersResponce(context.TODO(), test.request)
		if test.wantError {
			fmt.Println("getAllUsers fail")
			assert.Error(t, err)
			assert.Nil(t, users)
		} else {
			fmt.Println("getAllUsers success")
			assert.NoError(t, err)
			assert.NotNil(t, users)
			assert.Equal(t, test.expected, users)
		}
	}

}
