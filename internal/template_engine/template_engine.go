package template_engine

import (
	"fmt"
	"io/fs"
	"net/http"
	"text/template"

	"github.com/Gsvd/website/internal/blog"
	"github.com/Gsvd/website/internal/template_engine/layout"
	"github.com/Gsvd/website/internal/template_engine/view"
	"github.com/Gsvd/website/web"
	"github.com/pkg/errors"
)

type (
	TemplateEngine struct {
		Articles    []*blog.Article
		ArticlesMap map[string]*blog.Article
	}
	Option func(*TemplateEngine) error
)

func New() (*TemplateEngine, error) {
	// Articles are loaded at initialization.
	articles, articlesMap, err := blog.GetArticles()
	if err != nil {
		return nil, errors.Wrap(err, "blog.GetArticles")
	}

	return &TemplateEngine{
		Articles:    articles,
		ArticlesMap: articlesMap,
	}, nil
}

func (te *TemplateEngine) RenderView(w http.ResponseWriter, layout layout.Layout, view view.View, data any) error {
	layoutFile := fmt.Sprintf("layouts/%s.html", layout)
	viewFile := fmt.Sprintf("templates/%s.html", view)

	partialFiles, err := fs.Glob(web.ViewsFS(), "partials/*.html")
	if err != nil {
		return errors.Wrap(err, "fs.Glob")
	}

	files := append([]string{layoutFile, viewFile}, partialFiles...)

	tmpl, err := template.ParseFS(web.ViewsFS(), files...)
	if err != nil {
		return errors.Wrap(err, "template.ParseFS")
	}

	if err := tmpl.ExecuteTemplate(w, string(layout), data); err != nil {
		return errors.Wrap(err, "tmpl.ExecuteTemplate")
	}

	return nil
}
