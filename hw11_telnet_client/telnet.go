package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type telnetClient struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &telnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (t *telnetClient) Connect() error {
	var err error
	t.conn, err = net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return fmt.Errorf("could not connect to the server %s", t.address)
	}
	return nil
}

func (t *telnetClient) Close() error {
	return t.conn.Close()
}

func (t *telnetClient) Send() error {
	if _, err := io.Copy(t.conn, t.in); err != nil {
		return fmt.Errorf("could not send data %s to the server %s", t.in, t.address)
	}
	return nil
}

func (t *telnetClient) Receive() error {
	if _, err := io.Copy(t.out, t.conn); err != nil {
		return fmt.Errorf("could not receive data from the server %s", t.address)
	}
	return nil
}
