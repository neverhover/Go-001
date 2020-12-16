package http

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Server struct {
	Name string
	Addr string
	Handler http.Handler
	Ctx  context.Context
	ins  *http.Server
}

func (s *Server) Shutdown(ctx context.Context) error {
	t := rand.New(rand.NewSource(time.Now().UnixNano()))
	sleepTime := t.Intn(8)
	if s.ins != nil {
		fmt.Printf("Server %s: shutdown start... need work more %d seconds to finish some tasks\n", s.Name, sleepTime)
		time.Sleep(time.Duration(sleepTime) * time.Second)
		s.ins.Shutdown(context.Background())
		s.ins = nil
	}
	fmt.Printf("Server %s: shutdown done\n", s.Name)
	return nil
}

func (s *Server) Start(ctx context.Context) error {
	s.Ctx = ctx
	defer s.Shutdown(ctx)
	// Init some resources
	quit := make(chan struct{})
	s.ins = &http.Server{
		Addr:    s.Addr,
		Handler: s.Handler,
	}
	// Do the real work
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Server %s: panic with error -> %+v\n", s.Name, err)
				s.ins = nil
				quit <- struct{}{}
			}
		}()
		s.run()
	}()
	// Wait event which come from parent context or get a signal by any exception
	select {
	case <-s.Ctx.Done():
		fmt.Printf("%s: get a finish event by parent context , should shutdown Server next\n", s.Name)
		return nil
	case <-quit:
		fmt.Printf("%s goroutine error, quit Server\n", s.Name)
		return nil
	}

}

func (s *Server) run() error {

	fmt.Printf("Server %s: running...\n", s.Name)

	err := s.ins.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Printf("Server %s: Server exit with error -> %s\n", s.Name, err)
		return err
	}
	return nil
}

