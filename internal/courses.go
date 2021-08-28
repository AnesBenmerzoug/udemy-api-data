package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func makeRequest(ctx context.Context, client *http.Client, url string, queryParameters *string, client_id, client_secret string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if queryParameters != nil {
		req.URL.RawQuery = *queryParameters
	}
	req.SetBasicAuth(client_id, client_secret)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed getting courses data from api: %v", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(string(bodyBytes))
		return nil, fmt.Errorf("error while making request to udemy api: %v", url)
	}
	return resp, nil
}

func GetCourses(ctx context.Context, client *http.Client, client_id, client_secret string) ([]*Course, error) {
	log.Print("Getting Courses data from the api")
	var courses []*Course

	coursesUrl := "https://www.udemy.com/api-2.0/courses/"
	queryParameters := url.Values{}
	queryParameters.Add("page_size", "50")
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
		resp, err = makeRequest(ctx, client, coursesUrl, &encodedQueryParams, client_id, client_secret)
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
			resp, err = makeRequest(ctx, client, *apiResponse.Next, nil, client_id, client_secret)
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
