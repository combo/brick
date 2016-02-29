package main

import (
	"fmt"
	"github.com/combo/brick/pkg"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"
)

var (
	newCommand = kingpin.Command("new", "Create a new package")
	newPath    = newCommand.Arg("path", "Directory").Required().String()
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

func doNew(dir string) {
	name, err := packageNameFromDirectory(dir)
	if err != nil {
		abort(err)
	}

	err = pkg.New(dir, name)
	if err != nil {
		abort(err)
	}
}

func main() {
	switch kingpin.Parse() {
	case newCommand.FullCommand():
		doNew(*newPath)
	}
}
