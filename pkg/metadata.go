package pkg

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"regexp"
)

const METADATA_FILE = "brick.package.json"

var nameRegexp = regexp.MustCompile(`^[0-9a-z]+([-_\.][0-9a-z]+)*$`)

type Metadata struct {
	Name string `json:"name"`
}

func ValidateName(name string) bool {
	return nameRegexp.MatchString(name)
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
		return nil, err
	}

	return &m, nil
}

// Verify the validity of the metadata fields
func (m *Metadata) Validate() error {
	if !ValidateName(m.Name) {
		return errors.New("Invalid name")
	}

	return nil
}

// Save the package metadata into the specified package directory
func (m *Metadata) Save(dir string) error {
	err := m.Validate()
	if err != nil {
		return err
	}

	path := filepath.Join(dir, METADATA_FILE)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	bytes, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		return err
	}
	_, err = f.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
