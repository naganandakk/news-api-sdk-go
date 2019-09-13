package news

import (
	"github.com/mitchellh/mapstructure"
)

var topHeadlinesEndpoint = "/v2/top-headlines"

func (c *Client) TopHeadlines(params map[string]string) (Result, error) {
	var result Result

	data, err := c.fetchData(topHeadlinesEndpoint, params)
	if err != nil {
		return result, err
	}

	err = mapstructure.Decode(data, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
