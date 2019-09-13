package news

import (
	"github.com/mitchellh/mapstructure"
)

var sourcesEndpoint = "/v2/sources"

func (c *Client) Sources(params map[string]string) ([]Source, error) {
	var sources []Source

	data, err := c.fetchData(sourcesEndpoint, params)
	if err != nil {
		return sources, err
	}

	err = mapstructure.Decode(data["sources"], &sources)
	if err != nil {
		return sources, err
	}

	return sources, nil
}
