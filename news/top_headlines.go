package news

var topHeadlinesEndpoint = "/v2/top-headlines"

func (c *Client) TopHeadlines(params map[string]string) (map[string]interface{}, error) {
	return c.fetchData(topHeadlinesEndpoint, params)
}
