package renderer

import (
	"encoding/json"
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"gopkg.in/yaml.v3"
)

type Renderer struct {
	format Format
}

type Format string

const (
	FormatTable Format = "table"
	FormatJSON  Format = "json"
	FormatYAML  Format = "yaml"
)

type Formattable interface {
	TableHeader() table.Row
	TableRows() []table.Row
}

func New(f Format) *Renderer {
	return &Renderer{format: f}
}

func (r *Renderer) Render(f Formattable) (string, error) {
	switch r.format {
	case FormatTable:
		return r.renderTable(f)
	case FormatJSON:
		return r.renderJSON(f)
	case FormatYAML:
		return r.renderYAML(f)
	default:
		return "", fmt.Errorf("unsupported format: %s", r.format)
	}
}

func (r *Renderer) renderTable(f Formattable) (string, error) {
	t := table.NewWriter()
	t.AppendHeader(f.TableHeader())
	t.AppendRows(f.TableRows())
	return t.Render(), nil
}

func (r *Renderer) renderJSON(f Formattable) (string, error) {
	j, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func (r *Renderer) renderYAML(f Formattable) (string, error) {
	y, err := yaml.Marshal(f)
	if err != nil {
		return "", err
	}
	return string(y), nil
}
