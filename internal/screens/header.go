package screens

import (
	"strings"

	"github.com/ubselcuk/redkit/internal/style"
)

func header(m Session) string {

	tabs := []string{"Live", "Monitor", "CMD", "Settings"}

	header := make([]string, 0, len(tabs))
	for i, name := range tabs {
		if m.page == i {
			header = append(header, style.ActiveTabStyle.Render(name))
			continue
		}
		header = append(header, style.TabStyle.Render(name))
	}

	divider := style.DividerStyle.Render(strings.Repeat("─", m.width))

	return strings.Join(header, "") + "\n" + divider
}
