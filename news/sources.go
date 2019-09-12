package news

func (c *Client) Sources(params map[string]string) (map[string]interface{}, error) {
	return c.fetchData("/v2/sources", params)
}
