package style

import "charm.land/lipgloss/v2"

var (
	DividerColor   = lipgloss.Color("238")
	TabStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Padding(0, 1)
	ActiveTabStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("108")).Padding(0, 1)

	Footer = lipgloss.NewStyle().Foreground(lipgloss.Color("108"))

	Orange = lipgloss.NewStyle().Foreground(lipgloss.Color("214"))

	KeyStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("135")).Bold(true)
	DescStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	DividerStyle = lipgloss.NewStyle().Foreground(DividerColor)
	StatusStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("108"))
	MatchStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("183")).Bold(true)

	SelectedRowStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("230")).Background(lipgloss.Color("62")).Bold(true)

	LeftStyle   = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderRight(true).BorderForeground(DividerColor).Padding(0, 1)
	MiddleStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderRight(true).BorderForeground(DividerColor).Padding(0, 1)
	RightStyle  = lipgloss.NewStyle().Padding(0, 1)
)
