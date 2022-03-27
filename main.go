package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guidoxie/myblog/global"
	"github.com/guidoxie/myblog/internal/model"
	"github.com/guidoxie/myblog/internal/routers"
	"github.com/guidoxie/myblog/pkg/logger"
	"github.com/guidoxie/myblog/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"path"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	if err = setupLogger(); err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	if err = setupDBEngine(); err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 编程之旅： 一起用 Go 做项目
// @termsOfService https://github.com/guidoxie/myblog
func main() {
	gin.SetMode(global.ServerSetting.RunMode)

	global.Logger.Infof("xxxx")
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	if err = s.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	if err = s.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}
	if err = s.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  path.Join(global.AppSetting.LogSavePath, global.AppSetting.LogFileName+global.AppSetting.LogFileExt),
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	return err
}
