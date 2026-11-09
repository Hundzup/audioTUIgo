package ui

import (
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	choices []string
  cursor  int
	selected map[int]struct{}
}


func InitialModel() Model{
	return Model{
		choices: []string{"Something", "Kall"},
		selected: make(map[int]struct{}),
	}
}

func (m Model) Init() tea.Cmd{
	return nil
}
