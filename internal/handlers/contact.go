package handlers

import (
	"net/http"

	"github.com/gsvd/website/internal/tmplengine"
)

func (vh *ViewHandler) ShowContact(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title":       "Contact - Gsvd",
		"CurrentPath": r.URL.Path,
	}

	if err := vh.TemplateEngine.RenderView(w, tmplengine.BaseLayout, tmplengine.ContactView, data); err != nil {
		vh.Logger.Error("failed to render contact view", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
