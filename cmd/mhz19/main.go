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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	fmt.Println(v)
}
