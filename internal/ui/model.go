package ui

import (
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	tabs       []string
	tabContent []string
	styles     *styles
	activeTab  int

	choices  []string
	cursor   int
	selected map[int]struct{}

	width  int
	height int
}

func InitialModel() Model {
	tabs := []string{"Library", "Files", "PlayLists", "Recommendation"}
	tabContent := []string{"Scaned music", "Brose files", "Place for created playlists", "Хуй тебе а не рекоммендации"}
	return Model{
		choices:    []string{"Something", "Kall"},
		selected:   make(map[int]struct{}),
		tabs:       tabs,
		tabContent: tabContent,
		styles:     defaultStyles(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
