package components

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFooterShouldContainVersion(t *testing.T) {
	footer := Footer(100)
	regex := regexp.MustCompile(`goodmorning v\d{1}.\d{1}.\d{1}`)
	assert.Regexp(t, regex, footer)
}

func TestFooterShouldHaveMinWidth(t *testing.T) {
	given_len := 100
	footer := Footer(given_len)
	assert.GreaterOrEqual(t, len(footer), given_len)
}
