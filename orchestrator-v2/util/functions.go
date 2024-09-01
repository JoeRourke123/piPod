package util

import (
	"fmt"
	"net"
)

func Map[A any, B any](vs []A, f func(A) B) []B {
	vsm := make([]B, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("error fetching local IP: ", err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}
