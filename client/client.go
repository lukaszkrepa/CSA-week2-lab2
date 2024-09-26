package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n')
	fmt.Println(msg)

}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
}

func main() {

	//// Get the server address and port from the commandline arguments.
	//addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	//flag.Parse()
	////TODO Try to connect to the server
	////TODO Start asynchronously reading and displaying messages
	////TODO Start getting and sending user messages.

	stdin := bufio.NewReader(os.Stdin)
	conn, _ := net.Dial("tcp", "127.0.0.1:8030")

	for {
		fmt.Printf("Enter text->")
		msg, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, msg)
		read(conn)
	}
}
