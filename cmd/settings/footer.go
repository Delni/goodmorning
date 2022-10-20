package settings

import "github.com/charmbracelet/lipgloss"

var (
	footerStyle = lipgloss.NewStyle().
			Padding(0, 0, 0, 1).
			Background(lipgloss.AdaptiveColor{Light: "#ddd", Dark: "#333"})
	shortcutStyle = footerStyle.Copy().Foreground(lipgloss.AdaptiveColor{Light: "#666", Dark: "#aaa"})
	actionStyle   = footerStyle.Copy().Foreground(lipgloss.AdaptiveColor{Light: "#aaa", Dark: "#666"})
	versionStyle  = lipgloss.NewStyle().
			Padding(0, 1).
			Background(lipgloss.Color("32")).
			Foreground(lipgloss.Color("251"))
)

func Footer(width int) string {
	version := versionStyle.Render("goodmorning v0.0.1")
	shortcuts := shortcutStyle.Render("q/esc")
	action := actionStyle.Render("quit")
	return lipgloss.JoinHorizontal(lipgloss.Bottom,
		shortcuts,
		action,
		footerStyle.Width(width-lipgloss.Width(version+shortcuts+action)).Render(""),
		version,
	)
}
