package util

import (
	"fmt"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/stretchr/testify/assert"
)

func TestBorder(t *testing.T) {
	style := lipgloss.NewStyle()

	tests := []struct {
		str  string
		want string
	}{
		{
			"Hello World",
			`┌─────────────┐
│ Hello World │
└─────────────┘`,
		},
		{
			"Hello World\nGoodnight World",
			`┌─────────────────┐
│ Hello World     │
│ Goodnight World │
└─────────────────┘`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := Border(tt.str, style)
			assert.Equal(t, tt.want, got)
		})
	}
}
