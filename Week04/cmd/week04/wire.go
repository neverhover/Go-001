// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"github.com/neverhover/Go-001/tree/main/Week04/internal/pkg/config"
	"github.com/neverhover/Go-001/tree/main/Week04/internal/pkg/service"
	"github.com/neverhover/Go-001/tree/main/Week04/internal/pkg/storage"
	"github.com/neverhover/Go-001/tree/main/Week04/internal/pkg/logger"
	"io"
)

//var CoreService = wire.NewSet(service.NewCoreService, storage.NewConnection, logger.NewLoggerIns, config.NewConfig)

var configSet = wire.NewSet(config.NewConfig)

func InitializeCoreService(out io.Writer) (*service.CoreService, error) {
	wire.Build(service.NewCoreService, logger.GetLoggerLevel, logger.NewLoggerIns, configSet, storage.NewConnection)
	return &service.CoreService{}, nil
}
