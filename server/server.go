package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	for {
		conn, _ := ln.Accept()
		conns <- conn
	}

}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	reader := bufio.NewReader(client)
	for {
		msg, _ := reader.ReadString('\n')
		msgs <- Message{message: msg, sender: clientid}
	}
}

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Println(msg)
		fmt.Fprintln(conn, "OK")
	}
}
func main() {
	//Read in the network port we should listen on, from the commandline argument.
	//Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	ln, _ := net.Listen("tcp", *portPtr)
	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)
	id := 0
	//Start accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			clients[id] = conn
			go handleClient(conn, id, msgs)
			id++
		case msg := <-msgs:
			for i := 0; i < id; i++ {
				if !(msg.sender == i) {
					fmt.Fprintln(clients[i], msg.message)
				}
			}
		}
	}
}
