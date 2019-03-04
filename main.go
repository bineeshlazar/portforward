package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func copy(w io.WriteCloser, r io.Reader) {
	_, err := io.Copy(w, r)
	if err != nil {
		log.Println(err)
	}
	w.Close()
}

func forward(lc net.Conn, server string) {
	defer lc.Close()

	rc, err := net.Dial("tcp", server)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Forwarding connection to", server)

	go copy(rc, lc)
	copy(lc, rc)
	log.Printf("Terminated:  %s -> %s ", lc.RemoteAddr(), server)
}

func start(local, server string) {
	l, err := net.Listen("tcp", local)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	log.Println("Listening on ", local)

	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		log.Println("New connection from", conn.RemoteAddr())
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go forward(conn, server)
	}
}

func main() {

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf("Usage: %s <remote-addr:port> <local-addr:port>\n", os.Args[0])
		return
	}

	start(args[1], args[0])
}
