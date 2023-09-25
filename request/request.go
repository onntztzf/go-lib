package request

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
)

// GET sends a GET request to the specified URL with the given query parameters and returns the response body.
func GET(URL string, params map[string]string) ([]byte, error) {
	// Create a new GET request with the provided URL.
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %v", err)
	}

	// Encode and set the query parameters in the request URL.
	q := make(url.Values)
	for k, v := range params {
		q.Set(k, v)
	}
	req.URL.RawQuery = q.Encode()

	// Send the GET request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET request failed: %v", err)
	}
	defer resp.Body.Close()

	// Read and return the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil
}

// POST sends a POST request to the specified URL with the provided data and returns the response body.
func POST(URL string, data map[string]interface{}) ([]byte, error) {
	// Serialize the data to JSON format.
	bodyData, err := jsoniter.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize POST data: %v", err)
	}

	// Create a new POST request with the serialized data.
	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewReader(bodyData))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request: %v", err)
	}

	// Set the Content-Type header to indicate JSON data.
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("POST request failed: %v", err)
	}
	defer resp.Body.Close()

	// Read and return the response body.
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return res, nil
}
