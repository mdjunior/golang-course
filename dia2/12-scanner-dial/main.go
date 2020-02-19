// https://github.com/blackhat-go/bhg/blob/master/ch-2/dial/main.go
package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("Connection successful at port 80")
	}

	for port := 20; port < 30; port++ {
		host := fmt.Sprintf("%s:%d", "scanme.nmap.org", port)
		fmt.Println(host)
		_, err := net.Dial("tcp", host)
		if err == nil {
			fmt.Printf("%d -> open\n", port)
		} else {
			fmt.Printf("%d -> %s\n", port, err.Error())
		}
	}
}
