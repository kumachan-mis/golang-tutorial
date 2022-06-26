package greetings_test

import (
	"golang-tutorial/greetings"
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	want := regexp.MustCompile(`\bGladys\b`)
	message, err := greetings.Hello("Gladys")
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Gladys") is expected to be %#q, nil, but got %q, %v`, want, message, err)
	}
}

func TestEmpty(t *testing.T) {
	message, err := greetings.Hello("")
	if message != "" || err == nil {
		t.Fatalf(`Hello("Gladys") is expected to be "", error, but got %q, %v`, message, err)
	}
}
