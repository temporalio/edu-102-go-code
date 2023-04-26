package example

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func RetrieveEstimate(ctx context.Context, name string) (int, error) {
	base := "https://api.agify.io/?name=%s"
	url := fmt.Sprintf(base, url.QueryEscape(name))

	resp, err := http.Get(url)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	// This struct represents the JSON returned by the API
	type EstimatorResponse struct {
		Age   int    `json:"age"`
		Count int    `json:"count"`
		Name  string `json:"name"`
	}

	// read the HTTP response body into a string for parsing
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	// Create and populate the struct with data from the API call
	var response EstimatorResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return -1, err
	}

	return response.Age, nil
}
