package screens

import (
	"strings"

	"github.com/ubselcuk/redkit/internal/style"
)

func footer(m Session) string {
	divider := style.DividerStyle.Render(strings.Repeat("─", m.width)) + "\n"
	shortcuts := " "

	shortcuts += style.Footer.Render("q") + style.DescStyle.Render("uit") + style.DividerStyle.Render(" │ ")
	if m.page == 0 || m.page == 1 || m.page == 2 {
		shortcuts += style.Footer.Render("1 2 3") + style.DescStyle.Render(" tab") + style.DividerStyle.Render(" │ ")
	}

	if m.page == 0 {
		shortcuts += style.Footer.Render("↑↓") + style.DescStyle.Render(" navigate") + style.DividerStyle.Render(" │ ")
		shortcuts += style.Footer.Render("s") + style.DescStyle.Render("earch") + style.DividerStyle.Render(" │ ")
	}

	if m.page == 1 {
		shortcuts += style.Footer.Render("↑↓") + style.DescStyle.Render(" navigate") + style.DividerStyle.Render(" │ ")
		shortcuts += style.Footer.Render("←→") + style.DescStyle.Render(" tab") + style.DividerStyle.Render(" │ ")
		shortcuts += style.Footer.Render("enter") + style.DescStyle.Render(" select") + style.DividerStyle.Render(" │ ")
	}

	return divider + shortcuts
}
