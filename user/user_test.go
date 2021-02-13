package user

import (
	"testing"

	"bou.ke/monkey"
	"golang.org/x/crypto/bcrypt"
)

func TestCompareMatchingPasswords(t *testing.T) {
	hashedPassword := "$2a$04$1ORhJ3MFDnqPHdIPsUgzOuGNT/Vt8DVt4UhISBg1oYLlxtPTd0WXy"
	plainPassword := "password1234"

	if !ComparePasswords(hashedPassword, plainPassword) {
		t.Errorf("Passwords didn't match")
	}
}

func TestCompareDifferentPasswords(t *testing.T) {
	hashedPassword := "$2a$04$1ORhJ3MFDnqPHdIPsUgzOuGNT/Vt8DVt4UhISBg1oYL"
	plainPassword := "password1234"

	if ComparePasswords(hashedPassword, plainPassword) {
		t.Errorf("Passwords matched")
	}
}

func TestHashAndSalt(t *testing.T) {
	defer monkey.UnpatchAll()
	expected := []byte("asdf")
	plainPassword := "password1234"

	monkey.Patch(bcrypt.GenerateFromPassword, func([]byte, int) ([]byte, error) {
		return expected, nil
	})

	got := HashAndSalt(plainPassword)

	if got != string(expected) {
		t.Errorf("Password was not hashed correctly. Expected: " + string(expected) + ", but got " + got)
	}
}
