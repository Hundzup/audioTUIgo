package ui

import (
	"fmt"
  tea "charm.land/bubbletea/v2"
 
)


func (m Model) View() tea.View {
	s := "Expample from doc\n"

	for i, choice := range m.choices{
		cursor := " " // no cursor
		if m.cursor == i{
			cursor = ">"
		}

		checked := " " // no selected
		if _, ok := m.selected[i]; ok{
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"
	v := tea.NewView(s)
	v.AltScreen = true
	return v
}

