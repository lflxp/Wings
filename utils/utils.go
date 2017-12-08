package utils

import (
	"net"
	"time"
)

func ScannerPort(ipAndPort string) bool {
	rs := false
	//tcpaddr,_ := net.ResolveTCPAddr("tcp4",ipAndPort)
	//_,err := net.DialTCP("tcp",nil,tcpaddr)
	_, err := net.DialTimeout("tcp", ipAndPort, 500*time.Millisecond)
	if err == nil {
		rs = true
	}
	return rs
}