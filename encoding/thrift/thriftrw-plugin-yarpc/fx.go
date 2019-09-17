// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"path/filepath"

	"go.uber.org/thriftrw/plugin"
)

const fxDocTemplate = `
// Code generated by thriftrw-plugin-yarpc
// @generated

<$pkgname := printf "%sfx" (lower .Name)>
<$serverpkgname := printf "%sserver" (lower .Name)>
// Package <$pkgname> provides better integration for Fx for services
// implementing or calling <.Name>.
//
// Clients
//
// If you are making requests to <.Name>, use the Client function to inject a
// <.Name> client into your container.
//
// 	fx.Provide(<$pkgname>.Client("..."))
//
// Servers
//
// If you are implementing <.Name>, provide a <$serverpkgname>.Interface into
// the container and use the Server function.
//
// Given,
//
// 	func New<.Name>Handler() <$serverpkgname>.Interface
//
// You can do the following to have the procedures of <.Name> made available
// to an Fx application.
//
// 	fx.Provide(
// 		New<.Name>Handler,
// 		<$pkgname>.Server(),
// 	)
package <$pkgname>
`

const fxClientTemplate = `
// Code generated by thriftrw-plugin-yarpc
// @generated

<$pkgname := printf "%sfx" (lower .Name)>
package <$pkgname>

<$yarpc := import "go.uber.org/yarpc">
<$thrift := import "go.uber.org/yarpc/encoding/thrift">
<$client := import .ClientPackagePath>
<$fx := import "go.uber.org/fx">

// Params defines the dependencies for the <.Name> client.
type Params struct {
	<$fx>.In

	Provider <$yarpc>.ClientConfig
}

// Result defines the output of the <.Name> client module. It provides a
// <.Name> client to an Fx application.
type Result struct {
	<$fx>.Out

	Client <$client>.Interface

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// Client provides a <.Name> client to an Fx application using the given name
// for routing.
//
// 	fx.Provide(
// 		<$pkgname>.Client("..."),
// 		newHandler,
// 	)
func Client(name string, opts ...<$thrift>.ClientOption) interface{} {
	return func(p Params) Result {
		client := <$client>.New(p.Provider.ClientConfig(name), opts...)
		return Result{Client: client}
	}
}`

const fxServerTemplate = `
// Code generated by thriftrw-plugin-yarpc
// @generated

<$pkgname := printf "%sfx" (lower .Name)>
package <$pkgname>

<$transport := import "go.uber.org/yarpc/api/transport">
<$thrift := import "go.uber.org/yarpc/encoding/thrift">
<$server := import .ServerPackagePath>
<$fx := import "go.uber.org/fx">

// ServerParams defines the dependencies for the <.Name> server.
type ServerParams struct {
	<$fx>.In

	Handler <$server>.Interface
}

// ServerResult defines the output of <.Name> server module. It provides the
// procedures of a <.Name> handler to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group. Dig 1.2 or newer
// must be used for this feature to work.
type ServerResult struct {
	<$fx>.Out

	Procedures []<$transport>.Procedure ` + "`group:\"yarpcfx\"`" + `
	ThriftProcedures []<$transport>.Procedure ` + "`group:\"thrift\"`" + `
}

// Server provides procedures for <.Name> to an Fx application. It expects a
// <$pkgname>.Interface to be present in the container.
//
// 	fx.Provide(
// 		func(h *My<.Name>Handler) <$server>.Interface {
// 			return h
// 		},
// 		<$pkgname>.Server(),
// 	)
func Server(opts ...<$thrift>.RegisterOption) interface{} {
	return func(p ServerParams) ServerResult {
		procedures := <$server>.New(p.Handler, opts...)
		return ServerResult{
			Procedures: procedures,
			ThriftProcedures: procedures,
		}
	}
}
`

func fxGenerator(data *templateData, files map[string][]byte) (err error) {
	packageName := filepath.Base(data.FxPackagePath())

	// kv.thrift => .../kv/keyvaluefx/doc.go
	docPath := filepath.Join(data.Module.Directory, packageName, "doc.go")
	files[docPath], err = plugin.GoFileFromTemplate(docPath, fxDocTemplate, data, templateOptions...)

	// kv.thrift => .../kv/keyvaluefx/client.go
	clientPath := filepath.Join(data.Module.Directory, packageName, "client.go")
	files[clientPath], err = plugin.GoFileFromTemplate(clientPath, fxClientTemplate, data, templateOptions...)

	// kv.thrift => .../kv/keyvaluefx/server.go
	serverPath := filepath.Join(data.Module.Directory, packageName, "server.go")
	files[serverPath], err = plugin.GoFileFromTemplate(serverPath, fxServerTemplate, data, templateOptions...)

	return
}
