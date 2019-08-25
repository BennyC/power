package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fudge/power"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalf("no file provided as a parameter")
		return
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open gpx: %v", err)
		return
	}

	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)

	var gpx power.Gpx
	xml.Unmarshal(byteValue, &gpx)

	fmt.Printf("%v\n", gpx.Name)
	for _, i := range []int{1, 5, 10, 15, 20} {
		p := gpx.CalculateBestPower(60 * i)
		fmt.Printf("best %v min power: %v\n", i, p)
	}
}
