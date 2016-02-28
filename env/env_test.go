package env

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/combo/brick/pkg"
)

func TestCreate(t *testing.T) {
	name := "foo"

	// Get the name of a temporary directory we can use
	tempDir, err := ioutil.TempDir("", "brick_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create the virtual environment
	_, err = Create(tempDir, name)
	if err != nil {
		t.Fatalf("Create: %s\n", err)
	}

	// A valid metadata file should have been created
	m, err := pkg.OpenMetadata(tempDir)
	if err != nil {
		t.Fatalf("OpenMetadata: %s\n", err)
	}

	if m.Name != name {
		t.Fatalf("Expected package name = %s, got %s\n", name, m.Name)
	}
}
