# mhz19
Read CO2 concentration from MH-Z19 on Raspberry Pi.

## Installation
```shell script
$ go get github.com/tmsc4zht/mhz19/cmd/mhz19
```

## Usage
```shell script
$ mhz19
554
```

## Use as a library
```go
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
```