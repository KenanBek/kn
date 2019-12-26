package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	h1 := Hash("http://example.com")
	h2 := Hash("http://example.com/a")
	h3 := Hash("http://sub.example.com")

	assert.Equal(t, "f0e6a6", h1)
	assert.Equal(t, "5bd48f", h2)
	assert.Equal(t, "711459", h3)
}
