// https://gist.github.com/lmas/c13d1c9de3b2224f9c26435eb56e6ef3
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const (
	ProtocolICMP = 1
	TimeOUT      = 2 * time.Second
)

// Ping is function to Ping hosts by IP
func Ping(addr string) (*net.IPAddr, time.Duration, error) {
	// Start listening for icmp replies
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return nil, 0, err
	}
	defer conn.Close()

	// Resolve any DNS (if used) and get the real IP of the target
	dst, err := net.ResolveIPAddr("ip4", addr)
	if err != nil {
		panic(err)
		return nil, 0, err
	}

	// Make a new ICMP message
	m := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: []byte(""),
		},
	}
	b, err := m.Marshal(nil)
	if err != nil {
		return dst, 0, err
	}

	// Send it
	start := time.Now()
	n, err := conn.WriteTo(b, dst)
	if err != nil {
		return dst, 0, err
	} else if n != len(b) {
		return dst, 0, fmt.Errorf("got %v; want %v", n, len(b))
	}

	// Wait for a reply
	reply := make([]byte, 1500)
	err = conn.SetReadDeadline(time.Now().Add(TimeOUT))
	if err != nil {
		return dst, 0, err
	}
	n, peer, err := conn.ReadFrom(reply)
	if err != nil {
		return dst, 0, err
	}
	duration := time.Since(start)

	// Pack it up boys, we're done here
	rm, err := icmp.ParseMessage(ProtocolICMP, reply[:n])
	if err != nil {
		return dst, 0, err
	}
	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		return dst, duration, nil
	default:
		return dst, 0, fmt.Errorf("got %+v from %v; want echo reply", rm, peer)
	}
}

func main() {
	p := func(addr string) {
		dst, dur, err := Ping(addr)
		if err != nil {
			log.Printf("Ping %s (%s): %s\n", addr, dst, err)
			return
		}
		log.Printf("Ping %s (%s): %s\n", addr, dst, dur)
	}
	p("127.0.0.1")
	p("172.27.0.1")

}
