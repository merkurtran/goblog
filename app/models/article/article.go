package article

import (
	"strconv"

	"github.com/merkurtran/goblog/app/models"
	"github.com/merkurtran/goblog/pkg/route"
)

type Article struct {
	models.BaseModel

	Title string
	Body  string
}

func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))
}
