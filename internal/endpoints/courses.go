package endpoints

import (
	"context"
	"encoding/json"
	"net/http"

	dataTypes "github.com/AnesBenmerzoug/udemy-api-data/internal/data_types"
)

func GetCourses(ctx context.Context, client *http.Client, client_id, client_secret string, ch chan *dataTypes.Course) error {
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
	defer resp.Body.Close()
	var target = &dataTypes.CourseAPIResponse{}
	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return err
	}
	for _, course := range target.Courses {
		ch <- course
	}
	close(ch)
	return nil
}
