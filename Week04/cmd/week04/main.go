package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/neverhover/Go-001/tree/main/Week04/internal/service"
	"github.com/neverhover/Go-001/tree/main/Week04/pkg/app"
	transhttp "github.com/neverhover/Go-001/tree/main/Week04/pkg/transport/http"
	"log"
	"os"

	gw "github.com/neverhover/Go-001/tree/main/Week04/api/users"
)

func main() {

	serv, err := InitializeCoreService(os.Stdout)
	if err != nil {
		fmt.Printf("Init service error with %v\n", err)
		os.Exit(1)
	}

	gs := service.NewUserService(serv)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	err = gw.RegisterUsersServiceHandlerServer(ctx, mux, gs)
	if err != nil {
		fmt.Printf("Init service error with %v\n", err)
		os.Exit(1)
	}

	addr := "0.0.0.0:9911"
	serv01 := transhttp.Server{
		Name: "httpServer-01",
		Addr: addr,
		Handler: mux,
	}

	application := app.New()
	application.AppendServices(app.Services{Start: serv01.Start, Shutdown: serv01.Shutdown})

	if err := application.Run(); err != nil {
		log.Printf("startup failed: %v\n", err)
	}
}
