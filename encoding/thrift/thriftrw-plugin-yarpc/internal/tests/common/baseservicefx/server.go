// Code generated by thriftrw-plugin-yarpc
// @generated

package baseservicefx

import (
	fx "go.uber.org/fx"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	baseserviceserver "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/common/baseserviceserver"
)

// ServerParams defines the dependencies for the BaseService server.
type ServerParams struct {
	fx.In

	Handler baseserviceserver.Interface
}

// ServerResult defines the output of BaseService server module. It provides the
// procedures of a BaseService handler to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group. Dig 1.2 or newer
// must be used for this feature to work.
type ServerResult struct {
	fx.Out

	Procedures       []transport.Procedure `group:"yarpcfx"`
	ThriftProcedures []transport.Procedure `group:"thrift"`
}

// Server provides procedures for BaseService to an Fx application. It expects a
// baseservicefx.Interface to be present in the container.
//
// 	fx.Provide(
// 		func(h *MyBaseServiceHandler) baseserviceserver.Interface {
// 			return h
// 		},
// 		baseservicefx.Server(),
// 	)
func Server(opts ...thrift.RegisterOption) interface{} {
	return func(p ServerParams) ServerResult {
		procedures := baseserviceserver.New(p.Handler, opts...)
		return ServerResult{
			Procedures:       procedures,
			ThriftProcedures: procedures,
		}
	}
}
