package comet

import (
	"context"
	"fmt"
	"net"
)

type connection struct {
	conn   ConnWrapper
	cancel context.CancelFunc
}

type Handler struct {
	ctx   context.Context
	conns map[string]connection
}

var (
	handler = Handler{
		ctx:   context.Background(),
		conns: map[string]connection{},
	}
)

func Handle(conn net.Conn) {
	connInfo := fmt.Sprintf("%s-%s", conn.LocalAddr(), conn.RemoteAddr())
	if _, ok := handler.conns[connInfo]; !ok {
		connectObj := connection{}
		var ctx context.Context
		ctx, connectObj.cancel = context.WithCancel(context.Background())
		connectObj.conn = ConnWrapper{
			conn:    conn,
			msgChan: make(chan []byte),
			ctx:     ctx,
		}
		handler.conns[connInfo] = connectObj
	}
	go handler.conns[connInfo].conn.Handle()
}
