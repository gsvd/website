package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gsvd/website/internal/tmplengine"
)

func (vh *ViewHandler) ShowBlog(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title":       "Blog Articles - Gsvd",
		"Articles":    vh.TemplateEngine.Articles,
		"CurrentPath": r.URL.Path,
	}

	if err := vh.TemplateEngine.RenderView(w, tmplengine.BaseLayout, tmplengine.BlogView, data); err != nil {
		vh.Logger.Error("failed to render blog view", slog.Any("error", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (vh *ViewHandler) ShowArticle(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	article, ok := vh.TemplateEngine.ArticlesMap[slug]
	if !ok {
		http.NotFound(w, r)
		return
	}

	data := map[string]any{
		"Title":       fmt.Sprintf("%s - Gsvd", article.Title),
		"Article":     article,
		"Canonical":   fmt.Sprintf("blog/%s", slug),
		"CurrentPath": r.URL.Path,
	}

	if err := vh.TemplateEngine.RenderView(w, tmplengine.ArticleLayout, tmplengine.ArticleView, data); err != nil {
		vh.Logger.Error("failed to render article view", slog.Any("error", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
