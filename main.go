package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maxiaoyu123-coder/blog-service/global"
	"github.com/maxiaoyu123-coder/blog-service/internal/model"
	"github.com/maxiaoyu123-coder/blog-service/internal/routers"
	"github.com/maxiaoyu123-coder/blog-service/pkg/setting"
	"log"
	"net/http"
	"time"
)

func init() {
	if err := setupSetting(); err != nil {
		log.Fatalf("init.setupSetting error: %v\n", err)
	}
	if err := setupDBEngine(); err != nil {
		log.Fatalf("init.setupDBEngine error: %v\n", err)
	}
}

func main() {
	gin.SetMode(global.ServerSettings.RunMode)
	router := routers.NewRouter()
	server := &http.Server{
		Addr:           ":" + global.ServerSettings.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSettings.ReadTimeout * time.Second,
		WriteTimeout:   global.ServerSettings.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

func setupSetting() error {
	settings, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = settings.ReadSection("Server", &global.ServerSettings)
	if err != nil {
		return err
	}
	err = settings.ReadSection("App", &global.AppSettings)
	if err != nil {
		return err
	}
	settings.ReadSection("Database", &global.DataSettings)
	if err != nil {
		return err
	}
	global.ServerSettings.ReadTimeout *= time.Second //
	global.ServerSettings.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DataSettings)
	if err != nil {
		return err
	}
	return nil
}
