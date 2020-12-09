package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Service struct {
	Name string
	Addr string
	Ctx  context.Context
	ins  *http.Server
}

func (s *Service) Shutdown() {
	t := rand.New(rand.NewSource(time.Now().UnixNano()))
	sleepTime := t.Intn(8)
	if s.ins != nil {
		fmt.Printf("Service %s: shutdown start... need work more %d seconds to finish some tasks\n", s.Name, sleepTime)
		time.Sleep(time.Duration(sleepTime) * time.Second)
		s.ins.Shutdown(context.Background())
		s.ins = nil
	}
	fmt.Printf("Service %s: shutdown done\n", s.Name)
}

func (s *Service) Start() error {
	defer s.Shutdown()
	// Init some resources
	quit := make(chan struct{})
	s.ins = &http.Server{
		Addr:    s.Addr,
		Handler: nil,
	}
	// Do the real work
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Service %s: panic with error -> %+v\n", s.Name, err)
				s.ins = nil
				quit <- struct{}{}
			}
		}()
		s.run()
	}()
	// Wait event which come from parent context or get a signal by any exception
	select {
	case <-s.Ctx.Done():
		fmt.Printf("%s: get a finish event by parent context , should shutdown service next\n", s.Name)
		return nil
	case <-quit:
		fmt.Printf("%s goroutine error, quit service\n", s.Name)
		return nil
	}

}

func (s *Service) run() error {

	fmt.Printf("Service %s: running...\n", s.Name)

	err := s.ins.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Printf("Service %s: service exit with error -> %s\n", s.Name, err)
		return err
	}
	return nil
}

