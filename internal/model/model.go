package model

import (
	"fmt"
	"github.com/guidoxie/myblog/global"
	"github.com/guidoxie/myblog/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreateOn   uint32 `json:"create_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	/*
		parseTime 是查询结果是否自动解析为时间。
		loc 是MySQL的时区设置。
	*/
	db, err := gorm.Open(databaseSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			databaseSetting.UserName,
			databaseSetting.Password,
			databaseSetting.Host,
			databaseSetting.DBName,
			databaseSetting.Charset,
			databaseSetting.ParseTime,
		))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	/*
		gorm默认使用复数映射，go代码的单数、复数struct形式都匹配到复数表中：创建表、添加数据时都是如此。
		指定了db.SingularTable(true)之后，进行严格匹配。
	*/
	db.SingularTable(true)

	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
