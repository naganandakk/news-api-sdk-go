package news

import (
	"net/url"
)

func buildURL(host string, endpoint string, queryParms map[string]string) (string, error) {
	urlString, err := url.Parse(host + endpoint)
	if err != nil {
		return "", err
	}

	queryString := urlString.Query()

	for key, value := range queryParms {
		queryString.Add(key, value)
	}

	urlString.RawQuery = queryString.Encode()

	return urlString.String(), nil
}
