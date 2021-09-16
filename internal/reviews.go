package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func GetReviews(ctx context.Context, client *http.Client, courseId int, clientId, clientSecret string) ([]*Review, error) {
	log.Print("Getting Reviews data from the api")
	var reviews []*Review

	reviewsUrl := fmt.Sprintf("https://www.udemy.com/api-2.0/courses/%v/reviews", courseId)
	queryParameters := url.Values{}
	queryParameters.Add("page_size", "25")
	encodedQueryParams := queryParameters.Encode()

	try := 0
	maxAttempts := 10
	sleepTime := 30 * time.Second

	var resp *http.Response
	var err error

	for {
		try++
		log.Printf("Attempt %v", try)
		resp, err = makeRequest(ctx, client, reviewsUrl, &encodedQueryParams, clientId, clientSecret)
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

	var apiResponse = &ReviewAPIResponse{}

	err = json.NewDecoder(resp.Body).Decode(apiResponse)
	if err != nil {
		log.Printf("Failed parsing api data: %v", err)
		return nil, err
	}
	resp.Body.Close()

	reviews = append(reviews, apiResponse.Reviews...)
	page := 2

	for apiResponse.Next != nil {
		log.Printf("Getting Reviews data for page %v", page)

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
		reviews = append(reviews, apiResponse.Reviews...)
		page++
	}

	for _, review := range reviews {
		review.CourseId = courseId
	}

	return reviews, nil
}
