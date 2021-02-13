package command

import (
	"errors"
	"testing"

	"go.fmi/issuetracker/comment"
	"go.fmi/issuetracker/db"
	"go.fmi/issuetracker/issue"
	"go.fmi/issuetracker/project"
	"go.fmi/issuetracker/user"

	"bou.ke/monkey"
)

func TestParserRegisterCommand(t *testing.T) {
	rawCommand := "register|-|user|-|password"
	parsedCommand := ParseCommand(rawCommand)
	expectedCommand := RegisterCommand{
		user.User{
			Username: "user",
			Password: "password"}}

	switch parsedCommand.(type) {
	case RegisterCommand:
		if parsedCommand != expectedCommand {
			t.Errorf("Invalid parsing: command parameters were not properly assigned")
		}
	default:
		t.Errorf("Invalid command type: expected RegisterCommand")
	}
}

func TestParseLoginCommand(t *testing.T) {
	rawCommand := "login|-|user|-|password"
	parsedCommand := ParseCommand(rawCommand)
	expectedCommand := LoginCommand{
		user.User{
			Username: "user",
			Password: "password"}}

	switch parsedCommand.(type) {
	case LoginCommand:
		if parsedCommand != expectedCommand {
			t.Errorf("Invalid parsing: command parameters were not properly assigned")
		}
	default:
		t.Errorf("Invalid command type: expected LoginCommand")
	}
}

func TestParseProjectCommand(t *testing.T) {
	rawCommand := "project|-|name"
	parsedCommand := ParseCommand(rawCommand)
	expectedCommand := ProjectCommand{
		project.Project{
			Name: "name"}}

	switch parsedCommand.(type) {
	case ProjectCommand:
		if parsedCommand != expectedCommand {
			t.Errorf("Invalid parsing: command parameters were not properly assigned")
		}
	default:
		t.Errorf("Invalid command type: expected ProjectCommand")
	}
}

func TestParseIssueCommand(t *testing.T) {
	rawCommand := "issue|-|name|-|user|-|title|-|description|-|false"
	parsedCommand := ParseCommand(rawCommand)
	expectedCommand := IssueCommand{
		issue.Issue{
			Project:     "name",
			Reporter:    "user",
			Title:       "title",
			Description: "description",
			Resolved:    "false"}}

	switch parsedCommand.(type) {
	case IssueCommand:
		if parsedCommand != expectedCommand {
			t.Errorf("Invalid parsing: command parameters were not properly assigned")
		}
	default:
		t.Errorf("Invalid command type: expected IssueCommand")
	}
}

func TestParseResolveCommand(t *testing.T) {
	rawCommand := "resolve|-|name|-|title"
	parsedCommand := ParseCommand(rawCommand)
	expectedCommand := ResolveCommand{
		Project: "name",
		Title:   "title"}

	switch parsedCommand.(type) {
	case ResolveCommand:
		if parsedCommand != expectedCommand {
			t.Errorf("Invalid parsing: command parameters were not properly assigned")
		}
	default:
		t.Errorf("Invalid command type: expected ResolveCommand")
	}
}

func TestParseListCommand(t *testing.T) {
	rawCommand := "list|-|name"
	parsedCommand := ParseCommand(rawCommand)
	expectedCommand := ListCommand{
		Project: "name"}

	switch parsedCommand.(type) {
	case ListCommand:
		if parsedCommand != expectedCommand {
			t.Errorf("Invalid parsing: command parameters were not properly assigned")
		}
	default:
		t.Errorf("Invalid command type: expected ListCommand")
	}
}

func TestParseFindCommand(t *testing.T) {
	rawCommand := "find|-|name|-|title"
	parsedCommand := ParseCommand(rawCommand)
	expectedCommand := FindCommand{
		Project: "name",
		Title:   "title"}

	switch parsedCommand.(type) {
	case FindCommand:
		if parsedCommand != expectedCommand {
			t.Errorf("Invalid parsing: command parameters were not properly assigned")
		}
	default:
		t.Errorf("Invalid command type: expected FindCommand")
	}
}

func TestParseCommentCommand(t *testing.T) {
	rawCommand := "comment|-|name|-|title|-|content|-|commenter"
	parsedCommand := ParseCommand(rawCommand)
	expectedCommand := CommentCommand{
		comment.Comment{
			Project:   "name",
			Title:     "title",
			Content:   "content",
			Commenter: "commenter"}}

	switch parsedCommand.(type) {
	case CommentCommand:
		if parsedCommand != expectedCommand {
			t.Errorf("Invalid parsing: command parameters were not properly assigned")
		}
	default:
		t.Errorf("Invalid command type: expected CommentCommand")
	}
}

func TestRegisterUserExisting(t *testing.T) {
	defer monkey.UnpatchAll()
	userMock := user.User{
		Username: "user",
		Password: "password"}

	monkey.Patch(db.FindRegisteredUser, func(string) (user.User, error) {
		return userMock, nil
	})

	registerCommand := RegisterCommand{userMock}
	message, ok := registerCommand.Execute()
	if ok {
		t.Errorf("Command execution completed with OK, but shouldn't have")
	}

	if message != "Registration unsuccessful - username is not unique\n" {
		t.Errorf("Invalid command execution message. Expected: Registration unsuccessful - username is not unique\n, but got " + message)
	}
}

func TestRegisterUserUnique(t *testing.T) {
	defer monkey.UnpatchAll()
	userMock := user.User{
		Username: "user",
		Password: "password"}

	monkey.Patch(db.FindRegisteredUser, func(string) (user.User, error) {
		return user.User{}, errors.New("User not found")
	})

	monkey.Patch(db.InsertRegisteredUser, func(user.User) {
		return
	})

	registerCommand := RegisterCommand{userMock}
	message, ok := registerCommand.Execute()
	if !ok {
		t.Errorf("Command execution did not complet with OK, but should have")
	}

	if message != "Registration successful. You are now logged in as user\n" {
		t.Errorf("Invalid command execution message. Expected: Registration successful. You are now logged in as user\n, but got " + message)
	}
}

func TestLoginExisting(t *testing.T) {
	defer monkey.UnpatchAll()
	userMock := user.User{
		Username: "user",
		Password: "password1234"}

	monkey.Patch(db.FindRegisteredUser, func(string) (user.User, error) {
		return userMock, nil
	})

	monkey.Patch(user.ComparePasswords, func(string, string) bool {
		return true
	})

	loginCommand := LoginCommand{userMock}
	message, ok := loginCommand.Execute()
	if !ok {
		t.Errorf("Command execution did not complete with OK, but should have")
	}

	if message != "Login successful as user\n" {
		t.Errorf("Invalid command execution message. Expected: Login successful as user\n, but got " + message)
	}
}

func TestLoginMissingUser(t *testing.T) {
	defer monkey.UnpatchAll()
	userMock := user.User{
		Username: "user",
		Password: "password1234"}

	monkey.Patch(db.FindRegisteredUser, func(string) (user.User, error) {
		return user.User{}, errors.New("User doesn't exist")
	})

	monkey.Patch(user.ComparePasswords, func(string, string) bool {
		return false
	})

	loginCommand := LoginCommand{userMock}
	message, ok := loginCommand.Execute()
	if ok {
		t.Errorf("Command execution completed with OK, but shouldn't have")
	}

	if message != "Login unsuccessful - inavlid username/password\n" {
		t.Errorf("Invalid command execution message. Expected: Login unsuccessful - inavlid username/password\n, but got " + message)
	}
}

func TestLoginWrongPassword(t *testing.T) {
	defer monkey.UnpatchAll()
	userMock := user.User{
		Username: "user",
		Password: "password1234"}

	monkey.Patch(db.FindRegisteredUser, func(string) (user.User, error) {
		return userMock, nil
	})

	monkey.Patch(user.ComparePasswords, func(string, string) bool {
		return false
	})

	loginCommand := LoginCommand{userMock}
	message, ok := loginCommand.Execute()
	if ok {
		t.Errorf("Command execution completed with OK, but shouldn't have")
	}

	if message != "Login unsuccessful - inavlid username/password\n" {
		t.Errorf("Invalid command execution message. Expected: Login unsuccessful - inavlid username/password\n, but got " + message)
	}
}

func TestCreateExistingProject(t *testing.T) {
	defer monkey.UnpatchAll()
	projectMock := project.Project{
		Name: "project"}

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return projectMock, nil
	})

	projectCommand := ProjectCommand{projectMock}
	message, ok := projectCommand.Execute()
	if ok {
		t.Errorf("Command execution completed with OK, but shouldn't have")
	}

	if message != "Could not create new project - project name is not unique\n" {
		t.Errorf("Invalid command execution message. Expected: Could not create new project - project name is not unique\n, but got " + message)
	}
}

func TestCreateUniqueProject(t *testing.T) {
	defer monkey.UnpatchAll()
	projectMock := project.Project{
		Name: "project"}

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return project.Project{}, errors.New("Project does not exist")
	})

	monkey.Patch(db.InsertNewProject, func(project.Project) {
		return
	})

	projectCommand := ProjectCommand{projectMock}
	message, ok := projectCommand.Execute()
	if !ok {
		t.Errorf("Command execution didn't complete with OK, but should have")
	}

	if message != "Project created successfully\n" {
		t.Errorf("Invalid command execution message. Expected: Project created successfully\n, but got " + message)
	}
}

func TestCreateUniqueIssue(t *testing.T) {
	defer monkey.UnpatchAll()
	issueMock := issue.Issue{
		Project:     "project",
		Reporter:    "reporter",
		Title:       "title",
		Description: "description",
		Resolved:    "false"}

	projectMock := project.Project{
		Name: "project"}

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return projectMock, nil
	})

	monkey.Patch(db.FindExistingIssue, func(string, string) (issue.Issue, error) {
		return issue.Issue{}, errors.New("Issue not found")
	})

	monkey.Patch(db.InsertNewIssue, func(issue.Issue) {
		return
	})

	issueCommand := IssueCommand{issueMock}
	message, ok := issueCommand.Execute()
	if !ok {
		t.Errorf("Command execution didn't complete with OK, but should have")
	}

	if message != "Issue created successfully\n" {
		t.Errorf("Invalid command execution message. Expected: Issue created successfully\n, but got " + message)
	}
}

func TestCreateNonUniqueIssue(t *testing.T) {
	defer monkey.UnpatchAll()
	issueMock := issue.Issue{
		Project:     "project",
		Reporter:    "reporter",
		Title:       "title",
		Description: "description",
		Resolved:    "false"}

	projectMock := project.Project{
		Name: "project"}

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return projectMock, nil
	})

	monkey.Patch(db.FindExistingIssue, func(string, string) (issue.Issue, error) {
		return issueMock, nil
	})

	issueCommand := IssueCommand{issueMock}
	message, ok := issueCommand.Execute()
	if ok {
		t.Errorf("Command execution completed with OK, but shouldn't have")
	}

	if message != "Could not create new issue - issue name is not unique for project\n" {
		t.Errorf("Invalid command execution message. Expected: Could not create new issue - issue name is not unique for project\n, but got " + message)
	}
}

func TestCreateIssueMissingProject(t *testing.T) {
	defer monkey.UnpatchAll()

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return project.Project{}, errors.New("Project not found")
	})

	issueCommand := IssueCommand{issue.Issue{}}
	message, ok := issueCommand.Execute()
	if ok {
		t.Errorf("Command execution completed with OK, but shouldn't have")
	}

	if message != "Could not find project \n" {
		t.Errorf("Invalid command execution message. Expected: Could not find project \n, but got " + message)
	}
}

func TestResolveIssue(t *testing.T) {
	defer monkey.UnpatchAll()
	issueMock := issue.Issue{
		Project:     "project",
		Reporter:    "reporter",
		Title:       "title",
		Description: "description",
		Resolved:    "false"}

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return project.Project{}, nil
	})

	monkey.Patch(db.FindExistingIssue, func(string, string) (issue.Issue, error) {
		return issueMock, nil
	})

	monkey.Patch(db.ResolveIssue, func(string, string) {
		return
	})

	resolveCommand := ResolveCommand{Project: "project", Title: "title"}
	message, ok := resolveCommand.Execute()

	if !ok {
		t.Errorf("Command execution didn't complete with OK, but should have")
	}

	if message != "Issue resolved successfully\n" {
		t.Errorf("Invalid command execution message. Expected: Issue resolved successfully\n, but got " + message)
	}
}

func TestResolveMissingIssue(t *testing.T) {
	defer monkey.UnpatchAll()

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return project.Project{}, nil
	})

	monkey.Patch(db.FindExistingIssue, func(string, string) (issue.Issue, error) {
		return issue.Issue{}, errors.New("Issue not found")
	})

	resolveCommand := ResolveCommand{Project: "project", Title: "title"}
	message, ok := resolveCommand.Execute()

	if ok {
		t.Errorf("Command execution completed with OK, but shouldn't have")
	}

	if message != "Could not resolve issue - issue does not exist \n" {
		t.Errorf("Invalid command execution message. Expected: Could not resolve issue - issue does not exist \n, but got " + message)
	}
}

func TestResolveIssueMissingProject(t *testing.T) {
	defer monkey.UnpatchAll()

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return project.Project{}, errors.New("Project not found")
	})

	resolveCommand := ResolveCommand{Project: "project", Title: "title"}
	message, ok := resolveCommand.Execute()

	if ok {
		t.Errorf("Command execution completed with OK, but shouldn't have")
	}

	if message != "Could not find project \n" {
		t.Errorf("Invalid command execution message. Expected: Could not find project \n, but got " + message)
	}
}

func TestResolveResolvedIssue(t *testing.T) {
	defer monkey.UnpatchAll()
	issueMock := issue.Issue{
		Project:     "project",
		Reporter:    "reporter",
		Title:       "title",
		Description: "description",
		Resolved:    "true"}

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return project.Project{}, nil
	})

	monkey.Patch(db.FindExistingIssue, func(string, string) (issue.Issue, error) {
		return issueMock, nil
	})

	resolveCommand := ResolveCommand{Project: "project", Title: "title"}
	message, ok := resolveCommand.Execute()

	if ok {
		t.Errorf("Command execution completed with OK, but shouldn't have")
	}

	if message != "Issue is already resolved \n" {
		t.Errorf("Invalid command execution message. Expected: Issue is already resolved \n, but got " + message)
	}
}

func TestListCommand(t *testing.T) {
	defer monkey.UnpatchAll()

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return project.Project{}, nil
	})

	monkey.Patch(db.ListIssues, func(string) []issue.Issue {
		return []issue.Issue{issue.Issue{Title: "first issue"}, issue.Issue{Title: "second issue"}}
	})

	listCommand := ListCommand{Project: "project"}
	message, ok := listCommand.Execute()

	if !ok {
		t.Errorf("Command execution didn't complete with OK, but should have")
	}

	if message != "Issues in project: first issue, second issue\n" {
		t.Errorf("Invalid command execution message. Expected: Issues in project: first issue, second issue\n, but got " + message)
	}
}

func TestListCommandNoIssues(t *testing.T) {
	defer monkey.UnpatchAll()

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return project.Project{}, nil
	})

	monkey.Patch(db.ListIssues, func(string) []issue.Issue {
		return []issue.Issue{}
	})

	listCommand := ListCommand{Project: "project"}
	message, ok := listCommand.Execute()

	if !ok {
		t.Errorf("Command execution didn't complete with OK, but should have")
	}

	if message != "There aren't any issues in this project\n" {
		t.Errorf("Invalid command execution message. Expected: There aren't any issues in this project\n, but got " + message)
	}
}

func TestListCommandMissingProject(t *testing.T) {
	defer monkey.UnpatchAll()

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return project.Project{}, errors.New("Project not found")
	})

	listCommand := ListCommand{Project: "project"}
	message, ok := listCommand.Execute()

	if ok {
		t.Errorf("Command execution completed with OK, but shouldn't have")
	}

	if message != "Could not find project \n" {
		t.Errorf("Invalid command execution message. Expected: Could not find project \n, but got " + message)
	}
}

func TestFindCommand(t *testing.T) {
	defer monkey.UnpatchAll()
	commentMock := comment.Comment{
		Project:   "project",
		Title:     "title",
		Content:   "content",
		Commenter: "commenter"}
	issueMock := issue.Issue{
		Project:     "project",
		Reporter:    "reporter",
		Title:       "title",
		Description: "description",
		Resolved:    "true"}

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return project.Project{Name: "project"}, nil
	})

	monkey.Patch(db.FindExistingIssue, func(string, string) (issue.Issue, error) {
		return issueMock, nil
	})

	monkey.Patch(db.FindComments, func(string, string) []comment.Comment {
		return []comment.Comment{commentMock}
	})

	findCommand := FindCommand{Project: "project", Title: "title"}
	message, ok := findCommand.Execute()

	if !ok {
		t.Errorf("Command execution didn't complete with OK, but should have")
	}

	if message != "Project: project; Reporter: reporter; Title: title; Description: description; Resolved: true; Comments: \"content\" - commenter;\n" {
		t.Errorf("Invalid command execution message. Expected: Project: project; Reporter: reporter; Title: title; Description: description; Resolved: true; Comments: \"content\" - commenter;\n, but got " + message)
	}
}

func TestFindCommandMissingIssue(t *testing.T) {
	defer monkey.UnpatchAll()

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return project.Project{Name: "project"}, nil
	})

	monkey.Patch(db.FindExistingIssue, func(string, string) (issue.Issue, error) {
		return issue.Issue{}, errors.New("Missing issue")
	})

	findCommand := FindCommand{Project: "project", Title: "title"}
	message, ok := findCommand.Execute()

	if ok {
		t.Errorf("Command execution completed with OK, but shouldn't have")
	}

	if message != "Issue does not exist \n" {
		t.Errorf("Invalid command execution message. Expected: Issue does not exist \n, but got " + message)
	}
}

func TestFindCommandMissingProject(t *testing.T) {
	defer monkey.UnpatchAll()

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return project.Project{}, errors.New("Missing project")
	})

	findCommand := FindCommand{Project: "project", Title: "title"}
	message, ok := findCommand.Execute()

	if ok {
		t.Errorf("Command execution completed with OK, but shouldn't have")
	}

	if message != "Could not find project \n" {
		t.Errorf("Invalid command execution message. Expected: Could not find project \n, but got " + message)
	}
}

func TestCommentCommand(t *testing.T) {
	defer monkey.UnpatchAll()
	commentMock := comment.Comment{
		Project:   "project",
		Title:     "title",
		Content:   "content",
		Commenter: "commenter"}

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return project.Project{Name: "project"}, nil
	})

	monkey.Patch(db.FindExistingIssue, func(string, string) (issue.Issue, error) {
		return issue.Issue{}, nil
	})

	monkey.Patch(db.InsertComment, func(comment.Comment) {
		return
	})

	commentCommand := CommentCommand{commentMock}
	message, ok := commentCommand.Execute()

	if !ok {
		t.Errorf("Command execution didn't complete with OK, but should have")
	}

	if message != "Comment added successfully\n" {
		t.Errorf("Invalid command execution message. Expected: Comment added successfully\n, but got " + message)
	}
}

func TestCommentCommandMissingIssue(t *testing.T) {
	defer monkey.UnpatchAll()

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return project.Project{Name: "project"}, nil
	})

	monkey.Patch(db.FindExistingIssue, func(string, string) (issue.Issue, error) {
		return issue.Issue{}, errors.New("Issue not found")
	})

	commentCommand := CommentCommand{comment.Comment{}}
	message, ok := commentCommand.Execute()

	if ok {
		t.Errorf("Command execution completed with OK, but shouldn't have")
	}

	if message != "Issue does not exist \n" {
		t.Errorf("Invalid command execution message. Expected: Issue does not exist \n, but got " + message)
	}
}

func TestCommentCommandMissingProject(t *testing.T) {
	defer monkey.UnpatchAll()

	monkey.Patch(db.FindExistingProject, func(string) (project.Project, error) {
		return project.Project{}, errors.New("Project not found")
	})

	commentCommand := CommentCommand{comment.Comment{}}
	message, ok := commentCommand.Execute()

	if ok {
		t.Errorf("Command execution completed with OK, but shouldn't have")
	}

	if message != "Could not find project \n" {
		t.Errorf("Invalid command execution message. Expected: Could not find project \n, but got " + message)
	}
}
