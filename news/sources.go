package news

var sourcesEndpoint = "/v2/sources"

func (c *Client) Sources(params map[string]string) (map[string]interface{}, error) {
	return c.fetchData(sourcesEndpoint, params)
}
