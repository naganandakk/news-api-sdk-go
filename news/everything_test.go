package news

import (
	"testing"
)

func TestEverything(t *testing.T) {
	newsClient := NewClient("f251266458e947ba94e465c731e10a2a")

	_, err := newsClient.Everything(map[string]string{
		"domains": "google.com",
	})

	if err != nil {
		t.Error(err)
	}
}
