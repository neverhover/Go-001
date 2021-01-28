package comet

import (
	"bufio"
	"context"
	"fmt"
	"github.com/neverhover/Go-001/tree/main/Week09/pkg/logic"
	"log"
	"net"
)

type ConnWrapper struct {
	ctx     context.Context
	conn    net.Conn
	rr      *bufio.Reader
	wr      *bufio.Writer
	msgChan chan []byte
}

func (cw ConnWrapper) Handle() {
	defer cw.Stop()
	cw.rr = bufio.NewReader(cw.conn)
	cw.wr = bufio.NewWriter(cw.conn)

	go cw.write()
	go cw.read()

	select {
	case <-cw.ctx.Done():
		return
	}
}

func (cw ConnWrapper) Stop() {
	cw.conn.Close()
}

func (cw ConnWrapper) handle(data []byte) error {
	result, err := logic.Handle(data)
	if err != nil {
		fmt.Printf("Handle msg error")
		return err
	}
	if result == nil {
		return nil
	}
	cw.msgChan <- result
	return nil
}

func (cw ConnWrapper) read() {
	for {
		msg, _, err := cw.rr.ReadLine()
		if err != nil {
			log.Printf("read error: %v\n", err)
			return
		}
		// Do logic
		go cw.handle(msg)
	}

}

func (cw ConnWrapper) write() {
	for {
		select {
		case data := <-cw.msgChan:

			_, err := cw.wr.Write(data)
			if err != nil {
				fmt.Printf("Write msg [%s] to client error %s\n", data, err)
			} else {
				fmt.Printf("Write msg [%s] to client\n", data)
			}
			if err := cw.wr.Flush(); err != nil {
				fmt.Printf("Flush write error %s\n", err)
			}
		}
	}
}
