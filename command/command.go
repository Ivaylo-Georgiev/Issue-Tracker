package command

import (
	"strings"

	"go.fmi/issuetracker/db"
	"go.fmi/issuetracker/issue"
	"go.fmi/issuetracker/project"
	"go.fmi/issuetracker/user"
)

type Command interface {
	Execute() (string, bool)
}

func ParseCommand(rawCommand string) Command {
	commandElements := strings.Fields(rawCommand)
	commandType := commandElements[0]

	switch commandType {
	case "register":
		return RegisterCommand{user.User{
			Username: commandElements[1],
			Password: commandElements[2]}}
	case "login":
		return LoginCommand{
			user.User{
				Username: commandElements[1],
				Password: commandElements[2]}}
	case "project":
		return ProjectCommand{
			project.Project{
				Name: commandElements[1]}}
	case "issue":
		return IssueCommand{
			issue.Issue{
				Project:     commandElements[1],
				Reporter:    commandElements[2],
				Title:       commandElements[3],
				Description: commandElements[4],
				Resolved:    commandElements[5]}}
	default:
		return nil
	}
}

// REGISTER

type RegisterCommand struct {
	User user.User
}

func (rc RegisterCommand) Execute() (string, bool) {
	newUser := user.User{
		Username: rc.User.Username,
		Password: user.HashAndSalt(rc.User.Password)}

	if _, err := db.FindRegisteredUser(newUser.Username); err != nil {
		db.InsertRegisteredUser(newUser)
		return "Registration successful. You are now logged in as " + newUser.Username + "\n", true
	}
	return "Registration unsuccessful - username is not unique\n", false
}

// LOGIN

type LoginCommand struct {
	User user.User
}

func (lc LoginCommand) Execute() (string, bool) {
	loggingUser := user.User{
		Username: lc.User.Username,
		Password: lc.User.Password}

	registeredUser, err := db.FindRegisteredUser(loggingUser.Username)
	isPasswordCorrect := user.ComparePasswords(registeredUser.Password, loggingUser.Password)

	if err == nil && isPasswordCorrect {
		return "Login successful as " + loggingUser.Username + "\n", true
	} else {
		return "Login unsuccessful - inavlid username/password\n", false
	}
}

// PROJECT

type ProjectCommand struct {
	Project project.Project
}

func (pc ProjectCommand) Execute() (string, bool) {
	newProject := project.Project{
		Name: pc.Project.Name}

	if _, err := db.FindExistingProject(newProject.Name); err != nil {
		db.InsertNewProject(newProject)
		return "Project created successfully\n", true
	}

	return "Could not create new project - project name is not unique\n", false
}

// ISSUE

type IssueCommand struct {
	Issue issue.Issue
}

func (ic IssueCommand) Execute() (string, bool) {
	newIssue := issue.Issue{
		Project:     ic.Issue.Project,
		Reporter:    ic.Issue.Reporter,
		Title:       ic.Issue.Title,
		Description: ic.Issue.Description,
		Resolved:    ic.Issue.Resolved}

	if _, err := db.FindExistingProject(newIssue.Project); err != nil {
		return "Could not find project \n", false
	}

	if _, err := db.FindExistingIssue(newIssue.Project, newIssue.Title); err == nil {
		return "Could not create new issue - issue name is not unique for project\n", false
	}

	db.InsertNewIssue(newIssue)
	return "Issue created successfully\n", true
}
