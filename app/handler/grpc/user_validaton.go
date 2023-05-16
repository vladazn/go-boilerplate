package grpc

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	frontoffice "github.com/vladazn/go-boilerplate/api/client"
)

func validateUpdateUsernameRequest(req *frontoffice.SetUsernameRequest) error {
	errs := validation.Errors{
		"username": validation.Validate(req.GetUsername(),
			validation.Required,
			validation.Length(5, 20),
		),
	}

	return errs.Filter()
}
