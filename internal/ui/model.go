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
	layers []string
	activeLayer int

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

	layers := []string{"author", "alboms", "tracks"}
	return Model{
		choices:    []string{"Something", "Kall"},
		selected:   make(map[int]struct{}),
		tabs:       tabs,
		tabContent: tabContent,
		styles:     defaultStyles(),
		layers:     layers,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
