package faker

import (
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
)

func TestFakeUserName(t *testing.T) {
	name := User().Name()
	namePrefix := strings.Split(name, "-")[0]
	assert.Contains(t, names, namePrefix, namePrefix)
}

func TestFakeUserFirstName(t *testing.T) {
	assert.Contains(t, firstNames, User().FirstName())
}

func TestFakeUserPath(t *testing.T) {
	assert.Contains(t, paths, User().Path())
}
