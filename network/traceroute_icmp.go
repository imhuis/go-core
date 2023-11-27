package network

import (
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"log"
	"net"
	"os"
	"time"
)

func TraceRouteICMP() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s host", os.Args[0])
	}
	dst := os.Args[1]
	timeout := time.Second * 3
	// resolve the host name to an IP address
	ipAddr, err := net.ResolveIPAddr("ip4", dst)
	if err != nil {
		log.Fatalf("Failed to resolve IP address for %s: %v", dst, err)
	}
	// create a socket to listen for incoming ICMP packets
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatalf("Failed to create ICMP listener: %v", err)
	}
	defer conn.Close()
	id := os.Getpid() & 0xffff
	seq := 0
loop_ttl:
	for ttl := 1; ttl <= maxHops; ttl++ {
		// set the TTL on the socket
		if err := conn.IPv4PacketConn().SetTTL(ttl); err != nil {
			log.Fatalf("Failed to set TTL: %v", err)
		}
		seq++
		// create an ICMP message
		msg := icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Code: 0,
			Body: &icmp.Echo{
				ID:   id,
				Seq:  seq,
				Data: []byte("hello, are you there?"),
			},
		}
		// serialize the ICMP message
		msgBytes, err := msg.Marshal(nil)
		if err != nil {
			log.Fatalf("Failed to serialize ICMP message: %v", err)
		}
		// send the ICMP message
		start := time.Now()
		if _, err := conn.WriteTo(msgBytes, ipAddr); err != nil {
			log.Printf("%d: %v", ttl, err)
			continue loop_ttl
		}
		// listen for the reply
		replyBytes := make([]byte, 1500)
		if err := conn.SetReadDeadline(time.Now().Add(timeout)); err != nil {
			log.Fatalf("Failed to set read deadline: %v", err)
		}
		for i := 0; i < 3; i++ {
			n, peer, err := conn.ReadFrom(replyBytes)
			if err != nil {
				if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
					fmt.Printf("%d: *\n", ttl)
					continue loop_ttl
				} else {
					log.Printf("%d: Failed to parse ICMP message: %v", ttl, err)
				}
				continue
			}
			// parse the ICMP message
			replyMsg, err := icmp.ParseMessage(protocolICMP, replyBytes[:n])
			if err != nil {
				log.Printf("%d: Failed to parse ICMP message: %v", ttl, err)
				continue
			}
			// check if the reply is an echo reply
			if replyMsg.Type == ipv4.ICMPTypeEchoReply {
				echoReply, ok := msg.Body.(*icmp.Echo)
				if !ok || echoReply.ID != id || echoReply.Seq != seq {
					continue
				}
				fmt.Printf("%d: %v %v\n", ttl, peer, time.Since(start))
				break loop_ttl
			}
			if replyMsg.Type == ipv4.ICMPTypeTimeExceeded {
				echoReply, ok := msg.Body.(*icmp.Echo)
				if !ok || echoReply.ID != id || echoReply.Seq != seq {
					continue
				}
				var raddr = peer.String()
				names, _ := net.LookupAddr(raddr)
				if len(names) > 0 {
					raddr = names[0] + " (" + raddr + ")"
				} else {
					raddr = raddr + " (" + raddr + ")"
				}
				fmt.Printf("%d: %v %v\n", ttl, raddr, time.Since(start))
				continue loop_ttl
			}
		}
	}
}
