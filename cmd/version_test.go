package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnapshotVersion(t *testing.T) {
	out, err := executeCommand(swageCmd.Root(), versionCmd.Use)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("swage version %s\n", swageVersion), string(out))
}