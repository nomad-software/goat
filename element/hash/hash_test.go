package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomHash(t *testing.T) {
	hash := Generate()

	assert.Len(t, hash, 8)
}

func TestDeterministicHash(t *testing.T) {
	hash := Generate("foo", "bar", "baz")

	assert.Equal(t, "606D6255", hash)
}
