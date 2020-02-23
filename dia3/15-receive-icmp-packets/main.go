// https://www.darkcoding.net/uncategorized/raw-sockets-in-go-ip-layer/
package main

import (
	"fmt"
	"net"
)

func main() {
	protocol := "icmp"
	netaddr, _ := net.ResolveIPAddr("ip4", "127.0.0.1")
	conn, _ := net.ListenIP("ip4:"+protocol, netaddr)

	buf := make([]byte, 1024)
	numRead, addr, err := conn.ReadFrom(buf)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	conn.ReadMsgIP()
	fmt.Printf("% X\n", buf[:numRead])
}
