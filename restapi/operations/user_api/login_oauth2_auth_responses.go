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

package user_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/m3/models"
)

// LoginOauth2AuthCreatedCode is the HTTP code returned for type LoginOauth2AuthCreated
const LoginOauth2AuthCreatedCode int = 201

/*LoginOauth2AuthCreated A successful login.

swagger:response loginOauth2AuthCreated
*/
type LoginOauth2AuthCreated struct {

	/*
	  In: Body
	*/
	Payload *models.LoginResponse `json:"body,omitempty"`
}

// NewLoginOauth2AuthCreated creates LoginOauth2AuthCreated with default headers values
func NewLoginOauth2AuthCreated() *LoginOauth2AuthCreated {

	return &LoginOauth2AuthCreated{}
}

// WithPayload adds the payload to the login oauth2 auth created response
func (o *LoginOauth2AuthCreated) WithPayload(payload *models.LoginResponse) *LoginOauth2AuthCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login oauth2 auth created response
func (o *LoginOauth2AuthCreated) SetPayload(payload *models.LoginResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginOauth2AuthCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*LoginOauth2AuthDefault Generic error response.

swagger:response loginOauth2AuthDefault
*/
type LoginOauth2AuthDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewLoginOauth2AuthDefault creates LoginOauth2AuthDefault with default headers values
func NewLoginOauth2AuthDefault(code int) *LoginOauth2AuthDefault {
	if code <= 0 {
		code = 500
	}

	return &LoginOauth2AuthDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the login oauth2 auth default response
func (o *LoginOauth2AuthDefault) WithStatusCode(code int) *LoginOauth2AuthDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the login oauth2 auth default response
func (o *LoginOauth2AuthDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the login oauth2 auth default response
func (o *LoginOauth2AuthDefault) WithPayload(payload *models.Error) *LoginOauth2AuthDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login oauth2 auth default response
func (o *LoginOauth2AuthDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginOauth2AuthDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}