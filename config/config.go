package config

import (
	"github.com/gokits/cfg"
	"github.com/gokits/gotools"
	"github.com/gokits/stdlogger"
)


type Config struct {
	CrossConfig        CrossConfig  `json:"crossConfig" validate:"required"`
	HTTPServer         HTTPServer   `json:"httpServer" validate:"required"`
	Mysql              Mysql        `json:"mysql" validate:"required"`
	Logger             Logger       `json:"logger" validate:"required"`
}

type Mysql struct {
	User            string           `json:"user" validate:"required"`
	Host            string           `json:"host" validate:"required"`
	Database        string           `json:"database" validate:"required"`
	Password        string           `json:"password" validate:"required"`
	Port            string           `json:"port" validate:"required"`
	MaxOpenConns    int              `json:"maxOpenConns" validate:"required"`
	MaxIdleConns    int              `json:"maxIdleConns" validate:"required"`
	ConnMaxLifetime gotools.Duration `json:"connMaxLifetime" validate:"required"`
}

type Logger struct {
	Dir        string `json:"dir" validate:"required"`
	RemainDays int    `json:"remainDays" validate:"required"`
}

type CrossConfig struct {
	AccessControlAllowOrigin      string   `json:"Access-Control-Allow-Origin" validate:"required"`
	AccessControlAllowMethods     string   `json:"Access-Control-Allow-Methods" validate:"required"`
	AccessControlAllowHeaders     string   `json:"Access-Control-Allow-Headers" validate:"required"`
	AccessControlExposeHeaders    string   `json:"Access-Control-Expose-Headers" validate:"required"`
	AccessControlAllowCredentials string   `json:"Access-Control-Allow-Credentials" validate:"required"`
}

// HTTPServer configuration for HttpServer
type HTTPServer struct {
	// ListenAddr addr to listen, ":8080" for example
	ListenAddr string `json:"listenAddr" validate:"required,tcp4_addr"`
	// GraceShutdownPeriod time to wait before shutting down the server forcely
	GraceShutdownPeriod gotools.Duration `json:"graceShutdownPeriod" validate:"min=0"`
	// Dist dir
	DistFilePath string `json:"distFilePath" validate:"required"`
}


func Init(configfile string, logger stdlogger.LeveledLogger) (err error) {
	cfg.MustRegisterFile(&Config{}, configfile, cfg.WithDefaultConfiguration().WithLogger(logger))
	if err = cfg.WaitSyncedAll(); err != nil {
		logger.Errorf("Fail to synced all")
		return
	}
	return
}

func Finish() {
	cfg.Final()
}

func GetConfig() *Config {
	return cfg.MustGet(&Config{}).(*Config)
}