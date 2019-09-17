// Code generated by thriftrw-plugin-yarpc
// @generated

package foofx

import (
	fx "go.uber.org/fx"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	fooserver "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/extends/fooserver"
)

// ServerParams defines the dependencies for the Foo server.
type ServerParams struct {
	fx.In

	Handler fooserver.Interface
}

// ServerResult defines the output of Foo server module. It provides the
// procedures of a Foo handler to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group. Dig 1.2 or newer
// must be used for this feature to work.
type ServerResult struct {
	fx.Out

	Procedures       []transport.Procedure `group:"yarpcfx"`
	ThriftProcedures []transport.Procedure `group:"thrift"`
}

// Server provides procedures for Foo to an Fx application. It expects a
// foofx.Interface to be present in the container.
//
// 	fx.Provide(
// 		func(h *MyFooHandler) fooserver.Interface {
// 			return h
// 		},
// 		foofx.Server(),
// 	)
func Server(opts ...thrift.RegisterOption) interface{} {
	return func(p ServerParams) ServerResult {
		procedures := fooserver.New(p.Handler, opts...)
		return ServerResult{
			Procedures:       procedures,
			ThriftProcedures: procedures,
		}
	}
}
