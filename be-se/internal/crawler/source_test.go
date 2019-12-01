package crawler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJSONSourceLoader(t *testing.T) {
	jsl := NewJSONSourceLoader("../../assets/initial_sources.json")
	srcs, err := jsl.Load()

	assert.Nil(t, err)
	assert.NotNil(t, srcs)
}
