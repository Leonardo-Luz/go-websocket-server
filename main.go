package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

type Server struct {
	clients map[*websocket.Conn]bool
}

func newServer() *Server {
	return &Server{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		//on message
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error: ", err)
			continue
		}

		msg := buf[:n]

		fmt.Println(string(msg))
		ws.Write([]byte("Thank you for the message!!\nmessage: " + string(msg)))
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	//on connection
	fmt.Println("new incoming connection fomr client: ", ws.RemoteAddr())

	s.clients[ws] = true

	s.readLoop(ws)
}

func main() {
	server := newServer()

	const PORT = ":8080"

	fmt.Println("Server Listenning at ", PORT)

	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.ListenAndServe(PORT, nil)
}
