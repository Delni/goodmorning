package storage

import "strings"

type App struct {
	label   string
	command string
	active  bool
}

func (a App) slug() string {
	slug := strings.ToLower(a.label)
	slug = strings.Trim(slug, " ")
	slug = strings.ReplaceAll(slug, " ", "_")
	return slug
}
