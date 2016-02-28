package main

import (
	"github.com/combo/brick/env"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

var (
	envCommand       = kingpin.Command("env", "Manage environments")
	envCreateCommand = envCommand.Command("create", "Create an environment")
	envCreatePath    = envCreateCommand.Arg("path", "Directory").Required().String()
)

func main() {
	switch kingpin.Parse() {
	case envCreateCommand.FullCommand():
		_, err := env.Create(*envCreatePath)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
