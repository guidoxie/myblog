package model

import "github.com/guidoxie/myblog/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (b Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
