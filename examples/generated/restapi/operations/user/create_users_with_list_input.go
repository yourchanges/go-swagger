package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// CreateUsersWithListInputHandlerFunc turns a function with the right signature into a create users with list input handler
type CreateUsersWithListInputHandlerFunc func(CreateUsersWithListInputParams) error

func (fn CreateUsersWithListInputHandlerFunc) Handle(params CreateUsersWithListInputParams) error {
	return fn(params)
}

// CreateUsersWithListInputHandler interface for that can handle valid create users with list input params
type CreateUsersWithListInputHandler interface {
	Handle(CreateUsersWithListInputParams) error
}

// NewCreateUsersWithListInput creates a new http.Handler for the create users with list input operation
func NewCreateUsersWithListInput(ctx *middleware.Context, handler CreateUsersWithListInputHandler) *CreateUsersWithListInput {
	return &CreateUsersWithListInput{Context: ctx, Handler: handler}
}

// CreateUsersWithListInput
type CreateUsersWithListInput struct {
	Context *middleware.Context
	Params  CreateUsersWithListInputParams
	Handler CreateUsersWithListInputHandler
}

func (o *CreateUsersWithListInput) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	err := o.Handler.Handle(o.Params) // actually handle the request
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	o.Context.Respond(rw, r, route.Produces, route, nil)

}
