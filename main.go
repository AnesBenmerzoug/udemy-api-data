package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path"

	internal "github.com/AnesBenmerzoug/udemy-api-data/internal"
	"github.com/gocarina/gocsv"
	"github.com/joho/godotenv"
)

const DATA_DIR = "data"

var (
	CLIENT_ID, CLIENT_SECRET string
	COURSES_DATA_FILE        = path.Join(DATA_DIR, "courses.csv")
	REVIEWS_DATA_FILE        = path.Join(DATA_DIR, "reviews.csv")
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	CLIENT_ID = os.Getenv("CLIENT_ID")
	CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
	if _, err := os.Stat(DATA_DIR); os.IsNotExist(err) {
		os.Mkdir(DATA_DIR, 0777)
	}
}

func main() {
	client := &http.Client{}
	context := context.Background()

	coursesFile, err := os.OpenFile(COURSES_DATA_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	defer coursesFile.Close()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	reviewsFile, err := os.OpenFile(REVIEWS_DATA_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	defer reviewsFile.Close()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	courses, err := internal.GetCourses(context, client, CLIENT_ID, CLIENT_SECRET)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	log.Print("Writing Courses data to file")
	err = gocsv.MarshalFile(courses, coursesFile)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	var allReviews []*internal.Review

	for _, course := range courses {
		reviews, err := internal.GetReviews(context, client, course.Id, CLIENT_ID, CLIENT_SECRET)
		if err != nil {
			log.Fatalf("error: %v\n", err)
		}
		allReviews = append(allReviews, reviews...)
	}
	log.Print("Writing Reviews data to file")
	err = gocsv.MarshalFile(allReviews, reviewsFile)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
}
