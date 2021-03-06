package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/casualjim/go-swagger/errors"
	"github.com/casualjim/go-swagger/httpkit/middleware"
	"github.com/casualjim/go-swagger/strfmt"
)

// DeleteUserParams contains all the bound params for the delete user operation
// typically these are obtained from a http.Request
type DeleteUserParams struct {
	// The name that needs to be deleted
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	if err := o.bindUsername(route.Params.Get("username"), route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteUserParams) bindUsername(raw string, formats strfmt.Registry) error {

	o.Username = raw

	if err := o.validateUsername(formats); err != nil {
		return err
	}

	return nil
}

func (o *DeleteUserParams) validateUsername(formats strfmt.Registry) error {

	return nil
}
