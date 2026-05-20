package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Type     string   `json:"type"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Message  string   `json:"message"`
	Users    []string `json:"users,omitempty"`
}

var users = map[string]string{
	"himanshu": "1234",
	"test":     "1234",
}

var clients = make(map[string]*websocket.Conn)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	http.HandleFunc("/ws", handleWS)

	server := &http.Server{
		Addr: ":4433",
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	fmt.Println("Server running on https://localhost:4433")

	err := server.ListenAndServeTLS(
		"certs/localhost.pem",
		"certs/localhost-key.pem",
	)

	if err != nil {
		log.Fatal(err)
	}
}

func handleWS(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	handleConnection(conn)
}

func handleConnection(conn *websocket.Conn) {

	defer conn.Close()

	var currentUser string

	for {

		_, data, err := conn.ReadMessage()

		if err != nil {

			fmt.Println(err)

			if currentUser != "" {

				delete(clients, currentUser)

				broadcastUsers()
			}

			break
		}

		var msg Message

		err = json.Unmarshal(data, &msg)

		if err != nil {

			fmt.Println(err)
			continue
		}

		// LOGIN

		if msg.Type == "login" {

			pass, exists := users[msg.Username]

			if !exists || pass != msg.Password {

				conn.WriteJSON(Message{
					Type:    "chat",
					Message: "Invalid credentials",
				})

				continue
			}

			currentUser = msg.Username

			clients[currentUser] = conn

			conn.WriteJSON(Message{
				Type: "login_success",
			})

			broadcastUsers()

			fmt.Println(currentUser, "connected")

			continue
		}

		// CHAT

		if msg.Type == "chat" {

			text := currentUser + ": " + msg.Message

			broadcast(text)
		}
	}
}

func broadcast(message string) {

	msg := Message{
		Type:    "chat",
		Message: message,
	}

	for _, conn := range clients {

		conn.WriteJSON(msg)
	}
}

func broadcastUsers() {

	var list []string

	for username := range clients {

		list = append(list, username)
	}

	msg := Message{
		Type:  "users",
		Users: list,
	}

	for _, conn := range clients {

		conn.WriteJSON(msg)
	}
}
