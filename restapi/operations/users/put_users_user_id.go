// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutUsersUserIDHandlerFunc turns a function with the right signature into a put users user ID handler
type PutUsersUserIDHandlerFunc func(PutUsersUserIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutUsersUserIDHandlerFunc) Handle(params PutUsersUserIDParams) middleware.Responder {
	return fn(params)
}

// PutUsersUserIDHandler interface for that can handle valid put users user ID params
type PutUsersUserIDHandler interface {
	Handle(PutUsersUserIDParams) middleware.Responder
}

// NewPutUsersUserID creates a new http.Handler for the put users user ID operation
func NewPutUsersUserID(ctx *middleware.Context, handler PutUsersUserIDHandler) *PutUsersUserID {
	return &PutUsersUserID{Context: ctx, Handler: handler}
}

/*
	PutUsersUserID swagger:route PUT /users/{user_id} users putUsersUserId

update user

update user
*/
type PutUsersUserID struct {
	Context *middleware.Context
	Handler PutUsersUserIDHandler
}

func (o *PutUsersUserID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutUsersUserIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
