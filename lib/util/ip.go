package util

import (
	"fmt"
	"net"
)

var localIp string

func init() {
	localIp, _ = getLocalIP()
}

func GetLocalIp() string {
	return localIp
}

func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok || ipAddr.IP.IsLoopback() || !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return "", fmt.Errorf("get local ip failed")
}
