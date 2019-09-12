package news

import (
	"testing"
)

func TestSources(t *testing.T) {
	newsClient := NewClient("f251266458e947ba94e465c731e10a2a")

	_, err := newsClient.Sources(map[string]string{})

	if err != nil {
		t.Error(err)
	}
}
