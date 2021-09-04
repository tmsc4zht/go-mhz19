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

	if err := checkSum(buf); err != nil {
		return 0, fmt.Errorf("check sum failed: %v", err)
	}

	return int(buf[2])*256 + int(buf[3]), nil
}

func checkSum(b []byte) error {
	if len(b) != 9 {
		return fmt.Errorf("return value length must be 9 but %d", len(b))
	}
	sum := byte(0)
	for i := 1; i < 8; i++ {
		sum += b[i]
	}
	if b[8] != -sum {
		return fmt.Errorf("check sum failed expected %v, calclated: %v", b[8], -sum)
	}

	return nil
}
