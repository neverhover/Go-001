package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Service struct {
	Name string
	Addr string
	Ctx  context.Context
	ins  *http.Server
}

func main() {
	var err error
	rootCtx := context.Background()
	g, cancelCtx := errgroup.WithContext(rootCtx)
	// 不要使用errgroup返回的context

	serv01 := Service{
		Name: "httpServer-01",
		Addr: "0.0.0.0:10001",
		Ctx:  cancelCtx,
	}
	serv02 := Service{
		Name: "httpServer-02",
		Addr: "0.0.0.0:10002",
		Ctx:  cancelCtx,
	}
	g.Go(serv01.Start)
	g.Go(serv02.Start)
	fmt.Printf("Main: waiting all services work done\n")

	errs := make(chan error, 2)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)

	}()
	fmt.Printf("Main: all services started\n")
	select {
	case <-errs:
		// We got a signal, now shutdown all the services
		fmt.Printf("Main: Try to exit\n")
		go serv01.Shutdown()
		go serv02.Shutdown()
		// Waiting them
		err = g.Wait()
		if err != nil {
			fmt.Printf("Work error find: %+v\n", err)
			os.Exit(1)
		}
	}
	fmt.Printf("Main: all work done\n")
}

func (s *Service) Shutdown()  {
	t := rand.New(rand.NewSource(time.Now().UnixNano()))
	sleepTime := t.Intn(8)
	fmt.Printf("Service %s: shutdown start... need work more %d seconds\n", s.Name, sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Printf("Service %s: shutdown work done\n", s.Name)
	if s.ins != nil {
		s.ins.Shutdown(context.Background())
	}

}

func (s *Service) Start() error {
	quit := make(chan struct{})
	s.ins = &http.Server{
		Addr:    s.Addr,
		Handler: nil,
	}
	// Mock  panic
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Service %s: Mock panic got -> %+v\n", s.Name, err)
				quit<- struct{}{}
			}
		}()
		t := rand.New(rand.NewSource(time.Now().UnixNano()))
		sleepTime := t.Intn(8)
		time.Sleep(time.Duration(sleepTime) * time.Second)
		panic(fmt.Sprintf(fmt.Sprintf("panic %s", s.Name)))
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Service %s: panic with error -> %+v\n", s.Name, err)
			s.ins = nil
		}
	}()

	go s.run()
	select {
	case <-s.Ctx.Done():
		fmt.Printf("%s work done\n", s.Name)
		return nil
	case <-quit:
		fmt.Printf("%s goroutine error, quit service\n", s.Name)
		return errors.New("xxx eror")
	}
}

func (s *Service) run() error {

	fmt.Printf("Service %s: running...\n", s.Name)

	err := s.ins.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Printf("Service %s: service exit...", err)
		return err
	}
	return nil
}
