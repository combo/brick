package env

import (
	"os"

	"github.com/combo/brick/pkg"
)

type Environment struct {
}

func Create(dir string, name string) (*Environment, error) {
	// Create the directory if it doesn't exist
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		if !os.IsExist(err) {
			return nil, err
		}
	}

	// Create a metadata file
	m := pkg.Metadata{
		Name: name,
	}
	err = m.Save(dir)
	if err != nil {
		return nil, err
	}

	return &Environment{}, nil
}
