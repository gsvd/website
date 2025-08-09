package handlers

import (
	"log/slog"

	"github.com/gsvd/website/internal/tmplengine"
)

type ViewHandler struct {
	TemplateEngine *tmplengine.TmplEngine
	Logger         *slog.Logger
}

func NewViewHandler(te *tmplengine.TmplEngine, logger *slog.Logger) *ViewHandler {
	return &ViewHandler{TemplateEngine: te, Logger: logger}
}
