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

		case "right", "l", "n", "tab":
					m.activeTab = min(m.activeTab+1, len(m.tabs)-1)
					return m, nil
		case "left", "h", "p", "shift+tab":
			m.activeTab = max(m.activeTab-1, 0)
			return m, nil
		}
		
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}
