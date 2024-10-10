package protocol

import (
	"context"
	"io"

	"go.lsp.dev/jsonrpc2"
)

type ClientCloser interface {
	io.Closer
}

type connSender interface {
	io.Closer

	Notify(ctx context.Context, method string, params interface{}) error
	Call(ctx context.Context, method string, params, result interface{}) error
}

type clientDispatcher struct {
	sender connSender
}

func (c *clientDispatcher) Close() error {
	return c.sender.Close()
}

func NewClientDispatcher(conn jsonrpc2.Conn) ClientCloser {
	return &clientDispatcher{sender: clientConn{conn}}
}

type clientConn struct {
	conn jsonrpc2.Conn
}

func (c clientConn) Close() error {
	return c.conn.Close()
}

func (c clientConn) Notify(ctx context.Context, method string, params interface{}) error {
	return c.conn.Notify(ctx, method, params)
}

func (c clientConn) Call(ctx context.Context, method string, params, result interface{}) error {
	_, err := c.conn.Call(ctx, method, params, result)
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return err
}
