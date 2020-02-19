// https://github.com/blackhat-go/bhg/blob/master/ch-2/tcp-scanner-too-fast/main.go
package main

import (
	"fmt"
	"net"
	"sync"
)

func scan(port int, waitGroup *sync.WaitGroup) error {
	address := fmt.Sprintf("scanme.nmap.org:%d", port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("%d: %s\n", port, err.Error())
		waitGroup.Done()
		return err
	}
	conn.Close()
	fmt.Printf("%d: open\n", port)
	waitGroup.Done()
	return nil
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(10)

	for port := 20; port < 30; port++ {
		go scan(port, &waitGroup)
	}
	waitGroup.Wait()
}
