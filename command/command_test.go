package command

import (
	"testing"

	"go.fmi/issuetracker/comment"
	"go.fmi/issuetracker/issue"
	"go.fmi/issuetracker/project"
	"go.fmi/issuetracker/user"
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
