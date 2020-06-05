// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Kubernetes Cloud
// Copyright (c) 2020 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package admin_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// StartMirroringHandlerFunc turns a function with the right signature into a start mirroring handler
type StartMirroringHandlerFunc func(StartMirroringParams) middleware.Responder

// Handle executing the request and returning a response
func (fn StartMirroringHandlerFunc) Handle(params StartMirroringParams) middleware.Responder {
	return fn(params)
}

// StartMirroringHandler interface for that can handle valid start mirroring params
type StartMirroringHandler interface {
	Handle(StartMirroringParams) middleware.Responder
}

// NewStartMirroring creates a new http.Handler for the start mirroring operation
func NewStartMirroring(ctx *middleware.Context, handler StartMirroringHandler) *StartMirroring {
	return &StartMirroring{Context: ctx, Handler: handler}
}

/*StartMirroring swagger:route POST /mirror AdminAPI startMirroring

Start Mirroring

*/
type StartMirroring struct {
	Context *middleware.Context
	Handler StartMirroringHandler
}

func (o *StartMirroring) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewStartMirroringParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}