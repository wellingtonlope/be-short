package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashSize(t *testing.T) {
	hash := GenerateHash()

	assert.Equal(t, 6, len(hash), "The number of characters in hash mustn't be less or greater than 6")
}
