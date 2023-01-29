package util

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/jedib0t/go-pretty/v6/text"
)

func Border(str string, style lipgloss.Style) string {
	// NOTE: Don't use `lipgloss.Border“.
	//	See https://github.com/charmbracelet/lipgloss/issues/40 .
	lines := strings.Split(str, "\n")
	width := text.LongestLineLen(str)

	b := strings.Repeat("─", width+2)
	bt := style.Render(fmt.Sprintf("┌%s┐", b))
	bb := style.Render(fmt.Sprintf("└%s┘", b))

	rslt := []string{bt}
	for _, line := range lines {
		b := style.Render("│")
		rslt = append(rslt, fmt.Sprintf("%s %s %s", b, text.Pad(line, width, ' '), b))
	}
	rslt = append(rslt, bb)

	return strings.Join(rslt, "\n")
}
