package internal

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"
)

func GetCourses(ctx context.Context, client *http.Client, clientId, clientSecret string) ([]*Course, error) {
	log.Print("Getting Courses data from the api")
	var courses []*Course

	coursesUrl := "https://www.udemy.com/api-2.0/courses/"
	queryParameters := url.Values{}
	queryParameters.Add("page_size", "1")
	queryParameters.Add("fields[course]", "@all")
	encodedQueryParams := queryParameters.Encode()

	try := 0
	maxAttempts := 10
	sleepTime := 30 * time.Second

	var resp *http.Response
	var err error

	for {
		try++
		log.Printf("Attempt %v", try)
		resp, err = makeRequest(ctx, client, coursesUrl, &encodedQueryParams, clientId, clientSecret)
		if err != nil {
			if try < maxAttempts {
				log.Print(err)
				time.Sleep(sleepTime)
				continue
			} else {
				return nil, err
			}
		}
		break
	}

	var apiResponse = &CourseAPIResponse{}

	err = json.NewDecoder(resp.Body).Decode(apiResponse)
	if err != nil {
		log.Printf("Failed parsing api data: %v", err)
		return nil, err
	}
	resp.Body.Close()

	courses = append(courses, apiResponse.Courses...)
	page := 2

	for apiResponse.Next != nil {
		log.Printf("Getting Courses data for page %v", page)

		try := 0

		for {
			try++
			log.Printf("Attempt %v", try)
			resp, err = makeRequest(ctx, client, *apiResponse.Next, nil, clientId, clientSecret)
			if err != nil {
				if try < maxAttempts {
					log.Print(err)
					time.Sleep(sleepTime)
					continue
				} else {
					return nil, err
				}
			}
			break
		}

		err = json.NewDecoder(resp.Body).Decode(apiResponse)
		if err != nil {
			log.Printf("Failed parsing api data: %v", err)
			log.Print(resp.Body)
			return nil, err
		}
		resp.Body.Close()
		courses = append(courses, apiResponse.Courses...)
		page++
	}
	return courses, nil
}
