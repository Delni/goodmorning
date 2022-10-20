package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func RunRootOnBuffer() (*cobra.Command, *bytes.Buffer) {
	command := NewRootCmd()
	buffer := bytes.NewBufferString("")
	command.SetOut(buffer)
	return command, buffer
}
func TestRootCmd(t *testing.T) {
	command, _ := RunRootOnBuffer()
	err := command.Execute()
	assert.Nil(t, err)
}

func TestRootCmdShouldHaveSettingsFlag(t *testing.T) {
	command, _ := RunRootOnBuffer()
	command.SetArgs([]string{"settings"})
	err := command.Execute()
	assert.Nil(t, err, err)
}
