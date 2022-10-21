package settings

import (
	"goodmorning/cmd/settings/components"
	"goodmorning/cmd/settings/pages"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
)

var testModel = model{tabs: []string{"Test"}}

func TestUpdateWindowSize(t *testing.T) {
	message := tea.WindowSizeMsg{Width: 10, Height: 5}
	result, _ := testModel.Update(message)

	assert.Equal(t, 10, result.(model).width)
	assert.Equal(t, 5, result.(model).height)
}

func TestActiveTabsLeftBoundary(t *testing.T) {
	given := model{tabs: []string{"Test", "Test"}, activeTab: -1}
	result := given.SetActiveTabInBounds()

	assert.Equal(t, 1, result.activeTab)
}

func TestActiveTabsRightBoundary(t *testing.T) {
	given := model{tabs: []string{"Test", "Test"}, activeTab: 3}
	result := given.SetActiveTabInBounds()

	assert.Equal(t, 0, result.activeTab)
}

func TestLeftKeyShouldDecreaseActiveTab(t *testing.T) {
	activeTab := 1
	given := model{tabs: []string{"Test", "Test"}, activeTab: activeTab}
	result, _ := given.Update(tea.KeyMsg{
		Type: tea.KeyLeft,
	})

	assert.Equal(t, activeTab-1, result.(model).activeTab)
}

func TestRightKeyShouldIncreaseActiveTab(t *testing.T) {
	activeTab := 0
	given := model{tabs: []string{"Test", "Test"}, activeTab: activeTab}
	result, _ := given.Update(tea.KeyMsg{
		Type: tea.KeyRight,
	})

	assert.Equal(t, activeTab+1, result.(model).activeTab)
}

func TestTabKeyShouldIncreaseActiveTab(t *testing.T) {
	activeTab := 0
	given := model{tabs: []string{"Test", "Test"}, activeTab: activeTab}
	result, _ := given.Update(tea.KeyMsg{
		Type: tea.KeyTab,
	})

	assert.Equal(t, activeTab+1, result.(model).activeTab)
}

func TestQShouldSendQuitMessage(t *testing.T) {
	_, cmd := testModel.Update(tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'q'},
	})

	assert.NotNil(t, cmd)
}

func TestView(t *testing.T) {
	view := testModel.View()

	assert.Contains(t, view, components.RenderTabs(testModel.tabs, 0))
	assert.Contains(t, view, components.Footer(10))
}

func TestViewShoulDisplayCreditsWhenTabIsCredit(t *testing.T) {
	view := model{tabs: []string{"Credits"}}.View()
	assert.Contains(t,
		strings.ReplaceAll(view, " ", ""),
		strings.ReplaceAll(pages.Credits(), " ", ""),
	)
}
