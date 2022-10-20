package settings

import "github.com/charmbracelet/lipgloss"

var (
	loveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	urlStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("35"))
)

func Credits() string {
	return `
Open Source license (MIT)
Coded with ` + loveStyle.Render("♥️") + ` in Golang. 

Copyright (c) 2022 Nicolas Delauney
More info on ` + urlStyle.Render("https://github.com/delni/goodmorning") + "."

}
