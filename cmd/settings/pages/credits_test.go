package pages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCredits(t *testing.T) {
	result := Credits()
	assert.IsType(t, "string", result)
}
