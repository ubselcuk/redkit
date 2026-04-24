package screens

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (s Session) Init() tea.Cmd {
	return nil
}

func (s Session) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height

	case tea.KeyPressMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return s, tea.Quit
		}
	}

	return s, nil
}

func (s Session) View() tea.View {
	header := header(s)
	footer := footer(s)

	h := s.height - lipgloss.Height(header) - lipgloss.Height(footer)
	if h < 1 {
		h = 1
	}

	mid := "Press q to quit."

	body := lipgloss.NewStyle().Height(h).Render(mid)

	str := header + "\n" + body + "\n" + footer

	view := tea.NewView(str)
	view.AltScreen = true

	return view
}
