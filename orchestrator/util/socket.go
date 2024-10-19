package util

import (
	"context"
	"orchestrator/util/logger"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

const (
	PacketSize = 128
	SocketIP   = "127.0.0.1"
	SocketPort = 9090
)

func OpenSocketConnection() int {
	serverFD, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)

	if err != nil {
		logger.Error(
			context.Background(),
			"error occured while opening socket",
			err, logger.SocketOperationTag, logger.FromTag("OpenSocketConnection"),
		)
	}

	unix.SetNonblock(serverFD, true)

	serverAddr := &unix.SockaddrInet4{
		Port: SocketPort,
		Addr: inetAddr(SocketIP),
	}

	err = unix.Bind(serverFD, serverAddr)
	if err != nil {
		if err == unix.ECONNREFUSED {
			logger.Error(
				context.Background(),
				"error occured while connecting to socket",
				err, logger.SocketOperationTag, logger.FromTag("OpenSocketConnection"),
			)
			unix.Close(serverFD)
			return -1
		}
	}

	return serverFD
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
