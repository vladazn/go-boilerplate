package grpc

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	frontoffice "github.com/vladazn/go-boilerplate/api/client"
)

func validateAuthRequest(authRequest *frontoffice.AuthRequest) error {
	errs := validation.Errors{
		"key": validation.Validate(authRequest.GetKey(), validation.Required),
		"auth_type": validation.Validate(authRequest.GetAuthType(),
			validation.Required,
			validation.By(func(value interface{}) error {
				validAuthTypes := []frontoffice.AuthType{
					frontoffice.AuthType_AUTH_TYPE_GOOGLE,
					frontoffice.AuthType_AUTH_TYPE_APPLE,
				}

				authType, ok := value.(frontoffice.AuthType)
				if !ok {
					return errors.New("invalid auth type")
				}
				if !lo.Contains(validAuthTypes, authType) {
					return errors.New("invalid auth type")
				}

				return nil
			}),
		),
	}

	return errs.Filter()
}
