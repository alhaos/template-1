package main

import (
	"flag"
	"io"
	"log/slog"
	"os"

	"github.com/alhaos/okPublisher/internal/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {

	// Get config path from flag -config
	configFilenamePointer := flag.String("config", "config/config.yml", "okPublisher config file")
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

	mw := io.MultiWriter(lw, os.Stdout)
	th := slog.NewTextHandler(mw, nil)
	logger := slog.New(th)
	slog.SetDefault(logger)

}
