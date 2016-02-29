package pkg

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	name := "foo"

	// Get the name of a temporary directory we can use
	tempDir, err := ioutil.TempDir("", "brick_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create the package
	err = New(tempDir, name)
	if err != nil {
		t.Fatalf("Create: %s\n", err)
	}

	// A valid metadata file should have been created
	m, err := OpenMetadata(tempDir)
	if err != nil {
		t.Fatalf("OpenMetadata: %s\n", err)
	}

	if m.Name != name {
		t.Fatalf("Expected package name = %s, got %s\n", name, m.Name)
	}
}
