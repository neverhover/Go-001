package app

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Services struct {
	Start    func(context.Context) error
	Shutdown func(context.Context) error
}

// Option is an application option.
type Option func(o *options)

type options struct {
	startTimeout time.Duration
	stopTimeout  time.Duration

	sigs  []os.Signal
	sigFn func(*App, os.Signal)
}

type App struct {
	opts     options
	ctx      context.Context
	services []Services
	cancel   func()
}

func New(opts ...Option) *App {
	options := options{
		startTimeout: time.Second * 30,
		stopTimeout:  time.Second * 30,
		sigs: []os.Signal{
			syscall.SIGTERM,
			syscall.SIGQUIT,
			syscall.SIGINT,
		},
		sigFn: func(a *App, sig os.Signal) {
			switch sig {
			case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
				a.Stop()
			default:
			}
		},
	}
	for _, o := range opts {
		o(&options)
	}
	return &App{
		opts: options,
		ctx: context.Background(),
	}
}

func (a *App) AppendServices(service Services) {
	a.services = append(a.services, service)
}

func (a *App) Run() error {
	var ctx context.Context
	ctx, a.cancel = context.WithCancel(a.ctx)
	g, ctx := errgroup.WithContext(ctx)
	fmt.Printf("Application running\n")
	for _, srv := range a.services {
		fmt.Printf("Services %v\n", srv)
		srv := srv
		if srv.Shutdown != nil {
			g.Go(func() error {
				<-ctx.Done() // wait for stop signal
				stopCtx, cancel := context.WithTimeout(context.Background(), a.opts.stopTimeout)
				defer cancel()
				return srv.Shutdown(stopCtx)
			})
		}
		if srv.Start != nil {
			g.Go(func() error {
				startCtx, cancel := context.WithTimeout(context.Background(), a.opts.startTimeout)
				defer cancel()
				return srv.Start(startCtx)
			})
		}
	}
	if len(a.opts.sigs) == 0 {
		return g.Wait()
	}
	c := make(chan os.Signal, len(a.opts.sigs))
	signal.Notify(c, a.opts.sigs...)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case sig := <-c:
				if a.opts.sigFn != nil {
					a.opts.sigFn(a, sig)
				}
			}
		}
	})
	return g.Wait()
}

func (a App) Stop() {
	if a.cancel != nil {
		a.cancel()
	}
}
