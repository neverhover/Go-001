package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
)

const (
	ExitCodeSuccess = iota
	ExitCodeFailedStartup
	ExitCodeForceQuit
	ExitCodeFailedQuit
)


func TrapSignal(cb func(string)) {
	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt)

		for i := 0; true; i++ {
			<-shutdown

			if i > 0 {
				fmt.Printf("Main: Exit without waiting workers done\n")
				os.Exit(ExitCodeForceQuit)
			}
			//
			go cb("SIGINT")
		}
	}()
}


func main() {
	var err error
	rootCtx, rootCancelFunc := context.WithCancel(context.Background())
	g, _ := errgroup.WithContext(rootCtx)
	// 不要使用errgroup返回的context，错误的例子可以参考error_example中示范

	serv01 := Service{
		Name: "httpServer-01",
		Addr: "0.0.0.0:10001",
		Ctx:  rootCtx,
	}
	serv02 := Service{
		Name: "httpServer-02",
		Addr: "0.0.0.0:10002",
		Ctx:  rootCtx,
	}
	g.Go(serv01.Start)
	g.Go(serv02.Start)

	fmt.Printf("Main: all services started\n")

	TrapSignal(func(signal string) {
		fmt.Printf("Main: Try to exit by signal %s\n", signal)
		rootCancelFunc()
	})
	// Wait them to work done
	err = g.Wait()
	if err != nil {
		fmt.Printf("Main: service error find: %+v\n", err)
		os.Exit(1)
	}
	// It's an normal signal exit.. but should wait all workers to finish their works before main exit.
	/*
	errs := make(chan error, 2)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)

	}()
	select {
	case <-errs:
		// We got a signal, now shutdown all the services
		fmt.Printf("Main: Try to exit\n")
		rootCancelFunc()
		// Wait them to work done
		err = g.Wait()
		if err != nil {
			fmt.Printf("Main: service error find: %+v\n", err)
			os.Exit(1)
		}
	}
	 */
	fmt.Printf("Main: work done\n")
}
