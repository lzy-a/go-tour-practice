package handler

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func echoShout(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, shout)
	time.Sleep(delay)
	fmt.Fprintln(c, strings.ToLower(shout))
}

func Echo(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echoShout(c, input.Text(), time.Second)
	}
	c.Close()
}
