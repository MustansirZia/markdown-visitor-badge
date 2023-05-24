package svgRenderer

import (
	"io"
	"sync"
	"text/template"
)

// BadgeParams - Parameters for rendering the badge.
type BadgeParams struct {
	// Label - Label for the visitors.
	Label string
	// LabelBgColor - Background color for the visitors label area.
	LabelBgColor string
	// LabelColor - Color for the visitors label.
	LabelColor string
	// Count - Visitor count.
	Count uint64
	// CountBgColor - Background color for the count area.
	CountBgColor string
	// CountColor - Color for the count.
	CountColor string
}

// SvgRenderer - Interface for rendering the badge as an SVG.
type SvgRenderer interface {
	// RenderBadge - Writes contents of the badge into the writer.
	RenderBadge(w io.Writer, params BadgeParams) error
}

// NewSvgRenderer - Constructs and returns a new SVG Renderer.
func NewSvgRenderer() SvgRenderer {
	return &templateSvgRenderer{}
}

type templateSvgRenderer struct {
	svgTemplate *template.Template
	once        sync.Once
}

func (t *templateSvgRenderer) RenderBadge(w io.Writer, params BadgeParams) error {
	t.once.Do(func() {
		t.svgTemplate = template.Must(template.New("templateSvg").Parse(`
			<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="138" height="28" role="img" aria-label="visitors: 21">
				<title>visitors: {{.Count}}</title>
				<g shape-rendering="crispEdges">
					<rect width="77" height="28" fill="{{.LabelBgColor}}" />
					<rect x="77" width="61" height="28" fill="{{.CountBgColor}}" />
				</g>
				<g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="100">
					<text fill="{{.LabelColor}}" x="385" y="175" transform="scale(.1)" textLength="530">{{.Label}}</text>
					<text fill="{{.CountColor}}" x="1050" y="175" transform="scale(.1)" textLength="370" font-weight="bold">{{.Count | printf "%04d"}}</text>
				</g>
			</svg>
		`))
	})
	if err := t.svgTemplate.Execute(w, params); err != nil {
		return err
	}
	return nil
}
