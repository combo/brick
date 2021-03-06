package pkg

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/BurntSushi/toml"
)

const METADATA_FILE = "brick.package.toml"

var nameRegexp = regexp.MustCompile(`^[0-9a-z]+([-_\.][0-9a-z]+)*$`)

type Metadata struct {
	Name string `toml:"name"`
}

func ValidateName(name string) bool {
	return nameRegexp.MatchString(name)
}

// Read and return the metadata for the specified package directory
func OpenMetadata(dir string) (*Metadata, error) {
	path := filepath.Join(dir, METADATA_FILE)

	var m Metadata
	_, err := toml.DecodeFile(path, &m)
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

// Write the package metadata into the specified io.Writer
func (m *Metadata) Write(w io.Writer) error {
	err := m.Validate()
	if err != nil {
		return err
	}

	encoder := toml.NewEncoder(w)
	err = encoder.Encode(m)

	if err != nil {
		return err
	}

	return nil
}

// Save the package metadata into the specified package directory
func (m *Metadata) Save(dir string) error {
	path := filepath.Join(dir, METADATA_FILE)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return m.Write(f)
}
