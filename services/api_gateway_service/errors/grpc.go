package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	res "github.com/mohamed-karam/go-food-delivery/api-gateway-service/response"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WriteGRPCError(err error, c *gin.Context) {

	grpcStatus, ok := status.FromError(err)

	if !ok {
		c.JSON(http.StatusInternalServerError, res.APIResponse{
            Success: false,
            Message: "Internal server error",
        })
		return
	}

	switch grpcStatus.Code() {
		case codes.InvalidArgument:
			c.JSON(http.StatusBadRequest, res.APIResponse{
				Success: false,
				Message: grpcStatus.Message(),
				Errors: extractFieldErrors(grpcStatus),
			})
			return

		case codes.AlreadyExists:
			c.JSON(http.StatusConflict, res.APIResponse{
				Success: false,
				Message: grpcStatus.Message(),
				Errors: extractFieldErrors(grpcStatus),
			})
			return

		case codes.Unauthenticated:
			c.JSON(http.StatusUnauthorized, res.APIResponse{
				Success: false,
				Message: grpcStatus.Message(),
			})
			return

		case codes.NotFound:
			c.JSON(http.StatusNotFound, res.APIResponse{
            	Success: false,
				Message: grpcStatus.Message(),
			})
			return

		case codes.PermissionDenied:
			c.JSON(http.StatusForbidden, res.APIResponse{
				Success: false,
				Message: grpcStatus.Message(),
			})
			return

		case codes.Unavailable:
			c.JSON(http.StatusServiceUnavailable, res.APIResponse{
				Success: false,
				Message: "Identity service is currently unavailable. Please try again later.",
			})
			return

		default:
			c.JSON(http.StatusInternalServerError, res.APIResponse{
				Success: false,
				Message: grpcStatus.Message(),
			})
			return
	}
}

func extractFieldErrors(
    grpcStatus *status.Status,
) map[string]string {

    errorsMap := make(map[string]string)

    for _, detail := range grpcStatus.Details() {

        badRequest, ok := detail.(*errdetails.BadRequest)

        if !ok {
            continue
        }

        for _, violation := range badRequest.FieldViolations {
            errorsMap[violation.Field] = violation.Description
        }
    }

    return errorsMap
}