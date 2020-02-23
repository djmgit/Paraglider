package glidercore

import (
	"net"
	"strings"
	"strconv"
	"time"
	"paraglider/glider/models"
)

func CheckTcp(tcpHalthCheck models.TcpHealthCheckHolder) bool {

	targetAddr := strings.Join([]string{tcpHalthCheck.TargetIP, strconv.Itoa(tcpHalthCheck.TargetPort)}, ":")
	_, err := net.DialTimeout("tcp", targetAddr, time.Duration(tcpHalthCheck.TimeOut)*time.Second)

	if err != nil {
		return false
	}

	return true
}
