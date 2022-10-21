package components

import "github.com/charmbracelet/lipgloss"

var (
	activeTabBorder      = lipgloss.RoundedBorder()
	inactiveColor        = lipgloss.AdaptiveColor{Light: "#fff", Dark: "#000"}
	activeColor          = lipgloss.AdaptiveColor{Light: "#000", Dark: "#fff"}
	highlightBorderColor = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle     = lipgloss.NewStyle().
				Foreground(inactiveColor).
				Border(lipgloss.HiddenBorder()).
				BorderForeground(highlightBorderColor).
				Padding(0, 1)
	activeTabStyle = inactiveTabStyle.Copy().
			Border(activeTabBorder, true).
			Foreground(activeColor)
)

func RenderTabs(tabs []string, activeTab int) string {
	var renderedTabs []string

	for i, t := range tabs {
		var style lipgloss.Style
		isActive := i == activeTab
		if isActive {
			style = activeTabStyle.Copy()
		} else {
			style = inactiveTabStyle.Copy()
		}
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
}
