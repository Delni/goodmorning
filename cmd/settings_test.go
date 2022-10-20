package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func RunSettingsOnBuffer() (*cobra.Command, *bytes.Buffer){
	command := NewSettingsCmd()
	buffer := bytes.NewBufferString("")
	command.SetOut(buffer)
	return command, buffer
}
func TestSettingsCmd(t *testing.T) {
	command, _ := RunSettingsOnBuffer()
	err := command.Execute()
	assert.Nil(t, err)
}