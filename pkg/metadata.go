package pkg

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const METADATA_FILE = "brick.package.json"

type Metadata struct {
	Name string
}

// Read and return the metadata for the specified package directory
func OpenMetadata(dir string) (*Metadata, error) {
	path := filepath.Join(dir, METADATA_FILE)

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var m Metadata

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&m)
	if err != nil {
		return nil, nil
	}

	return &m, nil
}

// Save the package metadata into the specified package directory
func (m *Metadata) Save(dir string) error {
	path := filepath.Join(dir, METADATA_FILE)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	return encoder.Encode(m)
}
