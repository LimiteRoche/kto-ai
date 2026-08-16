package styles

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// logoLines contains the ASCII art for the k.to wordmark logo.
var logoLines = []string{
	"██╗  ██╗   ████████╗ ██████╗",
	"██║ ██╔╝   ╚══██╔══╝██╔═══██╗",
	"█████═╝  ██╗  ██║   ██║   ██║",
	"██╔═██╗  ╚═╝  ██║   ██║   ██║",
	"██║  ██╗      ██║   ╚██████╔╝",
	"╚═╝  ╚═╝      ╚═╝    ╚═════╝",
}

// gradientColors defines the top-to-bottom gradient for the logo.
// Distributed across rows: the k.to gradient.
var gradientColors = []lipgloss.Color{
	lipgloss.Color("#ff2b2b"), // band 1: red
	lipgloss.Color("#ff4d1a"), // band 2: red-orange
	lipgloss.Color("#ff6a00"), // band 3: orange
	lipgloss.Color("#ff8a00"), // band 4: orange-amber
	lipgloss.Color("#ffaa00"), // band 5: amber
}

// RenderLogo returns the ASCII logo with a top-to-bottom gradient.
func RenderLogo() string {
	total := len(logoLines)
	if total == 0 {
		return ""
	}

	bands := len(gradientColors)
	var b strings.Builder

	for i, line := range logoLines {
		bandIdx := (i * bands) / total
		if bandIdx >= bands {
			bandIdx = bands - 1
		}
		style := lipgloss.NewStyle().Foreground(gradientColors[bandIdx])
		b.WriteString(style.Render(line))
		if i < total-1 {
			b.WriteByte('\n')
		}
	}

	return b.String()
}
