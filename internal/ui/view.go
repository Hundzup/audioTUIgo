package ui

import (
	// "fmt"
	"strings"
  tea "charm.land/bubbletea/v2"
 	"charm.land/lipgloss/v2"
)


func (m Model) View() tea.View {
	if m.styles == nil {
		return tea.NewView("")
	}

	doc := strings.Builder{}
	s := m.styles

	var renderedTabs []string

	for i, t := range m.tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(m.tabs)-1, i == m.activeTab
		if isActive {
			style = s.activeTab
		} else {
			style = s.inactiveTab
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	doc.WriteString("\n")
	doc.WriteString(s.window.Width(m.width).Render(m.tabContent[m.activeTab]))
	
	v := tea.NewView(s.doc.Render(doc.String()))
	v.AltScreen = true
	
	return v
}

