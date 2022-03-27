package model

import "github.com/guidoxie/myblog/pkg/app"

type Article struct {
	*Model
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (b Article) TableName() string {
	return "blog_article"
}
