package main

import (
	"fmt"
	"github.com/combo/brick/env"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"
)

var (
	envCommand       = kingpin.Command("env", "Manage environments")
	envCreateCommand = envCommand.Command("create", "Create an environment")
	envCreatePath    = envCreateCommand.Arg("path", "Directory").Required().String()
)

func abort(message ...interface{}) {
	fmt.Println(message...)
	os.Exit(2)
}

func packageNameFromDirectory(dir string) (string, error) {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}

	return filepath.Base(absDir), nil
}

func doEnvCreate(dir string) {
	name, err := packageNameFromDirectory(dir)
	if err != nil {
		abort(err)
	}

	_, err = env.Create(dir, name)
	if err != nil {
		abort(err)
	}
}

func main() {
	switch kingpin.Parse() {
	case envCreateCommand.FullCommand():
		doEnvCreate(*envCreatePath)
	}
}
