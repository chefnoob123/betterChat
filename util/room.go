// Package util provides utils for chat rooms
package util

import (
	"net"
)

type messageType int

const (
	ClientConnected messageType = iota + 1 // iota basically makes the values int the enums numeric like 1, 2, 3 etc
	ClientDisconnected
	NewMessage
)

type Message struct {
	Type messageType
	Conn net.Conn
	Text string
}

type Client struct {
	Conn net.Conn
}
type Rooms struct {
	Name    string
	ID      string
	Members map[*Client]net.Conn
}

type AllRooms struct {
	AllRooms []Rooms
}
type Room interface {
	enter()
	kick()
}
