package test

import (
	"os"
	"testing"

	"github.com/akshay0074700747/user-service/db"
	"github.com/joho/godotenv"
)

func TestDatabaseConnection(t *testing.T) {

	if err := godotenv.Load("../cmd/.env"); err != nil {
		t.Fatal(err.Error())
	}

	addr := os.Getenv("TEST_DATABASE_ADDR")

	_, err := db.InitDB(addr)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log("Database connection successfull...")
}
