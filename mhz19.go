package mhz19

import (
	"io"

	"github.com/tarm/serial"
)

type Client struct {
	port *serial.Port
}

func (m *Client) Connect() error {
	c := &serial.Config{Name: "/dev/ttyS0", Baud: 9600}
	p, err := serial.OpenPort(c)
	if err != nil {
		return err
	}
	m.port = p
	return nil
}

func (m *Client) ReadCO2() (int, error) {
	err := m.writeToReadCO2()
	if err != nil {
		return 0, err
	}

	// read bytes
	buf := make([]byte, 9)
	_, err = io.ReadFull(m.port, buf)
	if err != nil {
		return 0, err
	}
	return int(buf[2])*256 + int(buf[3]), nil
}

func (m *Client) writeToReadCO2() error {
	_, err := m.port.Write([]byte{0xFF, 0x01, 0x86, 0x00, 0x00, 0x00, 0x00, 0x00, 0x79})
	if err != nil {
		return err
	}
	return nil
}
