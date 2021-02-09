package envd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// prepareEnv sets needed values in env.
func prepareEnv() {
	return
}

// restoreEnv sets values before tests.
func restoreEnv() {
	return
}

func TestNewSD(t *testing.T) {
	prepareEnv()

	esd, err := NewSD()

	assert.NoError(t, err)
	assert.NotNil(t, esd)
}
