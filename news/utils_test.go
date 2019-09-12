package news

import (
	"testing"
)

func TestBuildURL(t *testing.T) {
	host := "http://127.0.0.1"
	endpoint := "/endpoint"
	queryParams := map[string]string{
		"q1": "v1",
		"q2": "v2",
	}

	want := host + endpoint + "?q1=v1&q2=v2"
	if got, err := buildURL(host, endpoint, queryParams); (got != want) || (err != nil) {
		t.Error("buildURL() = ", got, ", want ", want, ", error ", err)
	}
}
