package main

import (
	"os"

	"github.com/fudge/power"
)

func main() {
	a := power.NewCli()
	exit, _ := a.Run(os.Args...)
	os.Exit(exit)
}
