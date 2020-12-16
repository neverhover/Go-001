package service

import (
	"github.com/mainflux/mainflux/logger"
	"github.com/neverhover/Go-001/tree/main/Week04/internal/pkg/config"
	"gorm.io/gorm"
)

type Service interface {
	DB() *gorm.DB
	Logger() logger.Logger
	Config() *config.Config
}

type CoreService struct {
	logIns logger.Logger
	cfg    *config.Config
	db     *gorm.DB
}

func (c CoreService) DB() *gorm.DB {
	return c.db
}

func (c CoreService) Logger() logger.Logger {
	return c.logIns
}

func (c CoreService) Config() *config.Config {
	return c.cfg
}

var _ Service = (*CoreService)(nil)

func NewCoreService(logIns logger.Logger, cfg *config.Config, db *gorm.DB) *CoreService {
	return &CoreService{
		logIns: logIns,
		cfg:    cfg,
		db:     db,
	}
}
