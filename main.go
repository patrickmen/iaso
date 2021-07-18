package main

import (
	"flag"
	"path/filepath"

	"go.uber.org/zap"

	"iaso/config"
	log "iaso/logger"
	"iaso/mysql"
	"iaso/server"
)

var (
	logger         *zap.SugaredLogger
	verbose        bool
)

func init() {
	flag.BoolVar(&verbose, "verbose", false, "Print log verbosely")
	flag.Parse()
}

func main() {
	if verbose {
		logger = log.DevModeRuntimeSugar
	} else {
		logger = log.RuntimeSugar
	}
	//Get the config
	if err := config.Init(logger); err != nil {
		logger.Fatalf("config init failed: %v", err)
	}
	defer config.Finish()
	cfg := config.GetConfig()

	// create log file
	err := log.Init(filepath.Join(cfg.Logger.Dir, "iaso-runtime.log"), cfg.Logger.RemainDays,
		filepath.Join(cfg.Logger.Dir, "iaso-access.log"), cfg.Logger.RemainDays)
	if err != nil {
		logger.Errorf("Logger init failed: %v", err)
		return
	}
	defer log.Final()

	// connect to database
	mysql.Client, err = mysql.NewMysqlClient(cfg.Mysql, logger)
	if err != nil {
		logger.Errorf("Fail to connect to mysql client: %s", err.Error())
		return
	}

	router := server.Init(logger, verbose, cfg.CrossConfig, cfg.HTTPServer.DistFilePath)
	server.Run(router, cfg, logger)

}
