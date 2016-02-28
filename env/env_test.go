package env

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	// Get the name of a temporary directory we can use
	tempDir, err := ioutil.TempDir("", "brick_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create the virtual environment
	_, err = Create(tempDir)
	if err != nil {
		t.Fatalf("Create: %s\n", err)
	}
}
