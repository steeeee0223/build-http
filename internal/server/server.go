package server

import (
	"fmt"
	"io"
	"net"

	"steeeee0223.http/internal/request"
	"steeeee0223.http/internal/response"
)

type HandlerError struct {
	StatusCode response.StatusCode
	Message    string
}
type Handler func(w *response.Writer, req *request.Request) *HandlerError

type Server struct {
	closed  bool
	handler Handler
}

func runConnection(s *Server, conn io.ReadWriteCloser) {
	defer conn.Close()

	w := response.NewWriter(conn)

	req, err := request.RequestFromReader(conn)
	if err != nil {
		w.WriteStatusLine(response.StatusBadRequest)
		w.WriteHeaders(*response.GetDefaultHeaders(0))
		return
	}

	s.handler(w, req)

}

func runServer(s *Server, listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if s.closed {
			return
		}

		if err != nil {
			return
		}
		go runConnection(s, conn)
	}
}

const port = 42069

func Serve(port uint16, handler Handler) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	server := &Server{
		closed:  false,
		handler: handler,
	}
	go runServer(server, listener)

	return server, nil
}

func (s *Server) Close() error {
	s.closed = true
	return nil
}
