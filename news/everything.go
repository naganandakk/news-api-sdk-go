package news

var everythingEndpoint = "/v2/everything"

func (c *Client) Everything(params map[string]string) (map[string]interface{}, error) {
	return c.fetchData(everythingEndpoint, params)
}
