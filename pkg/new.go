package pkg

import (
	"os"
)

func New(dir string, name string) error {
	// Create the directory if it doesn't exist
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	// Create a metadata file
	m := Metadata{
		Name: name,
	}
	err = m.Save(dir)
	if err != nil {
		return err
	}

	return nil
}
