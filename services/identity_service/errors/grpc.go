package errors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/mohamed-karam/go-food-delivery/identity-service/response"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func ToGRPCError(err error) (error) {

	// Register error
	var registerErr response.RegisterError

	if errors.As(err, &registerErr) {

		st := status.New(
			codes.AlreadyExists,
			"registration failed",
		)

		var violations []*errdetails.BadRequest_FieldViolation

		for field, message := range registerErr.Fields {
			violations = append(
				violations,
				&errdetails.BadRequest_FieldViolation{
					Field:       field,
					Description: message,
				},
			)
		}

		detailedStatus, detailErr := st.WithDetails(
			&errdetails.BadRequest{
				FieldViolations: violations,
			},
		)

		if detailErr != nil {
			return status.Errorf(
				codes.Internal,
				"failed to create error details: %v",
                detailErr,
			)
		}

		return detailedStatus.Err()
	}

	// Invalid credentials
	var creadentialErr response.InvalidCredentialsError

	if errors.As(err, &creadentialErr) {
		return status.Errorf(
			codes.Unauthenticated,
			creadentialErr.Error(),
		)
	}

	// Validation error
	var validationErr validator.ValidationErrors

	if errors.As(err, &validationErr) {
		var violations []*errdetails.BadRequest_FieldViolation

		for _, fieldErr := range validationErr {
			violations = append(
				violations, 
				&errdetails.BadRequest_FieldViolation{
					Field: strings.ToLower(fieldErr.Field()),
					Description: validateMessage(fieldErr),
				},
			)
		}

		st := status.New(
			codes.InvalidArgument,
			"Validation failed",
		)

		detailedStatus, detailErr := st.WithDetails(
			&errdetails.BadRequest{
				FieldViolations: violations,
			},
		)

		if detailErr != nil {
			return status.Errorf(
				codes.Internal,
				"failed to create validation error details: %v",
				detailErr,
			)
		}

		return detailedStatus.Err()
	}

	// Unexpected error
	return status.Errorf(
		codes.Internal,
		"internal error: %v",
        err,
	)
}


func validateMessage(err validator.FieldError) string { 
	switch err.Tag() {
		case "required":
			return "This field is requried"

		case "email":
			return "Please enter a valid email address"

		case "min":
			return fmt.Sprintf(
				"Must be at least %s characters", 
				err.Param(),
			)

		case "max":
			return fmt.Sprintf(
				"Must not exceed %s characters", 
				err.Param(),
			)

		case "len":
			return fmt.Sprintf(
				"Must be exactly %s characters",
				err.Param(),
			)

		case "numeric":
			return "Must contain only numbers"

		default:
			return "Invalid value"
	}
}