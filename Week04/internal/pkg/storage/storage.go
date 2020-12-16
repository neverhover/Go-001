package storage

import (
	"fmt"
	"github.com/neverhover/Go-001/tree/main/Week04/internal/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"time"
)

func convertDBlogLevel(level string) logger.LogLevel {
	levelStr := strings.ToLower(level)
	logLevel := logger.Silent
	if levelStr == "slient" {
		logLevel = logger.Silent
	} else if levelStr == "error" {
		logLevel = logger.Error
	} else if levelStr == "warn" {
		logLevel = logger.Warn
	} else if levelStr == "info" {
		logLevel = logger.Info
	}
	return logLevel
}

func NewConnection(appConfig *config.Config) (*gorm.DB, error) {
	cfg := appConfig.DbConfig
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s sslcert=%s sslkey=%s sslrootcert=%s TimeZone=Asia/Shanghai",
		cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Pass, cfg.SSLMode, cfg.SSLCert, cfg.SSLKey, cfg.SSLRootCert)

	//dsn := "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	logLevel := convertDBlogLevel(cfg.LogLevel)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	fmt.Printf("Database Log level is %s(%d)\n", cfg.LogLevel, logLevel)
	//db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	//fmt.Printf("Init DB Stats %+v\n", sqlDB.Stats())

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Printf("DB Stats %+v\n", sqlDB.Stats())
	return db, nil
}

