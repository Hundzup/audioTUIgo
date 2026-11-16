package ui

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {

		case "ctrl + c", "q", "й":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", "space":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		
		case "f1":
			m.activeTab = 0
			return m, nil

		case "f2":
			m.activeTab = 1
			return m, nil

		case "f3":
			m.activeTab = 2
			return m, nil

		case "f4":
			m.activeTab = 3
			return m, nil


		case "right", "l", "n", "tab":
			m.activeLayer = min(m.activeLayer+1, len(m.layers)-1)
			return m, nil
		case "left", "h", "p", "shift+tab":
			m.activeLayer = max(m.activeLayer-1, 0)
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}
