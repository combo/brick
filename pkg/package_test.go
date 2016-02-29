package pkg

import (
	"archive/tar"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// TODO: clean this mess
func TestCreate(t *testing.T) {
	name := "test"
	var files = []struct {
		path     string
		contents string
		seen     bool
	}{
		{"some/file", "foo", false},
		{"some/other/file", "bar", false},
	}

	// Create an empty package somewhere
	tempDir, err := ioutil.TempDir("", "brick_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	err = New(tempDir, name)
	if err != nil {
		t.Fatalf("Create: %s\n", err)
	}

	// Add some files to the tree
	for _, file := range files {
		path := filepath.Join(tempDir, TREE_DIR, file.path)

		// Create parent directories
		if err = os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatal(err)
		}

		// Write the contents
		if err = ioutil.WriteFile(path, []byte(file.contents), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Create the archive
	var archive bytes.Buffer

	if _, err = Package(&archive, tempDir); err != nil {
		t.Fatal(err)
	}

	// Check the archive contents
	tr := tar.NewReader(&archive)
	metadataSeen := false

outer:
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}

		if hdr.Name == METADATA_FILE {
			metadataSeen = true
			continue
		}

		for i, file := range files {
			if file.seen {
				continue
			}

			if hdr.Name != filepath.Join(TREE_DIR, file.path) {
				continue
			}

			files[i].seen = true
			// TODO: check file contents
			continue outer
		}
	}

	// TODO: check metadata validity
	if !metadataSeen {
		t.Fatalf("Metadata file not found in archive.\n")
	}

	for _, file := range files {
		if !file.seen {
			t.Fatalf("File %s not found in archive.\n", file.path)
		}
	}
}
