package brick

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateEnvironment(t *testing.T) {
	// Get the name of a temporary directory we can use
	tempDir, err := ioutil.TempDir("", "brick_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create the virtual environment
	_, err = CreateEnvironment(tempDir)
	if err != nil {
		t.Fatalf("CreateEnvironment: %s\n", err)
	}
}
