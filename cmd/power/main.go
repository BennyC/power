package main

import (
	"fmt"
	"os"

	"github.com/fudge/power"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("must provide a folder or gpx file")
		os.Exit(2)
	}

	a := power.NewCli()
	exit, _ := a.Run(os.Args[1:]...)
	os.Exit(exit)
}
