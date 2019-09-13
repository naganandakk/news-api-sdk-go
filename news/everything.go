package news

import (
	"github.com/mitchellh/mapstructure"
)

var everythingEndpoint = "/v2/everything"

func (c *Client) Everything(params map[string]string) (Result, error) {
	var result Result

	data, err := c.fetchData(everythingEndpoint, params)
	if err != nil {
		return result, err
	}

	err = mapstructure.Decode(data, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
