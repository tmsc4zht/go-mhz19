package main

import (
	"fmt"
	"os"

	"github.com/tmsc4zht/mhz19"
)

func main() {
	m := mhz19.New("/dev/serial0")
	v, err := m.ReadCO2()
	if err != nil {
		fmt.Printf("{\"error\": \"%s\"}\n", err)
		os.Exit(-1)
	}
	fmt.Printf("{\"co2\": \"%d\"}\n", v)
}
