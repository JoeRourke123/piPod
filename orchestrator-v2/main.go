package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

var (
	MAXMSGSIZE = 128
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("./client [IPv4] [Port]")
	}

	serverFD, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)
	if err != nil {
		log.Fatal("Socket: ", err)
	}

	unix.SetNonblock(serverFD, true)

	port, err := strconv.Atoi(args[1])
	if err != nil || (port < 0 || port > 100000) {
		os.Stderr.WriteString("Invalid port format\n")
		return
	}
	serverAddr := &unix.SockaddrInet4{
		Port: port,
		Addr: inetAddr(args[0]),
	}

	err = unix.Bind(serverFD, serverAddr)
	if err != nil {
		if err == unix.ECONNREFUSED {
			fmt.Println("* Connection failed")
			unix.Close(serverFD)
			return
		}
	}

	var response []byte

	response = make([]byte, MAXMSGSIZE)

	defer unix.Close(serverFD)

	for {
		_, _, err = unix.Recvfrom(serverFD, response, 0)
		if err == nil {
			fmt.Println(response[0])
			fmt.Println(response[1])
			fmt.Println(response[2])
			fmt.Println("________")
		}
	}

	return
}

func inetAddr(ipaddr string) [4]byte {
	var (
		ip                 = strings.Split(ipaddr, ".")
		ip1, ip2, ip3, ip4 uint64
	)
	ip1, _ = strconv.ParseUint(ip[0], 10, 8)
	ip2, _ = strconv.ParseUint(ip[1], 10, 8)
	ip3, _ = strconv.ParseUint(ip[2], 10, 8)
	ip4, _ = strconv.ParseUint(ip[3], 10, 8)
	return [4]byte{byte(ip1), byte(ip2), byte(ip3), byte(ip4)}
}
