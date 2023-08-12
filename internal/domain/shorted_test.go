package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewShorted(t *testing.T) {
	url := "http://www.site.com"

	t.Run("should open a valid url", func(t *testing.T) {
		shorted, err := NewShorted(url)

		assert.Nil(t, err)
		assert.NotNil(t, shorted)
		assert.Equal(t, url, shorted.URL)

	})

	t.Run("should return error if url is empty", func(t *testing.T) {
		shorted, err := NewShorted("")

		assert.NotNil(t, shorted)
		assert.NotNil(t, err)
		assert.Equal(t, ErrUrlIsInvalid, err)
	})

	t.Run("should open a invalid url", func(t *testing.T) {
		shorted, err := NewShorted("www.site.com")

		assert.Empty(t, shorted)
		assert.NotNil(t, err)
		assert.Equal(t, ErrUrlIsInvalid, err)
	})
}

func TestGenerateHash(t *testing.T) {
	t.Run("hash size", func(t *testing.T) {
		hash := GenerateHash()
		assert.Equal(t, 6, len(hash))
	})
}
