package endpoints

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	dataTypes "github.com/AnesBenmerzoug/udemy-api-data/internal/data_types"
)

func GetCourses(ctx context.Context, client *http.Client, client_id, client_secret string, ch chan *dataTypes.Course) error {
	log.Print("Getting Courses data from the api")
	page := 1
	defer close(ch)
	req, err := http.NewRequest("GET", "https://www.udemy.com/api-2.0/courses/", nil)
	if err != nil {
		return err
	}
	// Set query parameters
	queryParameters := req.URL.Query()
	queryParameters.Add("page_size", "100")
	req.URL.RawQuery = queryParameters.Encode()
	req.SetBasicAuth(client_id, client_secret)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	var apiResponse = &dataTypes.CourseAPIResponse{}
	err = json.NewDecoder(resp.Body).Decode(apiResponse)
	if err != nil {
		return err
	}
	for _, course := range apiResponse.Courses {
		ch <- course
	}
	resp.Body.Close()
	page++
	for apiResponse.Next != nil {
		log.Printf("Getting Courses data for page %v", page)
		req, err = http.NewRequest("GET", *apiResponse.Next, nil)
		if err != nil {
			return err
		}
		req.SetBasicAuth(client_id, client_secret)
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		err = json.NewDecoder(resp.Body).Decode(apiResponse)
		if err != nil {
			return err
		}
		for _, course := range apiResponse.Courses {
			ch <- course
		}
		resp.Body.Close()
		if apiResponse.Next == nil {
			break
		}
		page++
	}
	return nil
}
