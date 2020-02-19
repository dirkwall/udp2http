package main

import (
	"log"
	"net"
	"fmt"
	"net/http"
	"bytes"
)

func main() {
	pc, err := net.ListenPacket("udp", ":1053")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	for {
		buf := make([]byte, 1472)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}
		go serve(pc, addr, buf[:n])
	}
}

func serve(pc net.PacketConn, addr net.Addr, buf []byte) {
	fmt.Printf("datagram received: bytes %d from %s\n", len(buf), addr.String());

	resp, err := http.Post("http://127.0.0.1:51000", "text/html", bytes.NewBuffer(buf));
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("status: ", resp.Status)
	resp.Body.Close();

	pc.WriteTo([]byte("ACK"), addr)
}
