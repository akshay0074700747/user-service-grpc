package test

import (
	"context"
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

func TestSignup(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	adapter := mock_adapterinterfaces.NewMockAdapterInterface(ctrl)

	tests := []struct {
		name      string
		mockFunc  func(user entities.Clients) (entities.Clients, error)
		request   *pb.SignupUserRequest
		wantError bool
		wantUser  *pb.UserResponce
	}{
		{
			name: "Success",
			mockFunc: func(user entities.Clients) (entities.Clients, error) {
				return entities.Clients{Id: 1, Name: user.Name, Email: user.Email, Mobile: user.Mobile}, nil
			},
			request:   &pb.SignupUserRequest{Name: "John", Email: "john@gmail.com", Mobile: "9435467231", Password: "qw##w23Aw"},
			wantError: false,
			wantUser:  &pb.UserResponce{Id: 1, Name: "John", Email: "john@gmail.com", Mobile: "9435467231", IsAdmin: false, IsSuAdmin: false},
		},
		{
			name: "Success",
			mockFunc: func(user entities.Clients) (entities.Clients, error) {
				return entities.Clients{Id: 2, Name: user.Name, Email: user.Email, Mobile: user.Mobile}, nil
			},
			request:   &pb.SignupUserRequest{Name: "Akshay", Email: "akshay@gmail.com", Mobile: "9435467231", Password: "qw##w23Aw"},
			wantError: false,
			wantUser:  &pb.UserResponce{Id: 2, Name: "Akshay", Email: "akshay@gmail.com", Mobile: "9435467231", IsAdmin: false, IsSuAdmin: false},
		},
		{
			name: "Failure",
			mockFunc: func(user entities.Clients) (entities.Clients, error) {
				return entities.Clients{}, fmt.Errorf("here occures a error")
			},
			request:   &pb.SignupUserRequest{Name: "Akshay", Email: "akshay@gmail.com", Mobile: "9435467231", Password: "qw##w23Aw"},
			wantError: true,
			wantUser:  nil,
		},
	}

	for _, tt := range tests {

		adapter.EXPECT().Signup(gomock.Any()).DoAndReturn(tt.mockFunc).AnyTimes().Times(1)
		usrService := services.NewUserServiceServer(adapter)

		user, err := usrService.SignupUser(context.TODO(), tt.request)
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
		mockFunc  func(id uint) (entities.Clients, error)
		request   *pb.UserRequest
		wantError bool
		wantUser  *pb.UserResponce
	}{
		{
			name: "Success",
			mockFunc: func(id uint) (entities.Clients, error) {
				return entities.Clients{Id: id, Name: "Akshay", Email: "akshay@gmail.com", Mobile: "9345679876"}, nil
			},
			request:   &pb.UserRequest{Id: 1},
			wantError: false,
			wantUser:  &pb.UserResponce{Id: 1, Name: "Akshay", Email: "akshay@gmail.com", Mobile: "9345679876"},
		},
		{
			name: "Failure",
			mockFunc: func(id uint) (entities.Clients, error) {
				return entities.Clients{}, fmt.Errorf("its a test case")
			},
			request:   &pb.UserRequest{Id: 100},
			wantError: true,
			wantUser:  nil,
		},
		{
			name: "Success",
			mockFunc: func(id uint) (entities.Clients, error) {
				return entities.Clients{Id: id, Name: "Frank", Email: "frank@gmail.com", Mobile: "97654345678"}, nil
			},
			request:   &pb.UserRequest{Id: 2},
			wantError: false,
			wantUser:  &pb.UserResponce{Id: 2, Name: "Frank", Email: "frank@gmail.com", Mobile: "97654345678"},
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
		mockfunc  func() ([]entities.Clients, error)
		expected  *pb.AllUsersResponce
		wantError bool
	}{
		{
			name:    "Success",
			request: &empty.Empty{},
			mockfunc: func() ([]entities.Clients, error) {
				return []entities.Clients{
					{Id: 1, Name: "Akshay", Email: "akshay@gmail.com", Mobile: "9765434567"},
					{Id: 2, Name: "Frank", Email: "frank@gmail.com", Mobile: "987654567"},
				}, nil
			},
			expected: &pb.AllUsersResponce{Users: []*pb.UserResponce{
				{Id: 1, Name: "Akshay", Email: "akshay@gmail.com", Mobile: "9765434567"},
				{Id: 2, Name: "Frank", Email: "frank@gmail.com", Mobile: "987654567"},
			}},
			wantError: false,
		},
		{
			name:    "Failure",
			request: &empty.Empty{},
			mockfunc: func() ([]entities.Clients, error) {
				return []entities.Clients{}, fmt.Errorf("this is a failed case")
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
