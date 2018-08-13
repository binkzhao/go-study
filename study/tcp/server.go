package main

import (
	"fmt"
	"net"
	"os"
)

var clients []net.Conn

func main() {
	var (
		host   = "localhost"
		port   = "8000"
		remote = host + ":" + port
		data   = make([]byte, 1024)
	)

	fmt.Println("Initiating server...")

	// listen on all interfaces
	lis, err := net.Listen("tcp", remote)
	defer lis.Close()
	if err != nil {
		fmt.Printf("Error When listen: %s, Err: %s", remote, err)
		os.Exit(1)
	}

	for {
		var res string
		// accept connection on port
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("Error accepting client: ", err.Error())
			os.Exit(0)
		}

		clients = append(clients, conn)

		go func(con net.Conn) {
			fmt.Println("New connection: ", con.RemoteAddr())

			// Get client's name
			length, err := con.Read(data)
			if err != nil {
				fmt.Printf("Client %v quit.\n", con.RemoteAddr())
				con.Close()
				disconnect(con, con.RemoteAddr().String())
				return
			}
			name := string(data[:length])
			comeStr := name + " entered the room."
			notify(con, comeStr)

			// Begin recieve message from client
			for {
				length, err := con.Read(data)
				if err != nil {
					fmt.Printf("Client %s quit.\n", name)
					con.Close()
					disconnect(con, name)
					return
				}
				res = string(data[:length])
				sprdMsg := name + " said: " + res
				fmt.Println(sprdMsg)
				res = "You said:" + res
				con.Write([]byte(res))
				notify(con, sprdMsg)
			}

		}(conn)
	}
}

// notify all clients
func notify(conn net.Conn, msg string)  {
	for _, curCon := range clients {
		if curCon.RemoteAddr() != conn.RemoteAddr() {
			curCon.Write([]byte(msg))
		}
	}
}

// one client has left
func disconnect(conn net.Conn, name string)  {
	for index, curCon := range clients {
		if curCon.RemoteAddr() == conn.RemoteAddr() {
			disMsg := fmt.Sprintf("【%s】has left the room", name)
			fmt.Println(disMsg)
			clients = append(clients[:index], clients[index+1:]...)
			notify(conn, disMsg)
		}
	}
}