package news

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const host = "https://newsapi.org"

var httpClient = &http.Client{}

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey,
		httpClient,
	}
}

func (c *Client) fetchData(endpoint string, queryParms map[string]string) (map[string]interface{}, error) {
	response := make(map[string]interface{})

	apiURL, err := buildURL(host, endpoint, queryParms)
	if err != nil {
		return response, err
	}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return response, err
	}

	req.Header.Set("X-Api-Key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return response, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return response, errors.New("api response status code is other than 2xx")
	}

	return response, nil
}
