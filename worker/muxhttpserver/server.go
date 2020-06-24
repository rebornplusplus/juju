// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package muxhttpserver

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/juju/errors"
	"github.com/juju/worker/v2/catacomb"

	"github.com/juju/juju/apiserver/apiserverhttp"
	"github.com/juju/juju/pki"
	pkitls "github.com/juju/juju/pki/tls"
)

type Config struct {
	Address string
	Port    string
}

type Logger interface {
	Debugf(string, ...interface{})
	Errorf(string, ...interface{})
	Infof(string, ...interface{})
}

// Server is the http server running inside this worker handling requests to
// the http mux
type Server struct {
	catacomb catacomb.Catacomb
	info     ServerInfo
	listener net.Listener
	logger   Logger
	Mux      *apiserverhttp.Mux
	server   *http.Server
}

// ServerInfo represents a small interface for informining interested party's
// through manifolds outputs the listening information of the server
type ServerInfo interface {
	Port() string
	PortInt() (int, error)
}

type serverInfo struct {
	port string
}

var (
	defaultPort = "17071"
)

func NewServer(authority pki.Authority, logger Logger, conf Config) (*Server, error) {
	mux := apiserverhttp.NewMux()

	tlsConfig := &tls.Config{
		GetCertificate: pkitls.AuthoritySNITLSGetter(authority, logger),
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", conf.Address, conf.Port))
	if err != nil {
		return nil, errors.Annotate(err, "creating mux http server listener")
	}

	httpServ := &http.Server{
		Handler:   mux,
		TLSConfig: tlsConfig,
	}

	server := &Server{
		info:     &serverInfo{conf.Port},
		listener: listener,
		logger:   logger,
		Mux:      mux,
		server:   httpServ,
	}

	if err := catacomb.Invoke(catacomb.Plan{
		Site: &server.catacomb,
		Work: server.loop,
	}); err != nil {
		return server, errors.Trace(err)
	}
	return server, nil
}

func DefaultConfig() Config {
	return Config{
		Port: defaultPort,
	}
}

// Kill implements the worker interface
func (s *Server) Kill() {
	s.catacomb.Kill(nil)
}

func (s *serverInfo) Port() string {
	return s.port
}

func (s *serverInfo) PortInt() (int, error) {
	return net.LookupPort("tcp", s.port)
}

func (s *Server) Port() string {
	splits := strings.Split(s.listener.Addr().String(), ":")
	if len(splits) == 0 {
		return ""
	}
	return splits[len(splits)-1]
}

// Wait implements the worker interface
func (s *Server) Wait() error {
	return s.catacomb.Wait()
}

func (s *Server) Info() ServerInfo {
	return s.info
}

func (s *Server) loop() error {
	httpCh := make(chan error)

	go func() {
		s.logger.Infof("starting http server on %s", s.listener.Addr())
		httpCh <- s.server.ServeTLS(s.listener, "", "")
		close(httpCh)
	}()

	for {
		select {
		case <-s.catacomb.Dying():
			s.server.Close()
		case err := <-httpCh:
			if err != nil && err != http.ErrServerClosed {
				return err
			}
			return s.catacomb.ErrDying()
		}
	}
}
