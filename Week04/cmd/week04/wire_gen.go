// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/neverhover/Go-001/tree/main/Week04/internal/pkg/config"
	"github.com/neverhover/Go-001/tree/main/Week04/internal/pkg/logger"
	"github.com/neverhover/Go-001/tree/main/Week04/internal/pkg/service"
	"github.com/neverhover/Go-001/tree/main/Week04/internal/pkg/storage"
	"io"
)

// Injectors from wire.go:

func InitializeCoreService(out io.Writer) (*service.CoreService, error) {
	configConfig := config.NewConfig()
	string2 := logger.GetLoggerLevel(configConfig)
	loggerLogger, err := logger.NewLoggerIns(out, string2)
	if err != nil {
		return nil, err
	}
	db, err := storage.NewConnection(configConfig)
	if err != nil {
		return nil, err
	}
	coreService := service.NewCoreService(loggerLogger, configConfig, db)
	return coreService, nil
}

// wire.go:

var configSet = wire.NewSet(config.NewConfig)
