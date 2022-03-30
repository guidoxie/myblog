package service

import (
	"context"
	"github.com/guidoxie/myblog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}
