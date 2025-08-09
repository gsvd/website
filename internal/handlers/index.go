package handlers

import (
	"net/http"

	"github.com/Gsvd/website/internal/template_engine/layout"
	"github.com/Gsvd/website/internal/template_engine/view"
)

func (vh *ViewHandler) ShowIndex(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title":    "Gsvd - Software Engineer",
		"Articles": vh.TemplateEngine.Articles,
	}

	if err := vh.TemplateEngine.RenderView(w, layout.Base, view.Index, data); err != nil {
		vh.Logger.Error("failed to render index view", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
