package constants

import (
	"net"
	"os"
	"strings"
)

func GetIp(key string) string {
	ip := os.Getenv(key)
	if ip == "" {
		ip = "localhost"
	}
	return ip
}

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return "localhost", nil
}
