package internal

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func makeRequest(ctx context.Context, client *http.Client, url string, queryParameters *string, clientId, clientSecret string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if queryParameters != nil {
		req.URL.RawQuery = *queryParameters
	}
	req.SetBasicAuth(clientId, clientSecret)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed getting data from api: %v", err)
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
