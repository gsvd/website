package handlers

import (
	"net/http"

	"github.com/Gsvd/website/internal/template_engine/layout"
	"github.com/Gsvd/website/internal/template_engine/view"
)

func (vh *ViewHandler) ShowContact(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title": "Contact - Gsvd",
	}

	if err := vh.TemplateEngine.RenderView(w, layout.Base, view.Contact, data); err != nil {
		vh.Logger.Error("failed to render contact view", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
