package main

import (
	"flag"
	"io"
	"log/slog"
	"os"

	"github.com/alhaos/telegramBot/internal/config"
	"github.com/alhaos/telegramBot/internal/telegram"
	"github.com/alhaos/telegramBot/internal/worker"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {

	// Get config path from flag -config
	configFilenamePointer := flag.String("config", "config/config.yml", "app config file")
	flag.Parse()
	configFilename := *configFilenamePointer

	// Init config
	cfg, err := config.NewConfig(configFilename)
	if err != nil {
		panic(err)
	}

	// Init logging
	lw := &lumberjack.Logger{
		Filename:   cfg.Logfile,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}
	defer lw.Close()

	mw := io.MultiWriter(lw, os.Stdout)
	th := slog.NewTextHandler(mw, nil)
	logger := slog.New(th)
	slog.SetDefault(logger)

	// Init telegram
	tg := telegram.NewTelegram(cfg.Token)

	// Init worker
	w := worker.NewWorker(tg)

	err = w.SendMessage("Hello world")
	if err != nil {
		slog.Error("telegram send message", "error", err.Error())
		os.Exit(1)
	}
}
