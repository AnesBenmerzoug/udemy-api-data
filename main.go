package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"

	dataTypes "github.com/AnesBenmerzoug/udemy-api-data/internal/data_types"
	endpoints "github.com/AnesBenmerzoug/udemy-api-data/internal/endpoints"
	"github.com/joho/godotenv"
)

var CLIENT_ID, CLIENT_SECRET string

const DATA_DIR = "data"

var COURSES_DATA_FILE = path.Join(DATA_DIR, "courses")

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
	context := context.Background()
	coursesChannel := make(chan *dataTypes.Course)
	go func() {
		err := endpoints.GetCourses(context, client, CLIENT_ID, CLIENT_SECRET, coursesChannel)
		if err != nil {
			fmt.Fprintf(os.Stdout, "error: %v\n", err)
			os.Exit(1)
		}
	}()
	var courses []*dataTypes.Course
	for course := range coursesChannel {
		courses = append(courses, course)
	}
	for _, course := range courses {
		fmt.Println(course)
	}
}
