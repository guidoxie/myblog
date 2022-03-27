package global

import (
	"github.com/guidoxie/myblog/pkg/logger"
	"github.com/guidoxie/myblog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
