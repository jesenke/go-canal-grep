package util

import (
	"fmt"
	"golang.org/x/exp/slog"
	"os"
	"sync"
	"time"
)

var Log *slog.Logger

func InitLog() {
	lock := &sync.RWMutex{}
	createLogFile(lock)
}

func createLogFile(lock *sync.RWMutex) {
	lock.Lock()
	defer lock.Unlock()
	path := Config.Get("log.Path")
	level := Config.Get("log.Level")
	logFile := fmt.Sprintf("%s/grep.log.%s.log", path, time.Now().Format("20060102"))
	logout, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic("log.file.err:" + err.Error())
	}
	lvel := slog.LevelDebug
	switch level {
	case "all":
	case "debug":
	case "info":
		lvel = slog.LevelInfo
	case "warn":
		lvel = slog.LevelWarn
	case "error":
		lvel = slog.LevelError
	}
	Log = slog.New(slog.NewJSONHandler(logout, &slog.HandlerOptions{
		AddSource: true,
		Level:     lvel,
	}))
	slog.SetDefault(Log)
	println("logLevel:", lvel.String())
	Log.Warn("log服务开始了")
}
