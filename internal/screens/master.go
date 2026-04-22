package screens

import tea "charm.land/bubbletea/v2"

func (s Session) Init() tea.Cmd {
	return nil
}

func (s Session) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		s.Width = msg.Width
		s.Height = msg.Height

	case tea.KeyPressMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return s, tea.Quit
		}
	}

	return s, nil
}

func (s Session) View() tea.View {

	str := "Press q to quit."

	view := tea.NewView(str)
	view.AltScreen = true

	return view
}
