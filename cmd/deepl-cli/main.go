package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/pkg/auth"
	"github.com/leschuster/deepl-cli/ui"
)

const (
	appId = "com.leschuster.deepl-cli"
	user  = "deepl api key"
)

func main() {
	auth := auth.New(appId, user)

	// Redirect logs to a file
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	ui.Run(auth)
}
