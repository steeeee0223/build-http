package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"steeeee0223.http/internal/request"
	"steeeee0223.http/internal/response"
	"steeeee0223.http/internal/server"
)

const port = 42069

func res400() []byte {
	return []byte(`<html>
  <head>
    <title>400 Bad Request</title>
  </head>
  <body>
    <h1>Bad Request</h1>
    <p>Your request honestly kinda sucked.</p>
  </body>
</html>`)
}
func res500() []byte {
	return []byte(`<html>
  <head>
    <title>500 Internal Server Error</title>
  </head>
  <body>
    <h1>Internal Server Error</h1>
    <p>Okay, you know what? This one is on me.</p>
  </body>
</html>`)
}
func res200() []byte {
	return []byte(`<html>
  <head>
    <title>200 OK</title>
  </head>
  <body>
    <h1>Success!</h1>
    <p>Your request was an absolute banger.</p>
  </body>
</html>`)
}

func main() {
	server, err := server.Serve(port, func(w *response.Writer, req *request.Request) *server.HandlerError {

		h := response.GetDefaultHeaders(0)
		b := res200()
		status := response.StatusOK

		if req.RequestLine.RequestTarget == "/dog" {
			b = res400()
			status = response.StatusBadRequest

		} else if req.RequestLine.RequestTarget == "/cat" {
			b = res500()
			status = response.StatusInternalServerError

		}

		h.Replace("Content-type", "text/html")
		h.Replace("Content-length", fmt.Sprintf("%d", len(b)))
		w.WriteStatusLine(status)
		w.WriteHeaders(*h)
		w.WriteBody(b)
		return nil
	})
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer server.Close()
	log.Println("Server started on port", port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Server gracefully stopped")
}
