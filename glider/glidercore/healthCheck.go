package glidercore

import (
	"net"
	"strings"
	"strconv"
	"paraglider/glider/models"
)

func CheckTcp(tcpHalthCheck models.TcpHealthCheckHolder) bool {

	targetAddr := strings.Join([]string{tcpHalthCheck.TargetIP, strconv.Itoa(tcpHalthCheck.TargetPort)}, ":")
	_, err := net.Dial("tcp", targetAddr)

	if err != nil {
		return false
	}

	return true
}
