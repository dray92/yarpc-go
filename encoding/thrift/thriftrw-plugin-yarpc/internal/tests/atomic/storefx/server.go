// Code generated by thriftrw-plugin-yarpc
// @generated

package storefx

import (
	fx "go.uber.org/fx"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	storeserver "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/atomic/storeserver"
)

// ServerParams defines the dependencies for the Store server.
type ServerParams struct {
	fx.In

	Handler storeserver.Interface
}

// ServerResult defines the output of Store server module. It provides the
// procedures of a Store handler to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group. Dig 1.2 or newer
// must be used for this feature to work.
type ServerResult struct {
	fx.Out

	Procedures       []transport.Procedure `group:"yarpcfx"`
	ThriftProcedures []transport.Procedure `group:"thrift"`
}

// Server provides procedures for Store to an Fx application. It expects a
// storefx.Interface to be present in the container.
//
// 	fx.Provide(
// 		func(h *MyStoreHandler) storeserver.Interface {
// 			return h
// 		},
// 		storefx.Server(),
// 	)
func Server(opts ...thrift.RegisterOption) interface{} {
	return func(p ServerParams) ServerResult {
		procedures := storeserver.New(p.Handler, opts...)
		return ServerResult{
			Procedures:       procedures,
			ThriftProcedures: procedures,
		}
	}
}
