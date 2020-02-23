// https://github.com/blackhat-go/bhg/blob/master/ch-2/tcp-scanner-too-fast/main.go
package main

import (
	"fmt"
	"net"
)

func scan(port int) error {
	address := fmt.Sprintf("scanme.nmap.org:%d", port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("%d: %s", port, err.Error())
		return err
	}
	conn.Close()
	fmt.Printf("%d: open", port)
	return nil
}

func main() {
	for port := 20; port <= 30; port++ {
		go scan(port)
	}
}
