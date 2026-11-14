package ui

import (
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	//super base + header
	tabs       []string
	tabContent []string
	styles     *styles
	activeTab  int

	//body active tab: need to seletc some thing (like waves)
	firstWindow  []string
	secondWindow []string
	thirdWindow  []string

	//fuck
	choices  []string
	cursor   int
	selected map[int]struct{}

	//some consts
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
