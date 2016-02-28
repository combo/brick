package pkg

import (
	"encoding/json"
	"testing"
)

func TestMetadataCanBeEncodedToJson(t *testing.T) {
	m := Metadata{
		Name: "foo",
	}

	_, err := json.Marshal(&m)
	if err != nil {
		t.Fatal(err)
	}
}
