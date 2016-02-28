package env

import (
	"os"
)

type Environment struct {
}

func Create(dir string) (*Environment, error) {
	// Create the directory if it doesn't exist
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		if !os.IsExist(err) {
			return nil, err
		}
	}

	return &Environment{}, nil
}
