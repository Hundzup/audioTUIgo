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
		isActive := i == m.activeTab
		if isActive {
			style = s.activeTab
		} else {
			style = s.inactiveTab
		}
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	header := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.width - 1).
		Height(int(float64(m.height)*0.05)).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...))
	body := s.body.
		Width(m.width - 1).
		// Height(int(float64(m.height)*0.80)).
		Height(m.height - 2*lipgloss.Height(header)).
		Render(m.tabContent[m.activeTab])
	footer := s.footer.
		Height(int(float64(m.height)*0.05)).
		Width(m.width - 1).
		Render("footer")

	doc.WriteString(header)
	doc.WriteString("\n")
	doc.WriteString(body)
	doc.WriteString("\n")
	doc.WriteString(footer)

	v := tea.NewView(s.doc.Render(doc.String()))
	v.AltScreen = true

	return v
}
