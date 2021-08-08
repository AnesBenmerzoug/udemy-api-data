package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	dataTypes "github.com/AnesBenmerzoug/udemy-api-data/internal/data_types"
	endpoints "github.com/AnesBenmerzoug/udemy-api-data/internal/endpoints"
	"github.com/gocarina/gocsv"
	"github.com/joho/godotenv"
)

var CLIENT_ID, CLIENT_SECRET string

const DATA_DIR = "data"

var COURSES_DATA_FILE = path.Join(DATA_DIR, "courses.csv")

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
	coursesChannel := make(chan *dataTypes.Course, 100)
	coursesFile, err := os.OpenFile(COURSES_DATA_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	defer coursesFile.Close()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	go func() {
		err := endpoints.GetCourses(context, client, CLIENT_ID, CLIENT_SECRET, coursesChannel)
		if err != nil {
			log.Fatalf("error: %v\n", err)
		}
	}()
	var courses []*dataTypes.Course
	for {
		select {
		case course, ok := <-coursesChannel:
			if !ok {
				coursesChannel = nil

			}
			courses = append(courses, course)
		case <-time.After(3 * time.Second):
		}

		if len(courses) >= 100 {
			log.Print("Appending Courses data to file")
			err = gocsv.MarshalFile(&courses, coursesFile)
			if err != nil {
				log.Fatalf("error: %v\n", err)
			}
			courses = nil
		}

		if coursesChannel == nil {
			if len(courses) > 0 {
				err = gocsv.MarshalFile(&courses, coursesFile)
				if err != nil {
					log.Fatalf("error: %v\n", err)
				}
			}
			courses = nil
			break
		}
	}
}
