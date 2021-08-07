package main

import (
	"fmt"
	"net/http"
	"os"

	endpoints "github.com/AnesBenmerzoug/udemy-api-data/internal/endpoints"
	"github.com/joho/godotenv"
)

var CLIENT_ID, CLIENT_SECRET string

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	CLIENT_ID = os.Getenv("CLIENT_ID")
	CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
}

func main() {
	client := &http.Client{}
	courses, err := endpoints.GetCourses(client, CLIENT_ID, CLIENT_SECRET)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(courses)
}
