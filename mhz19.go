package mhz19

import (
	"fmt"
	"io"

	"github.com/tarm/serial"
)

type Client struct {
	s *serial.Config
}

func New(name string) *Client {
	s := &serial.Config{
		Name:     name,
		Baud:     9600,
		Size:     8,
		StopBits: serial.Stop1,
		Parity:   serial.ParityNone,
	}
	return &Client{s: s}
}

func (m *Client) ReadCO2() (int, error) {
	p, err := serial.OpenPort(m.s)
	if err != nil {
		return 0, fmt.Errorf("could not open port: %v", err)
	}
	defer p.Close()

	readCO2concentrationCommand := []byte{0xFF, 0x01, 0x86, 0x00, 0x00, 0x00, 0x00, 0x00, 0x79}
	if _, err := p.Write(readCO2concentrationCommand); err != nil {
		return 0, fmt.Errorf("could not send command: %v", err)
	}

	buf := make([]byte, 9)
	n, err := io.ReadFull(p, buf)
	if err != nil {
		return 0, fmt.Errorf("could not read result: %v", err)
	}
	if n != 9 {
		return 0, fmt.Errorf("return value must 9 bytes but got %d byte(s)", n)
	}

	return int(buf[2])*256 + int(buf[3]), nil
}
