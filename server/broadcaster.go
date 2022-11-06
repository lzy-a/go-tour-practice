package server

import (
	"bufio"
	"fmt"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string)
)

// 记录每个客户的信息，并发送相关内容
func Broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-message:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}

	}
}

func HandleConn(c net.Conn) {
	ch := make(chan string)
	go clientWriter(c, ch)
	who := c.RemoteAddr().String()
	ch <- "YOU ARE :" + who
	message <- who + " has arrived"
	entering <- ch
	input := bufio.NewScanner(c)
	for input.Scan() {
		message <- who + ": " + input.Text()
	}

	leaving <- ch
	message <- who + " has left"
	c.Close()
}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}

}
