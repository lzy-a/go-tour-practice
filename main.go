package main

import (
	"io"
	"log"
	"net"
	"os"
	"server/handler"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	// go client()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// go handler.PrintTime(conn)
		go handler.Echo(conn)
	}
}
func client() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Print(err)
		return
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
