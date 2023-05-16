package grpc

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"github.com/vladazn/go-boilerplate/app/service/userService"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func buildError(err error) error {
	switch errors.Cause(err).(type) {
	case validation.Errors:
		return status.New(codes.InvalidArgument, "invalid arguments").Err()
	case *userService.NotFoundError:
		return status.New(codes.FailedPrecondition, err.Error()).Err()
	default:
		return status.New(codes.Internal, err.Error()).Err()
	}
}
