package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"

	"go.fmi/issuetracker/command"
	"go.fmi/issuetracker/db"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		log.Fatalln(err)
	}

	db.Connect()
	defer listener.Close()

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// If you want, you can increment a counter here and inject to handleClientRequest below as client identifier
		go handleClientRequest(con)
	}
}

func handleClientRequest(con net.Conn) {
	defer con.Close()

	clientReader := bufio.NewReader(con)

	for {
		// Waiting for the client request
		clientRequest, err := clientReader.ReadString('\n')

		switch err {
		case nil:
			clientRequest := strings.TrimSpace(clientRequest)
			if clientRequest == "disconnect" {
				log.Println("Connection to server is closed")
				return
			}

			parsedCommand := command.ParseCommand(clientRequest)
			message, _ := parsedCommand.Execute()
			con.Write([]byte(message))
		case io.EOF:
			log.Println("Client closed the connection by terminating the process")
			return
		default:
			log.Printf("Error: %v\n", err)
			return
		}
	}
}
