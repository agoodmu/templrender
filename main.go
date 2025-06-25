package templrender

import (
	"context"
	"errors"
	"net/http"

	"github.com/a-h/templ"
)

type TemplRender struct {
	Context   context.Context
	Component templ.Component
}

func (t TemplRender) Render(w http.ResponseWriter) error {
	t.WriteContentType(w)
	if t.Component == nil {
		return errors.New("invalid templ componet")
	}
	return t.Component.Render(t.Context, w)

}

func (t TemplRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func New(ctx context.Context, component templ.Component) *TemplRender {
	return &TemplRender{
		Context:   ctx,
		Component: component,
	}
}
