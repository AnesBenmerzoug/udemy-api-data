package endpoints

import (
	"encoding/json"
	"net/http"

	dataTypes "github.com/AnesBenmerzoug/udemy-api-data/internal/data_types"
)

func GetCourses(client *http.Client, client_id, client_secret string) (*dataTypes.CourseAPIResponse, error) {
	req, err := http.NewRequest("GET", "https://www.udemy.com/api-2.0/courses/", nil)
	if err != nil {
		return nil, err
	}
	// Set query parameters
	queryParameters := req.URL.Query()
	queryParameters.Add("page_size", "100")
	req.URL.RawQuery = queryParameters.Encode()
	req.SetBasicAuth(client_id, client_secret)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var target = &dataTypes.CourseAPIResponse{}
	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return nil, err
	}
	return target, nil
}
