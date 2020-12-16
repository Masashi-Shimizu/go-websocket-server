package main

import (
	"./src/github.com/gorilla/websocket"
)

type Client struct {
	connection *websocket.Conn
	sessionId string
}
