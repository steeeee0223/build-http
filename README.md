# Build HTTP from TCP

[Full course by Primagen](https://youtu.be/FknTw9bJsXM?si=n1V1CYWtO51UHTf8)

## Commands

```bash
go run ./cmd/tcplistener | tee /tmp/tcp.txt

nc -v localhost 42069

# run all tests
go run ./...
```

## Reference

- [HTTP 1.1](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/Connection_management_in_HTTP_1.x)
- [RFC 9110](https://datatracker.ietf.org/doc/html/rfc9110)
- [RFC 9112](https://datatracker.ietf.org/doc/html/rfc9112)
