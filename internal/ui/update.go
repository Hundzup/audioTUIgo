package ui

import (
	tea "charm.land/bubbletea/v2"
)


func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type){
	case tea.KeyPressMsg:		
		switch msg.String(){

		case "ctrl + c", "q":
			return m, tea.Quit
			

		case "up", "k":
			if m.cursor > 0{
				m.cursor --
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1{
				m.cursor ++
			}

		case "enter", "space":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else{
				m.selected[m.cursor] = struct{}{}
			}
		}

		
	}
	return m, nil
}
