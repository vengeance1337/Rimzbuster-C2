package enum

import (
	"bufio"
	"net"
	"strings"
)

func GetUsername(conn net.Conn) string {
	conn.Write([]byte("whoami\n"))
	reader := bufio.NewReader(conn)
	username, _ := reader.ReadString('\n')
	return strings.TrimSpace(username)
}

func GetOSInfo(conn net.Conn) string {
	conn.Write([]byte("osinfo\n"))
	reader := bufio.NewReader(conn)
	osInfo, _ := reader.ReadString('\n')
	return strings.TrimSpace(osInfo)
}

