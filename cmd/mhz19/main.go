package main

import (
	"fmt"
	"os"

	"github.com/tmsc4zht/mhz19"
)

func main() {
	m := mhz19.Client{}
	if err := m.Connect(); err != nil {
		fmt.Printf("{\"error\": \"%s\"}", err)
		os.Exit(0)
	}
	v, err := m.ReadCO2()
	if err != nil {
		fmt.Printf("{\"error\": \"%s\"}", err)
		os.Exit(0)
	}
	fmt.Printf("{\"co2\": \"%d\"}", v)
}
