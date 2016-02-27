package brick

import (
	"os"
)

type Environment struct {
}

func CreateEnvironment(dir string) (*Environment, error) {
	// Create the directory if it doesn't exist
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		if !os.IsExist(err) {
			return nil, err
		}
	}

	return &Environment{}, nil
}
