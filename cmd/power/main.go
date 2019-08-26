package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fudge/power/cli"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("must provide a folder or gpx file")
		os.Exit(2)
	}

	a := cli.NewCli()
	flag.Var(&a.OutputType, "output", "0: text, 1: json, 2: svg")
	flag.Parse()

	exit, _ := a.Run(flag.Args()...)
	os.Exit(exit)
}
