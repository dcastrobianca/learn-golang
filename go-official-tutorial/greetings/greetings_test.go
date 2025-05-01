package greetings

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	//given
	assert := assert.New(t)
	name := "Gladys"
	//when
	msg, err := Hello("Gladys")
	//then
	assert.Empty(err)
	assert.Regexp(regexp.MustCompile(`\b`+name+`\b`), msg)
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	//when
	msg, err := Hello("")
	//then
	assert.Empty(t, msg)
	assert.EqualError(t, err, "empty name")
}
