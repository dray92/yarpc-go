// Code generated by thriftrw-plugin-yarpc
// @generated

package metafx

import (
	"go.uber.org/fx"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/thrift"
	"go.uber.org/yarpc/x/shardproxy/internal/gen/go/github.com/uber/tchannel/meta/metaserver"
)

// ServerParams defines the dependencies for the Meta server.
type ServerParams struct {
	fx.In

	Handler metaserver.Interface
}

// ServerResult defines the output of Meta server module. It provides the
// procedures of a Meta handler to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group. Dig 1.2 or newer
// must be used for this feature to work.
type ServerResult struct {
	fx.Out

	Procedures []transport.Procedure `group:"yarpcfx"`
}

// Server provides procedures for Meta to an Fx application. It expects a
// metafx.Interface to be present in the container.
//
// 	fx.Provide(
// 		func(h *MyMetaHandler) metaserver.Interface {
// 			return h
// 		},
// 		metafx.Server(),
// 	)
func Server(opts ...thrift.RegisterOption) interface{} {
	return func(p ServerParams) ServerResult {
		procedures := metaserver.New(p.Handler, opts...)
		return ServerResult{Procedures: procedures}
	}
}