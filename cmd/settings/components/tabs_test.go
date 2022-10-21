package components

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTabsRendering(t *testing.T) {
	tabs := RenderTabs([]string{"Test1", "Test2"}, 0)
	assert.Contains(t, tabs, "Test1")
	assert.Contains(t, tabs, "Test2")
}

func TestTabsShouldRenderAsThreeLines(t *testing.T) {
	tabs := RenderTabs([]string{"Test1", "Test2"}, 0)
	assert.Equal(t, 2, strings.Count(tabs, "\n"))
}
