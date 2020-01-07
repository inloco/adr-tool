package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/inloco/adr-tool/config"
)

var (
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	name    = kingpin.Arg("name", "Name of user.").Required().String()
)

func main() {
	p, _ := config.LoadConfigFile(".")

	kingpin.Parse()
	fmt.Printf("%v, %s %s\n", *verbose, *name, p)
}
