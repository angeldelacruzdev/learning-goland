package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {

	name := "Gladys"

	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Gladys")

	if !want.MatchString(msg) || err != nil {

	}
}
