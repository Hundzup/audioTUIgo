package main

import (
	"fmt"
	"os"
  tea "charm.land/bubbletea/v2"
  "audioTUIgo/internal/ui"
)



func main() {

	p := tea.NewProgram(
		ui.InitialModel(),
		// tea.WithWindowSize(),
		// tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil{
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
