package news

import (
	"testing"
)

func TestTopHeadlines(t *testing.T) {
	newsClient := NewClient("f251266458e947ba94e465c731e10a2a")

	_, err := newsClient.TopHeadlines(map[string]string{
		"country": "in",
	})

	if err != nil {
		t.Error(err)
	}
}
