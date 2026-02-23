// Package main runs the main code
package main

import (
	"fmt"
	"log"
	"navit/projects/betterChat/util"
	"net"
	"unicode/utf8"
)

const (
	PORT = "6969"
)

// MainServer listens to incoming requests and
// places each request into the correct room
func MainServer(message chan util.Message) {
	clientsInMain := map[string]*util.Client{}
	for {
		mes := <-message
		addr := mes.Conn.RemoteAddr().(*net.TCPAddr)
		switch mes.Type {
		case util.ClientConnectedtoMain:
			log.Printf("Client %s connected to the main server!\n", mes.Conn.RemoteAddr())
			clientsInMain[addr.String()] = &util.Client{
				Conn: mes.Conn,
			}
			log.Printf("Here is the list of Rooms you can join %s", mes.Conn.RemoteAddr())

		case util.NewMessage:
			authorAddr := mes.Conn.RemoteAddr()
			// author := clientsInMain[author_addr.String()]
			if utf8.Valid([]byte(mes.Text)) {
				if addr.String() != authorAddr.String() {
					_, err := mes.Conn.Write([]byte(mes.Text))
					if err != nil {
						// TODO: remove the connection from the list
						fmt.Printf("Could not send data to %s because %s", mes.Conn.RemoteAddr(), err)
					}
					log.Printf("Client %s sent message %s\n", mes.Conn.RemoteAddr(), mes.Text)
				}
			}
		case util.ClientDisconnectedfromMain:
			log.Printf("Client %s disconnected from the main server", addr)
			delete(clientsInMain, addr.String())
		}
	}
}

func client(conn net.Conn, messages chan util.Message) {
	buffer := make([]byte, 512)
	for {
		n, err := conn.Read([]byte(buffer))
		if err != nil {
			log.Printf("Could not read from %s", conn.RemoteAddr())
			messages <- util.Message{
				Type: util.ClientDisconnectedfromMain,
				Text: string(buffer[0:n]),
				Conn: conn,
			}
			err := conn.Close()
			if err != nil {
				fmt.Printf("could not close connection: %s\n", err)
			}
			return
		}
		text := string(buffer[0:n])

		messages <- util.Message{
			Type: util.NewMessage,
			Text: text,
			Conn: conn,
		}
	}
}

func main() {
	// Display all existing rooms
	all := util.AllRooms{}
	for _, n := range all.AllRooms {
		fmt.Printf("Room name: %s\n", n.Name)
	}
	ln, err := net.Listen("tcp", ":"+PORT)
	fmt.Printf("listening on port %s\n", PORT)
	if err != nil {
		fmt.Printf("Could not Listen at epic port %s cause.. %s", PORT, err)
		return
	}
	message := make(chan util.Message)
	go MainServer(message)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Could not accept connection due to: %s", err)
		}
		fmt.Printf("Accepted Connection from: %s", conn.RemoteAddr())
		message <- util.Message{
			Type: util.ClientConnectedtoMain,
			Conn: conn,
		}
		go client(conn, message)
	}
}
