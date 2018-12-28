package backend

import (
	"context"
	"github.com/kooksee/grpcweb-boilerplate/proto/server"
)

// Backend should be used to implement the server interface
// exposed by the generated server proto.
type Backend struct {
}

func (t *Backend) Foo(ctx context.Context, request *server.FooRequest) (resp *server.FooResponse, err error) {
	resp = &server.FooResponse{Text: request.Text}
	return
}

// Ensure struct implements interface
var _ server.BackendServer = (*Backend)(nil)
