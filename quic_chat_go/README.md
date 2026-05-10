# QUIC Multi Client Chat

A secure multi-client real-time chat application built using Go, WebSocket/WebTransport concepts, HTTPS, and JavaScript frontend.

## Features

- Multi-client chat
- Real-time messaging
- Username-based login
- Broadcast messages
- Private messaging using `@username`
- Secure HTTPS communication
- Modern UI

## Tech Stack

- Go
- Gorilla WebSocket
- HTML
- CSS
- JavaScript

## Project Structure

```bash
quic_chat_go/
│
├── certs/
├── frontend/
├── server/
├── go.mod
└── README.md
```

## Run the Backend

```bash
go run ./server
```

Server runs on:

```bash
https://localhost:4433
```

## Run the Frontend

```bash
npx http-server frontend -S -C certs/localhost.pem -K certs/localhost-key.pem -p 5500
```

Frontend URL:

```bash
https://localhost:5500
```

## Usage

1. Open multiple browser tabs.
2. Join with different usernames.
3. Send broadcast messages.
4. Send private messages using:

```bash
@username your_message
```

## Notes

- HTTPS certificates are required for secure browser communication.
- Built for Software Defined Networking (SDN) lab experimentation.

## Author

Himanshu Dihingia