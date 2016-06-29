// Copyright (c) 2016 Uber Technologies, Inc.
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

package yarpc

import (
	"golang.org/x/net/context"

	"github.com/yarpc/yarpc-go/transport"
)

// CallReqMeta contains information about an outgoing YARPC request.
type CallReqMeta interface {
	Procedure(string) CallReqMeta
	Headers(Headers) CallReqMeta

	GetContext() context.Context
	GetProcedure() string
	GetHeaders() Headers
}

// ReqMeta contains information about an incoming YARPC request.
type ReqMeta interface {
	Caller() string
	Context() context.Context
	Encoding() transport.Encoding
	Headers() Headers
	Procedure() string
	Service() string
}

// NewReqMeta constructs a CallReqMeta with the given Context.
//
// The context MUST NOT be nil.
func NewReqMeta(ctx context.Context) CallReqMeta {
	if ctx == nil {
		panic("invalid usage of ReqMeta: context cannot be nil")
	}
	return &callReqMeta{ctx: ctx}
}

type callReqMeta struct {
	ctx       context.Context
	procedure string
	headers   Headers
}

func (r *callReqMeta) Procedure(p string) CallReqMeta {
	r.procedure = p
	return r
}

func (r *callReqMeta) Headers(h Headers) CallReqMeta {
	r.headers = h
	return r
}

func (r *callReqMeta) GetContext() context.Context {
	return r.ctx
}

func (r *callReqMeta) GetProcedure() string {
	return r.procedure
}

func (r *callReqMeta) GetHeaders() Headers {
	return r.headers
}