package settings

import (
	"goodmorning/cmd/utils"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	tabs      []string
	activeTab int
	height int
}

func Program() error {
	tabs := []string{"Menu", "Apps", "Credits"}
	m := model{
		tabs: tabs,
	}
	program := tea.NewProgram(m, tea.WithAltScreen())
	return program.Start()
}

var (
	activeTabBorder      = lipgloss.RoundedBorder()
	docStyle             = lipgloss.NewStyle()
	inactiveColor        = lipgloss.AdaptiveColor{Light: "#fff", Dark: "#000"}
	activeColor          = lipgloss.AdaptiveColor{Light: "#000", Dark: "#fff"}
	highlightBorderColor = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle     = lipgloss.NewStyle().
				Foreground(inactiveColor).
				Border(lipgloss.HiddenBorder()).
				BorderForeground(highlightBorderColor).
				Padding(0, 1)
	activeTabStyle = inactiveTabStyle.Copy().Border(activeTabBorder, true).Foreground(activeColor)
	windowStyle    = lipgloss.NewStyle().Padding(1, 0)
)

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
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
		m.activeTab = len(m.tabs) -1
	}

	return m, nil
}

func (m model) View() string {
	doc := strings.Builder{}

	var renderedTabs []string

	for i, t := range m.tabs {
		var style lipgloss.Style
		_, _, isActive := i == 0, i == len(m.tabs)-1, i == m.activeTab
		if isActive {
			style = activeTabStyle.Copy()
		} else {
			style = inactiveTabStyle.Copy()
		}
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)

	var content string;
	switch m.tabs[m.activeTab] {
	case "Apps":
		content = "Apps"
	case "Credits":
		content = Credits()
	}
	doc.WriteString(windowStyle.Render(content))
	
	return docStyle.Render(doc.String())
}
