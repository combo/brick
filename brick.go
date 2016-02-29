package main

import (
	"fmt"
	"github.com/combo/brick/pkg"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	newCommand = kingpin.Command("new", "Create a new package")
	newPath    = newCommand.Arg("path", "Directory").Required().String()

	packageCommand = kingpin.Command("package", "Build a package")
	packagePath    = packageCommand.Arg("path", "Package directory").String()
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

func doPackage(dir string) {
	// Create a temporary file for the package
	packageTmp, err := ioutil.TempFile(dir, "package")
	if err != nil {
		abort(err)
	}
	defer os.Remove(packageTmp.Name())

	m, err := pkg.Package(packageTmp, dir)
	if err != nil {
		abort(err)
	}
	packageTmp.Close()

	// Rename package
	packageFile := filepath.Join(dir, fmt.Sprintf("%s.brick", m.Name))
	os.Rename(packageTmp.Name(), packageFile)

	fmt.Println("Package written at", packageFile)
}

func main() {
	switch kingpin.Parse() {
	case newCommand.FullCommand():
		doNew(*newPath)

	case packageCommand.FullCommand():
		var err error
		dir := *packagePath
		if dir == "" {
			dir, err = os.Getwd()
			if err != nil {
				abort(err)
			}
		}
		doPackage(dir)
	}
}
