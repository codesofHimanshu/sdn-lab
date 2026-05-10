package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Name string
	Conn *websocket.Conn
}

var clients = make(map[string]*Client)
var mutex sync.Mutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Secure Chat Server Running")
	})

	http.HandleFunc("/ws", handleConnections)

	server := &http.Server{
		Addr: ":4433",
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
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

func handleConnections(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	_, usernameBytes, err := conn.ReadMessage()
	if err != nil {
		return
	}

	username := strings.TrimSpace(string(usernameBytes))

	client := &Client{
		Name: username,
		Conn: conn,
	}

	mutex.Lock()
	clients[username] = client
	mutex.Unlock()

	log.Println(username + " connected")

	broadcast(
		"SERVER: "+username+" joined the chat",
		username,
	)

	for {

		_, msgBytes, err := conn.ReadMessage()

		if err != nil {

			mutex.Lock()
			delete(clients, username)
			mutex.Unlock()

			log.Println(username + " disconnected")

			broadcast(
				"SERVER: "+username+" left the chat",
				username,
			)

			return
		}

		message := string(msgBytes)

		log.Println(username + ": " + message)

		if strings.HasPrefix(message, "@") {

			parts := strings.SplitN(message, " ", 2)

			if len(parts) < 2 {
				continue
			}

			target := strings.TrimPrefix(parts[0], "@")
			privateMessage := parts[1]

			sendPrivate(
				username,
				target,
				privateMessage,
			)

		} else {

			broadcast(
				username+": "+message,
				username,
			)
		}
	}
}

func broadcast(message string, sender string) {

	mutex.Lock()
	defer mutex.Unlock()

	for username, client := range clients {

		if username == sender {
			continue
		}

		client.Conn.WriteMessage(
			websocket.TextMessage,
			[]byte(message),
		)
	}
}

func sendPrivate(sender string, target string, message string) {

	mutex.Lock()
	defer mutex.Unlock()

	client, exists := clients[target]

	if !exists {
		return
	}

	privateText :=
		"[PRIVATE] " + sender + ": " + message

	client.Conn.WriteMessage(
		websocket.TextMessage,
		[]byte(privateText),
	)
}
