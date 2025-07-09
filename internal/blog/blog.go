package blog

import (
	"fmt"
	"html/template"
	"io/fs"
	"sort"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/pkg/errors"
	"github.com/russross/blackfriday/v2"
	"gopkg.in/yaml.v2"
)

type Article struct {
	Id          int       `yaml:"id"`
	Slug        string    `yaml:"slug"`
	Title       string    `yaml:"title"`
	Description string    `yaml:"description"`
	Date        time.Time `yaml:"date"`
	Content     template.HTML
}

func GetArticles() (articles []*Article, articlesMap map[string]*Article, err error) {
	articleFiles, err := fs.ReadDir(FS, "articles")
	if err != nil {
		return nil, nil, errors.Wrap(err, "fs.ReadDir")
	}

	articles = make([]*Article, 0, len(articleFiles))
	articlesMap = make(map[string]*Article, len(articleFiles))

	for _, file := range articleFiles {
		if file.IsDir() {
			continue
		}

		var article *Article

		filePath := fmt.Sprintf("articles/%s", file.Name())
		content, err := FS.ReadFile(filePath)
		if err != nil {
			return nil, nil, errors.Wrap(err, "FS.ReadFile")
		}

		content, err = frontmatter.Parse(strings.NewReader(string(content)), &article, frontmatter.NewFormat("---", "---", yaml.Unmarshal))
		if err != nil {
			return nil, nil, errors.Wrap(err, "frontmatter.Parse")
		}

		output := blackfriday.Run(content)
		article.Content = template.HTML(output)

		articles = append(articles, article)
		articlesMap[article.Slug] = article
	}

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Date.After(articles[j].Date)
	})

	return articles, articlesMap, nil
}

func (a *Article) FormattedDate() string {
	return a.Date.Format("January 2, 2006")
}
