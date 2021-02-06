package command

import (
	"strings"

	"go.fmi/issuetracker/db"
	"go.fmi/issuetracker/issue"
	"go.fmi/issuetracker/project"
	"go.fmi/issuetracker/user"
)

// Command is a common interface for all supported command types
type Command interface {
	Execute() (string, bool)
}

// ParseCommand is a factory function that instantiates a Command using raw input
func ParseCommand(rawCommand string) Command {
	commandElements := strings.Split(rawCommand, "|-|")
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
	case "resolve":
		return ResolveCommand{
			Project: commandElements[1],
			Title:   commandElements[2]}
	case "list":
		return ListCommand{
			Project: commandElements[1]}
	case "find":
		return FindCommand{
			Project: commandElements[1],
			Title:   commandElements[2]}
	default:
		return nil
	}
}

// REGISTER

// RegisterCommand is used to create a new user
type RegisterCommand struct {
	User user.User
}

// Execute creates a new user
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

// LoginCommand is used to log a user in to their account
type LoginCommand struct {
	User user.User
}

// Execute logs a user in to their account
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

// ProjectCommand is used to create a new project
type ProjectCommand struct {
	Project project.Project
}

// Execute creates a new project
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

// IssueCommand is used to create a new issue in a project
type IssueCommand struct {
	Issue issue.Issue
}

// Execute creates a new issue in a project
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

// RESOLVE

// ResolveCommand is use to resolve an issue
type ResolveCommand struct {
	Project string
	Title   string
}

// Execute resolves an issue
func (rc ResolveCommand) Execute() (string, bool) {
	if _, err := db.FindExistingProject(rc.Project); err != nil {
		return "Could not find project \n", false
	}

	resolvableIssue, err := db.FindExistingIssue(rc.Project, rc.Title)
	if err != nil {
		return "Could not resolve issue - issue does not exist \n", false
	}

	if resolvableIssue.Resolved == "true" {
		return "Issue is already resolved \n", false
	}

	db.ResolveIssue(resolvableIssue.Project, resolvableIssue.Title)
	return "Issue resolved successfully\n", true
}

// LIST

// ListCommand is used to list all issues in a project
type ListCommand struct {
	Project string
}

// Execute lists all issues in a project
func (lc ListCommand) Execute() (string, bool) {
	if _, err := db.FindExistingProject(lc.Project); err != nil {
		return "Could not find project \n", false
	}

	issues := db.ListIssues(lc.Project)
	if len(issues) == 0 {
		return "There aren't any issues in this project\n", true
	}

	issuesTitles := "Issues in project: "
	for _, issue := range issues {
		issuesTitles += issue.Title + ", "
	}

	return issuesTitles[:len(issuesTitles)-2] + "\n", true
}

// FIND

// FindCommand is used to find the details for an issue in a project
type FindCommand struct {
	Project string
	Title   string
}

// Execute finds the details for an issue in a project
func (fc FindCommand) Execute() (string, bool) {
	if _, err := db.FindExistingProject(fc.Project); err != nil {
		return "Could not find project \n", false
	}

	foundIssue, err := db.FindExistingIssue(fc.Project, fc.Title)
	if err != nil {
		return "Issue does not exist \n", false
	}

	foundIssueStr := "Project: " + foundIssue.Project + " Reporter: " +
		foundIssue.Reporter + " Title: " + foundIssue.Title + " Description: " +
		foundIssue.Description + " Resolved: " + foundIssue.Resolved + "\n"

	return foundIssueStr, true
}
