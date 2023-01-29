package styles

import "github.com/charmbracelet/lipgloss"

var (
	// colors
	ColorMain = lipgloss.Color("#E44332")

	// styles
	StyleNotificationBorder = lipgloss.NewStyle().Foreground(ColorMain)
	StyleNotificationText   = lipgloss.NewStyle().Bold(true)
	StyleLink               = lipgloss.NewStyle().Underline(true)
)
