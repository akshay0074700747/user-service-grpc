package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/akshay0074700747/proto-files-for-microservices/pb"
	initializer "github.com/akshay0074700747/user-service/Initializer"
	"github.com/akshay0074700747/user-service/db"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestAddUserIntegration(t *testing.T) {
	fmt.Println("Add user integration called first...")
	if err := godotenv.Load("../cmd/.env"); err != nil {
		log.Fatal(err.Error())
	}

	addr := os.Getenv("TEST_DATABASE_ADDR")

	db, err := db.InitDB(addr)
	if err != nil {
		log.Fatal(err.Error())
	}

	usrService := initializer.Initialize(db)

	tests := []struct {
		name      string
		request   *pb.AddUserRequest
		wantError bool
		wantUser  *pb.UserResponce
	}{
		{
			name:      "Success",
			request:   &pb.AddUserRequest{Name: "Trevor", IsAdmin: false},
			wantError: false,
			wantUser:  &pb.UserResponce{Id: 1, Name: "Trevor", IsAdmin: false},
		},
		{
			name:      "Failure",
			request:   &pb.AddUserRequest{},
			wantError: true,
			wantUser:  nil,
		},
		{
			name:      "Success",
			request:   &pb.AddUserRequest{Name: "Messi", IsAdmin: true},
			wantError: false,
			wantUser:  &pb.UserResponce{Id: 2, Name: "Messi", IsAdmin: true},
		},
	}

	for _, test := range tests {
		responce, err := usrService.AddUser(context.TODO(), test.request)
		if test.wantError {
			assert.Error(t, err)
			assert.Nil(t, responce)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, responce)
			assert.Equal(t, test.wantUser, responce)
		}
	}
}

func TestGetUserIntegration(t *testing.T) {
	fmt.Println("Get user integration called second...")
	if err := godotenv.Load("../cmd/.env"); err != nil {
		log.Fatal(err.Error())
	}

	addr := os.Getenv("TEST_DATABASE_ADDR")

	db, err := db.InitDB(addr)
	if err != nil {
		log.Fatal(err.Error())
	}

	usrService := initializer.Initialize(db)

	tests := []struct {
		name      string
		request   *pb.UserRequest
		wantError bool
		wantUser  *pb.UserResponce
	}{
		{
			name:      "Success",
			request:   &pb.UserRequest{Id: 1},
			wantError: false,
			wantUser:  &pb.UserResponce{Id: 1, Name: "Trevor", IsAdmin: false},
		},
		{
			name:      "Failure",
			request:   &pb.UserRequest{Id: 0},
			wantError: true,
			wantUser:  nil,
		},
		{
			name:      "Success",
			request:   &pb.UserRequest{Id: 2},
			wantError: false,
			wantUser:  &pb.UserResponce{Id: 2, Name: "Messi", IsAdmin: true},
		},
	}

	for _, test := range tests {
		responce, err := usrService.GetUser(context.TODO(), test.request)
		if test.wantError {
			assert.Error(t, err)
			assert.Nil(t, responce)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, responce)
			assert.Equal(t, test.wantUser, responce)
		}
	}
}

func TestGetAllUsersIntegration(t *testing.T) {
	fmt.Println("Get all users integration called third...")
	if err := godotenv.Load("../cmd/.env"); err != nil {
		log.Fatal(err.Error())
	}

	addr := os.Getenv("TEST_DATABASE_ADDR")

	db, err := db.InitDB(addr)

	defer func() {
		db.Exec("drop table users")
	}()

	if err != nil {
		log.Fatal(err.Error())
	}

	usrService := initializer.Initialize(db)

	tests := []struct {
		name      string
		request   *empty.Empty
		expected  *pb.AllUsersResponce
		wantError bool
	}{
		{
			name:    "Success",
			request: &empty.Empty{},
			expected: &pb.AllUsersResponce{Users: []*pb.UserResponce{
				{Id: 1, Name: "Trevor", IsAdmin: false},
				{Id: 2, Name: "Messi", IsAdmin: true},
			}},
			wantError: false,
		},
		// {
		// 	name:    "Failure",
		// 	request: &empty.Empty{},
		// 	expected:  nil,
		// 	wantError: true,
		// },
	}

	for _, test := range tests {

		responce, err := usrService.GetAllUsersResponce(context.TODO(), test.request)
		if test.wantError {
			assert.Error(t, err)
			assert.Nil(t, responce)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, responce)
			assert.Equal(t, test.expected, responce)
		}
	}
}
