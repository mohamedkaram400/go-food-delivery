package errors

import (
	"net/http"

	res "github.com/mohamed-karam/go-food-delivery/api-gateway-service/response"
	"github.com/gin-gonic/gin"
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

		case codes.AlreadyExists:
			c.JSON(http.StatusConflict, res.APIResponse{
				Success: false,
				Message: grpcStatus.Message(),
				Errors: extractFieldErrors(grpcStatus),
			})

		case codes.Unauthenticated:
			c.JSON(http.StatusUnauthorized, res.APIResponse{
				Success: false,
				Message: grpcStatus.Message(),
			})

		case codes.NotFound:
			c.JSON(http.StatusNotFound, res.APIResponse{
            	Success: false,
				Message: grpcStatus.Message(),
			})

		case codes.PermissionDenied:
			c.JSON(http.StatusForbidden, res.APIResponse{
				Success: false,
				Message: grpcStatus.Message(),
			})

		default:
			c.JSON(http.StatusInternalServerError, res.APIResponse{
				Success: false,
				Message: "Internal server error",
			})
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