package mysql

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"xorm.io/xorm"

	"iaso/config"
)

const ErrorDuplicateEntry uint16 = 1062

var Client xorm.Interface

type Mysql struct {
	*xorm.Engine
	Logger *zap.SugaredLogger
}

func NewMysqlClient(conf config.Mysql, logger *zap.SugaredLogger) (xorm.Interface, error) {
	mysqlLogger := logger.Named("MysqlClient")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	mysqlLogger.Debugf("DSN: %s", dsn)

	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		mysqlLogger.Errorf("Fail to create database engine %s: %s", dsn, err.Error())
		return nil, err
	}

	err = engine.Ping()
	if err != nil {
		mysqlLogger.Errorf("Fail to ping database %s: %s", dsn, err.Error())
		return nil, err
	}

	engine.SetConnMaxLifetime(conf.ConnMaxLifetime.Duration)
	engine.SetMaxOpenConns(conf.MaxOpenConns)
	engine.SetMaxIdleConns(conf.MaxIdleConns)
	engine.TZLocation = time.FixedZone("Asia/Shanghai", 8*60*60)

	return &Mysql{
		Engine:  engine,
		Logger:  mysqlLogger,
	}, nil
}

