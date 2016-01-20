package goonix

import (
	"fmt"
	"net"
	"time"
)

//Port on network host is reachable/exists (telnet foo.com 25000)

type Network struct {
}

func (n *Network) CheckPort(host string, port int, timeout time.Duration) (bool, error) {
	buf := make([]byte, 512)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%v", host, port))
	if err != nil {
		return false, err
	}
	conn.SetReadDeadline(time.Now().Add(timeout))
	// Conn.Read will raise a timeout error after <timeout>
	_, err = conn.Read(buf)
	return err == nil, err
}
