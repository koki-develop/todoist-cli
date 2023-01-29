package renderer

import "fmt"

type Renderer struct {
	format Format
}

type Format string

const (
	FormatTable Format = "table"
	FormatJSON  Format = "json"
)

type Formattable interface {
	Table() (string, error)
	JSON() (string, error)
}

func New(f Format) *Renderer {
	return &Renderer{format: f}
}

func (r *Renderer) Render(f Formattable) (string, error) {
	switch r.format {
	case FormatTable:
		return f.Table()
	case FormatJSON:
		return f.JSON()
	default:
		return "", fmt.Errorf("unsupported format: %s", r.format)
	}
}
