// Package tgtest provides test Telegram server for end-to-end test.
package tgtest

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"net"

	"github.com/gotd/td/transport"

	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/internal/proto"
)

type Server struct {
	server *transport.Server

	key     *rsa.PrivateKey
	cipher  crypto.Cipher
	handler Handler

	ctx    context.Context
	cancel context.CancelFunc

	users *users
}

func (s *Server) Key() *rsa.PublicKey {
	return &s.key.PublicKey
}

func (s *Server) Addr() net.Addr {
	return s.server.Addr()
}

func (s *Server) Start() {
	go func() {
		_ = s.serve()
	}()
}

func (s *Server) Close() {
	if s.cancel != nil {
		s.cancel()
	}

	_ = s.server.Close()
}

func NewServer(ctx context.Context, tb TB, codec transport.Codec, h Handler) *Server {
	s := NewUnstartedServer(ctx, tb, codec)
	s.SetHandler(h)
	s.Start()
	return s
}

func NewUnstartedServer(ctx context.Context, tb TB, codec transport.Codec) *Server {
	k, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(ctx)
	s := &Server{
		server: transport.NewCustomServer(codec, newLocalListener()),
		key:    k,
		cipher: crypto.NewServerCipher(rand.Reader),
		ctx:    ctx,
		cancel: cancel,
		users:  newUsers(),
	}
	return s
}

func (s *Server) SetHandler(handler Handler) {
	s.handler = handler
}

func (s *Server) checkMsgID(id int64) error {
	if proto.MessageID(id).Type() != proto.MessageFromClient {
		return errors.New("bad msg type")
	}
	return nil
}

func (s *Server) serve() error {
	return s.server.Serve(s.ctx, s.serveConn)
}
