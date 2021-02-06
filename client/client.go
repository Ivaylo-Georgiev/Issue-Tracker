package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var loggedUser string

func main() {
	con, err := net.Dial("tcp", "0.0.0.0:9999")
	if err != nil {
		log.Fatalln(err)
	}
	defer con.Close()

	clientReader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(con)

	for {
		// Waiting for the client request
		fmt.Print("Command: ")
		clientRequest, err := clientReader.ReadString('\n')

		switch err {
		case nil:
			clientRequest := strings.TrimSpace(clientRequest)

			if clientRequest == "logout" {
				if loggedUser != "" {
					loggedUser = ""
					log.Println("Successfully logged out")
				} else {
					log.Println("You are not logged in")
				}
				continue
			}

			if command, ok := constructCommand(clientRequest); ok {
				con.Write([]byte(command + "\n"))
			} else {
				log.Println(command)
				continue
			}
		case io.EOF:
			log.Println("Client closed the connection")
			return
		default:
			log.Printf("Client error: %v\n", err)
			return
		}

		// Waiting for the server response
		serverResponse, err := serverReader.ReadString('\n')
		if strings.Index(serverResponse, "Login successful") == 0 ||
			strings.Index(serverResponse, "Registration successful") == 0 {
			serverResponseFields := strings.Fields(serverResponse)
			loggedUser = serverResponseFields[len(serverResponseFields)-1]
		}

		switch err {
		case nil:
			log.Println(strings.TrimSpace(serverResponse))
		case io.EOF:
			log.Println("Server closed the connection")
			return
		default:
			log.Printf("Server error: %v\n", err)
			return
		}
	}
}

func constructCommand(clientRequest string) (string, bool) {
	switch clientRequest {
	case "disconnect":
		return "disconnect", true
	case "login":
		return constructLoginCommand()
	case "register":
		return constructRegisterCommand()
	case "project":
		return constructProjectCommand()
	case "issue":
		return constructIssueCommand()
	case "resolve":
		return constructResolveCommand()
	case "list":
		return constructListCommand()
	case "find":
		return constructFindCommand()
	case "comment":
		return constructCommentCommand()
	default:
		return "Invallid command", false
	}
}

func constructLoginCommand() (string, bool) {
	if loggedUser != "" {
		return "You are already logged in", false
	}

	scanner := bufio.NewScanner(os.Stdin)

	var username, password string
	fmt.Print("Username: ")
	if scanner.Scan() {
		username = scanner.Text()
	}
	fmt.Print("Password: ")
	if scanner.Scan() {
		password = scanner.Text()
	}

	return "login|-|" + strings.TrimSpace(username) + "|-|" + strings.TrimSpace(password), true
}

func constructRegisterCommand() (string, bool) {
	if loggedUser != "" {
		return "You are already logged in", false
	}

	scanner := bufio.NewScanner(os.Stdin)

	var username, password string
	fmt.Print("Username: ")
	if scanner.Scan() {
		username = scanner.Text()
	}
	fmt.Print("Password: ")
	if scanner.Scan() {
		password = scanner.Text()
	}

	return "register|-|" + strings.TrimSpace(username) + "|-|" + strings.TrimSpace(password), true
}

func constructProjectCommand() (string, bool) {
	if loggedUser == "" {
		return "You are not logged in", false
	}

	scanner := bufio.NewScanner(os.Stdin)

	var projectName string
	fmt.Print("Project name: ")
	if scanner.Scan() {
		projectName = scanner.Text()
	}

	return "project|-|" + strings.TrimSpace(projectName), true
}

func constructIssueCommand() (string, bool) {
	if loggedUser == "" {
		return "You are not logged in", false
	}

	scanner := bufio.NewScanner(os.Stdin)

	var project string
	fmt.Print("Project name: ")
	if scanner.Scan() {
		project = scanner.Text()
	}

	var title string
	fmt.Print("Title: ")
	if scanner.Scan() {
		title = scanner.Text()
	}

	var description string
	fmt.Print("Description: ")
	if scanner.Scan() {
		description = scanner.Text()
	}

	return "issue|-|" + strings.TrimSpace(project) + "|-|" + loggedUser + "|-|" + strings.TrimSpace(title) + "|-|" + strings.TrimSpace(description) + "|-|false", true
}

func constructResolveCommand() (string, bool) {
	if loggedUser == "" {
		return "You are not logged in", false
	}

	scanner := bufio.NewScanner(os.Stdin)

	var project string
	fmt.Print("Project name: ")
	if scanner.Scan() {
		project = scanner.Text()
	}

	var title string
	fmt.Print("Title: ")
	if scanner.Scan() {
		title = scanner.Text()
	}

	return "resolve|-|" + strings.TrimSpace(project) + "|-|" + strings.TrimSpace(title), true
}

func constructListCommand() (string, bool) {
	if loggedUser == "" {
		return "You are not logged in", false
	}

	scanner := bufio.NewScanner(os.Stdin)

	var project string
	fmt.Print("Project name: ")
	if scanner.Scan() {
		project = scanner.Text()
	}

	return "list|-|" + strings.TrimSpace(project), true
}

func constructFindCommand() (string, bool) {
	if loggedUser == "" {
		return "You are not logged in", false
	}

	scanner := bufio.NewScanner(os.Stdin)

	var project string
	fmt.Print("Project name: ")
	if scanner.Scan() {
		project = scanner.Text()
	}

	var title string
	fmt.Print("Title: ")
	if scanner.Scan() {
		title = scanner.Text()
	}

	return "find|-|" + strings.TrimSpace(project) + "|-|" + strings.TrimSpace(title), true
}

func constructCommentCommand() (string, bool) {
	if loggedUser == "" {
		return "You are not logged in", false
	}

	scanner := bufio.NewScanner(os.Stdin)

	var project string
	fmt.Print("Project name: ")
	if scanner.Scan() {
		project = scanner.Text()
	}

	var title string
	fmt.Print("Title: ")
	if scanner.Scan() {
		title = scanner.Text()
	}

	var comment string
	fmt.Print("Comment: ")
	if scanner.Scan() {
		comment = scanner.Text()
	}

	return "comment|-|" + strings.TrimSpace(project) + "|-|" + strings.TrimSpace(title) + "|-|" + strings.TrimSpace(comment) + "|-|" + loggedUser, true
}
