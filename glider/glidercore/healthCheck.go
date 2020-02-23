package glidercore

import (
	"net"
	"net/http"
	"io/ioutil"
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

func CheckHTTP(HTTPHealthCheck models.HTTPHealthCheckHolder) bool {

	targetAddr := strings.Join([]string{HTTPHealthCheck.TargetIP, strconv.Itoa(HTTPHealthCheck.TargetPort)}, ":")
	HTTPHealthCheckPath := strings.Join([]string{targetAddr, HTTPHealthCheck.HealthCheckPath}, "/")

	response, err := http.Get(HTTPHealthCheckPath)
	if err != nil {
		return false
	}

	body, errBody := ioutil.ReadAll(response.Body)
	if errBody != nil {
		return false
	}

	bodyStr := string(body)

	if bodyStr != HTTPHealthCheck.PassMessage {
		return false
	}

	return true
}
