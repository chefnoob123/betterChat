// Package main runs the main code
package main

import (
	"fmt"
	"net"

	"navit/projects/betterChat/util"
)

const (
	PORT = "69420"
)

// Server listens to incoming requests inside a room
// So each room has its own server
func server(message chan util.Message) {
}

func main() {
	// Display all existing rooms
	all := util.AllRooms{}
	for _, n := range all.AllRooms {
		fmt.Printf("Room name: %s\n", n.Name)
	}
	ln, err := net.Listen("tcp", PORT)
	fmt.Printf("listening on port %s\n", PORT)
	if err != nil {
		fmt.Printf("Could not Listen at epic port %s cause.. %s", PORT, err)
	}
	message := make(chan util.Message)
}
