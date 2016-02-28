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

func TestMetadataValidateName(t *testing.T) {
	var tests = []struct {
		name     string
		expected bool
	}{
		{"foo", true},
		{"foo-bar", true},
		{"foo_bar", true},
		{"foo.bar", true},
		{"foo_", false},
		{".foo", false},
		{"foo--bar", false},
		{"foo__bar", false},
	}

	for _, test := range tests {
		if valid := ValidateName(test.name); valid != test.expected {
			t.Errorf("ValidateName(%q) = %v, expected %v\n", test.name, valid, test.expected)
		}
	}
}
