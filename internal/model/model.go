package model

import (
	"fmt"
	"github.com/guidoxie/myblog/global"
	"github.com/guidoxie/myblog/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
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
	// 注册回调
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

// 新增行为的回调
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		// 当前是否包含所需的字段
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			// 该字段是否为空
			if createTimeField.IsBlank {
				// 设置字段的值
				_ = createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

// 更新行为的回调
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	// 获取当前设置的标识gorm:update_column的字段属性
	if _, ok := scope.Get("gorm:update_column"); !ok {
		// 在更新回调内设置ModifiedOn的值为当前的时间戳
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// 删除行为的回调
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		// 获取标识 gorm:delete_option的字段属性
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		// 判断是否存在DeletedOn和IsDel，否则执行delete操作进行硬删除
		deleteOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")

		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(), // 获取表名
				scope.Quote(deleteOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			),
			).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			),
			).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
