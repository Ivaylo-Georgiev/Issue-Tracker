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

// LoggedUser is the user who is currently logged in the system
var LoggedUser string

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
				if LoggedUser != "" {
					LoggedUser = ""
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
			LoggedUser = serverResponseFields[len(serverResponseFields)-1]
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
		return ConstructLoginCommand()
	case "register":
		return ConstructRegisterCommand()
	case "project":
		return ConstructProjectCommand()
	case "issue":
		return ConstructIssueCommand()
	case "resolve":
		return ConstructResolveCommand()
	case "list":
		return ConstructListCommand()
	case "find":
		return ConstructFindCommand()
	case "comment":
		return ConstructCommentCommand()
	default:
		return "Invallid command", false
	}
}

// ConstructLoginCommand parses the user input for a login command into a string, which the server can handle
func ConstructLoginCommand() (string, bool) {
	if LoggedUser != "" {
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

// ConstructRegisterCommand parses the user input for a register command into a string, which the server can handle
func ConstructRegisterCommand() (string, bool) {
	if LoggedUser != "" {
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

// ConstructProjectCommand parses the user input for a project command into a string, which the server can handle
func ConstructProjectCommand() (string, bool) {
	if LoggedUser == "" {
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

// ConstructIssueCommand parses the user input for an issue command into a string, which the server can handle
func ConstructIssueCommand() (string, bool) {
	if LoggedUser == "" {
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

	return "issue|-|" + strings.TrimSpace(project) + "|-|" + LoggedUser + "|-|" + strings.TrimSpace(title) + "|-|" + strings.TrimSpace(description) + "|-|false", true
}

// ConstructResolveCommand parses the user input for a resolve command into a string, which the server can handle
func ConstructResolveCommand() (string, bool) {
	if LoggedUser == "" {
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

// ConstructListCommand parses the user input for a list command into a string, which the server can handle
func ConstructListCommand() (string, bool) {
	if LoggedUser == "" {
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

// ConstructFindCommand parses the user input for a find command into a string, which the server can handle
func ConstructFindCommand() (string, bool) {
	if LoggedUser == "" {
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

// ConstructCommentCommand parses the user input for a comment command into a string, which the server can handle
func ConstructCommentCommand() (string, bool) {
	if LoggedUser == "" {
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

	return "comment|-|" + strings.TrimSpace(project) + "|-|" + strings.TrimSpace(title) + "|-|" + strings.TrimSpace(comment) + "|-|" + LoggedUser, true
}
