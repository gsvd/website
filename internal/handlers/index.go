package handlers

import (
	"net/http"

	"github.com/gsvd/website/internal/tmplengine"
)

func (vh *ViewHandler) ShowIndex(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title":       "Gsvd - Software Developer",
		"Articles":    vh.TemplateEngine.Articles,
		"CurrentPath": r.URL.Path,
	}

	if err := vh.TemplateEngine.RenderView(w, tmplengine.BaseLayout, tmplengine.IndexView, data); err != nil {
		vh.Logger.Error("failed to render index view", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
