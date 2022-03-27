package model

import "github.com/guidoxie/myblog/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (b Tag) TableName() string {
	return "blog_tag"
}
