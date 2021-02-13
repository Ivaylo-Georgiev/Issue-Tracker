package main

import (
	"testing"
)

func TestConstructLoginCommandSkeletonLogged(t *testing.T) {
	LoggedUser = "test"
	expected := "You are already logged in"
	command, _ := ConstructLoginCommand()
	if command != expected {
		t.Errorf("Command skeleton was not costructed properly. Expected: " + expected + ", but got: " + command)
	}
}
func TestConstructLoginCommandSkeletonNotLogged(t *testing.T) {
	LoggedUser = ""
	expected := "login|-||-|"
	command, _ := ConstructLoginCommand()
	if command != expected {
		t.Errorf("Command skeleton was not costructed properly. Expected: " + expected + ", but got: " + command)
	}
}

func TestConstructRegisterCommandSkeletonLogged(t *testing.T) {
	LoggedUser = "test"
	expected := "You are already logged in"
	command, _ := ConstructRegisterCommand()
	if command != expected {
		t.Errorf("Command skeleton was not costructed properly. Expected: " + expected + ", but got: " + command)
	}
}

func TestConstructRegisterCommandSkeletonNotLogged(t *testing.T) {
	LoggedUser = ""
	expected := "register|-||-|"
	command, _ := ConstructRegisterCommand()
	if command != expected {
		t.Errorf("Command skeleton was not costructed properly. Expected: " + expected + ", but got: " + command)
	}
}

func TestConstructProjectCommandSkeletonLogged(t *testing.T) {
	LoggedUser = "test"
	expected := "project|-|"
	command, _ := ConstructProjectCommand()
	if command != expected {
		t.Errorf("Command skeleton was not costructed properly. Expected: " + expected + ", but got: " + command)
	}
}

func TestConstructProjectCommandSkeletonNotLogged(t *testing.T) {
	LoggedUser = ""
	expected := "You are not logged in"
	command, _ := ConstructProjectCommand()
	if command != expected {
		t.Errorf("Command skeleton was not costructed properly. Expected: " + expected + ", but got: " + command)
	}
}

func TestConstructIssueCommandSkeletonLogged(t *testing.T) {
	LoggedUser = "test"
	expected := "issue|-||-|test|-||-||-|false"
	command, _ := ConstructIssueCommand()
	if command != expected {
		t.Errorf("Command skeleton was not costructed properly. Expected: " + expected + ", but got: " + command)
	}
}

func TestConstructIssueCommandSkeletonNotLogged(t *testing.T) {
	LoggedUser = ""
	expected := "You are not logged in"
	command, _ := ConstructIssueCommand()
	if command != expected {
		t.Errorf("Command skeleton was not costructed properly. Expected: " + expected + ", but got: " + command)
	}
}

func TestConstructResolveCommandSkeletonLogged(t *testing.T) {
	LoggedUser = "test"
	expected := "resolve|-||-|"
	command, _ := ConstructResolveCommand()
	if command != expected {
		t.Errorf("Command skeleton was not costructed properly. Expected: " + expected + ", but got: " + command)
	}
}

func TestConstructResolveCommandSkeletonNotLogged(t *testing.T) {
	LoggedUser = ""
	expected := "You are not logged in"
	command, _ := ConstructResolveCommand()
	if command != expected {
		t.Errorf("Command skeleton was not costructed properly. Expected: " + expected + ", but got: " + command)
	}
}

func TestConstructListCommandSkeletonLogged(t *testing.T) {
	LoggedUser = "test"
	expected := "list|-|"
	command, _ := ConstructListCommand()
	if command != expected {
		t.Errorf("Command skeleton was not costructed properly. Expected: " + expected + ", but got: " + command)
	}
}

func TestConstructListCommandSkeletonNotLogged(t *testing.T) {
	LoggedUser = ""
	expected := "You are not logged in"
	command, _ := ConstructListCommand()
	if command != expected {
		t.Errorf("Command skeleton was not costructed properly. Expected: " + expected + ", but got: " + command)
	}
}

func TestConstructFindCommandSkeletonLogged(t *testing.T) {
	LoggedUser = "test"
	expected := "find|-||-|"
	command, _ := ConstructFindCommand()
	if command != expected {
		t.Errorf("Command skeleton was not costructed properly. Expected: " + expected + ", but got: " + command)
	}
}

func TestConstructFindCommandSkeletonNotLogged(t *testing.T) {
	LoggedUser = ""
	expected := "You are not logged in"
	command, _ := ConstructFindCommand()
	if command != expected {
		t.Errorf("Command skeleton was not costructed properly. Expected: " + expected + ", but got: " + command)
	}
}

func TestConstructCommentCommandSkeletonLogged(t *testing.T) {
	LoggedUser = "test"
	expected := "comment|-||-||-||-|test"
	command, _ := ConstructCommentCommand()
	if command != expected {
		t.Errorf("Command skeleton was not costructed properly. Expected: " + expected + ", but got: " + command)
	}
}

func TestConstructCommentCommandSkeletonNotLogged(t *testing.T) {
	LoggedUser = ""
	expected := "You are not logged in"
	command, _ := ConstructCommentCommand()
	if command != expected {
		t.Errorf("Command skeleton was not costructed properly. Expected: " + expected + ", but got: " + command)
	}
}
