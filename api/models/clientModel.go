package models

import "github.com/gorilla/websocket"

type ClientModel struct {
	Conn *websocket.Conn
	UserName string
	RoomName string
}