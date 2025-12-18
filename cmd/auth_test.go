package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"shorten/internal/auth"
	"shorten/internal/user"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "a@ya.ru",
		Password: "$2a$10$cHsy6q1TH6Glf4eBf/xptOGKKfml82AgCwcz1KNMOusXl44uz1yM6",
		Name:     "Вася",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "a@ya.ru").
		Delete(&user.User{})
}

func TestLoginSuccess(t *testing.T) {
	// Prepare
	db := initDb()
	initData(db)
	defer removeData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a@ya.ru",
		Password: "password",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)

	if err != nil {
		t.Fatal(err)
	}

	if resData.Token == "" {
		t.Fatal("Expected token, got empty")
	}

	// Clean
	removeData(db)
}

func TestLoginFail(t *testing.T) {
	// Prepare
	db := initDb()
	initData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a@a.com",
		Password: "1",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 401 {
		t.Fatalf("Expected status code 401, got %d", res.StatusCode)
	}

	// Clean
	removeData(db)
}
