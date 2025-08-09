package tmplengine

import (
	"fmt"
	"io/fs"
	"net/http"
	"strings"
	"text/template"

	"github.com/gsvd/website/internal/blog"
	"github.com/gsvd/website/web"
	"github.com/pkg/errors"
)

type (
	TmplEngine struct {
		Articles    []*blog.Article
		ArticlesMap map[string]*blog.Article
		CommitHash  string
	}
	Option func(*TmplEngine) error
	Layout string
	View   string
)

const (
	BaseLayout    Layout = "base"
	ArticleLayout Layout = "article"
	IndexView     View   = "index"
	ContactView   View   = "contact"
	ResumeView    View   = "resume"
	ArticleView   View   = "article"
	BlogView      View   = "blog"
)

func New(commitHash string) (*TmplEngine, error) {
	// Articles are loaded at initialization.
	articles, articlesMap, err := blog.GetArticles()
	if err != nil {
		return nil, errors.Wrap(err, "blog.GetArticles")
	}

	return &TmplEngine{
		Articles:    articles,
		ArticlesMap: articlesMap,
		CommitHash:  commitHash,
	}, nil
}

func (te *TmplEngine) RenderView(w http.ResponseWriter, layout Layout, view View, data any) error {
	layoutFile := fmt.Sprintf("layouts/%s.html", layout)
	viewFile := fmt.Sprintf("templates/%s.html", view)

	partialFiles, err := fs.Glob(web.ViewsFS(), "partials/*.html")
	if err != nil {
		return errors.Wrap(err, "fs.Glob")
	}

	files := append([]string{layoutFile, viewFile}, partialFiles...)

	funcMap := template.FuncMap{
		"hasPrefix": func(s, prefix string) bool {
			return strings.HasPrefix(s, prefix)
		},
		"commitHash": func() string {
			return te.CommitHash
		},
	}

	tmpl, err := template.New("").Funcs(funcMap).ParseFS(web.ViewsFS(), files...)
	if err != nil {
		return errors.Wrap(err, "template.ParseFS")
	}

	if err := tmpl.ExecuteTemplate(w, string(layout), data); err != nil {
		return errors.Wrap(err, "tmpl.ExecuteTemplate")
	}

	return nil
}
