package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/ubselcuk/redkit/internal/screens"
)

func main() {
	p := tea.NewProgram(screens.Session{})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Err: %v", err)
		os.Exit(1)
	}
}
