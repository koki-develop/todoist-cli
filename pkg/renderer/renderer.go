package renderer

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/koki-develop/todoist-cli/pkg/util"
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
	Maps() []map[string]interface{}
}

func New(f Format) *Renderer {
	return &Renderer{format: f}
}

func (r *Renderer) Render(f Formattable, cols []string) (string, error) {
	switch r.format {
	case FormatTable:
		return r.renderTable(f, cols)
	case FormatCSV:
		return r.renderCSV(f, cols)
	case FormatHTML:
		return r.renderHTML(f, cols)
	case FormatMarkdown:
		return r.renderMarkdown(f, cols)
	case FormatJSON:
		return r.renderJSON(f)
	case FormatYAML:
		return r.renderYAML(f)
	default:
		return "", fmt.Errorf("unsupported format: %s", r.format)
	}
}

func (r *Renderer) newTable(f Formattable, cols []string) (table.Writer, error) {
	t := table.NewWriter()

	// header
	h := make([]string, len(cols))
	for i, col := range cols {
		h[i] = strings.ToUpper(col)
	}
	t.AppendHeader(util.StringsToInterfaces(h))

	// rows
	rows := []table.Row{}
	for _, m := range f.Maps() {
		row := table.Row{}

		for _, col := range cols {
			row = append(row, r.renderColumn(m[col]))
		}

		rows = append(rows, row)
	}
	t.AppendRows(rows)

	return t, nil
}

func (r *Renderer) renderColumn(v interface{}) string {
	if v == nil {
		return ""
	}

	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(v)
		cols := make([]string, s.Len())
		for i := 0; i < s.Len(); i++ {
			cols[i] = r.renderColumn(s.Index(i))
		}
		return strings.Join(cols, ", ")
	case reflect.Map:
		return fmt.Sprintf("%#v", v)
	default:
		return fmt.Sprint(v)
	}
}

func (r *Renderer) renderTable(f Formattable, cols []string) (string, error) {
	t, err := r.newTable(f, cols)
	if err != nil {
		return "", err
	}
	return t.Render(), nil
}

func (r *Renderer) renderCSV(f Formattable, cols []string) (string, error) {
	t, err := r.newTable(f, cols)
	if err != nil {
		return "", err
	}
	return t.RenderCSV(), nil
}

func (r *Renderer) renderHTML(f Formattable, cols []string) (string, error) {
	t, err := r.newTable(f, cols)
	if err != nil {
		return "", err
	}
	return t.RenderHTML(), nil
}

func (r *Renderer) renderMarkdown(f Formattable, cols []string) (string, error) {
	t, err := r.newTable(f, cols)
	if err != nil {
		return "", err
	}
	return t.RenderMarkdown(), nil
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
