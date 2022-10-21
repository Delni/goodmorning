package settings

import (
	"goodmorning/cmd/settings/components"
	"goodmorning/cmd/settings/pages"
	"goodmorning/cmd/utils"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	tabs      []string
	activeTab int
	height    int
	width     int
}

func Program() error {
	tabs := []string{"Credits"}
	m := model{
		tabs: tabs,
	}
	program := tea.NewProgram(m, tea.WithAltScreen())
	return program.Start()
}

var (
	docStyle    = lipgloss.NewStyle()
	windowStyle = lipgloss.NewStyle().Padding(1, 0)
)

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, utils.Keys.Left):
			m.activeTab--
		case key.Matches(msg, utils.Keys.Right, utils.Keys.Tab):
			m.activeTab++
		case key.Matches(msg, utils.Keys.Quit):
			return m, tea.Quit
		}
	}

	if m.activeTab >= len(m.tabs) {
		m.activeTab = 0
	}

	if m.activeTab < 0 {
		m.activeTab = len(m.tabs) - 1
	}

	return m, nil
}

func (m model) View() string {
	doc := strings.Builder{}

	doc.WriteString(components.RenderTabs(m.tabs, m.activeTab))

	var content string
	switch m.tabs[m.activeTab] {
	case "Credits":
		content = pages.Credits()
	}

	doc.WriteString(windowStyle.Render(content))

	height := m.height - strings.Count(doc.String(), "\n") - 1
	if height > 0 {
		doc.WriteString(strings.Repeat("\n", height))
	}
	doc.WriteString(components.Footer(m.width))
	return docStyle.Render(doc.String())
}
