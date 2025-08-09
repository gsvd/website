package handlers

import (
	"log/slog"

	"github.com/Gsvd/website/internal/template_engine"
)

type ViewHandler struct {
	TemplateEngine *template_engine.TemplateEngine
	Logger         *slog.Logger
}

func NewViewHandler(te *template_engine.TemplateEngine, logger *slog.Logger) *ViewHandler {
	return &ViewHandler{TemplateEngine: te, Logger: logger}
}
