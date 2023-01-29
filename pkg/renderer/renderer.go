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
	FormatTable    Format = "table"
	FormatCSV      Format = "csv"
	FormatHTML     Format = "html"
	FormatMarkdown Format = "markdown"
	FormatJSON     Format = "json"
	FormatYAML     Format = "yaml"
)

var Formats = []string{
	string(FormatTable),
	string(FormatCSV),
	string(FormatHTML),
	string(FormatMarkdown),
	string(FormatJSON),
	string(FormatYAML),
}

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
	case FormatCSV:
		return r.renderCSV(f)
	case FormatHTML:
		return r.renderHTML(f)
	case FormatMarkdown:
		return r.renderMarkdown(f)
	case FormatJSON:
		return r.renderJSON(f)
	case FormatYAML:
		return r.renderYAML(f)
	default:
		return "", fmt.Errorf("unsupported format: %s", r.format)
	}
}

func (r *Renderer) newTable(f Formattable) table.Writer {
	t := table.NewWriter()
	t.AppendHeader(f.TableHeader())
	t.AppendRows(f.TableRows())
	return t
}

func (r *Renderer) renderTable(f Formattable) (string, error) {
	return r.newTable(f).Render(), nil
}

func (r *Renderer) renderCSV(f Formattable) (string, error) {
	return r.newTable(f).RenderCSV(), nil
}

func (r *Renderer) renderHTML(f Formattable) (string, error) {
	return r.newTable(f).RenderHTML(), nil
}

func (r *Renderer) renderMarkdown(f Formattable) (string, error) {
	return r.newTable(f).RenderMarkdown(), nil
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
