package model

import "github.com/guidoxie/myblog/pkg/app"

type Article struct {
	*Model
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (b Article) TableName() string {
	return "blog_article"
}

type TagArticle struct {
	List  []*Article
	Pager *app.Pager
}
