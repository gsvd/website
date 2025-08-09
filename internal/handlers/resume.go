package handlers

import (
	"net/http"

	"github.com/gsvd/website/internal/tmplengine"
)

func (vh *ViewHandler) ShowResume(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title":       "Resume - Gsvd",
		"CurrentPath": r.URL.Path,
	}

	if err := vh.TemplateEngine.RenderView(w, tmplengine.BaseLayout, tmplengine.ResumeView, data); err != nil {
		vh.Logger.Error("failed to render resume view", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
