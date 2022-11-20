package usecase

import (
	v1 "github.com/Skijetler/alphinium/auth/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserNotFound      = status.Errorf(codes.NotFound, "reason: %v", v1.ErrorReason_USER_NOT_FOUND.String())
	ErrUserAlreadyExists = status.Errorf(codes.AlreadyExists, "reason: %v", v1.ErrorReason_USER_ALREADY_EXISTS.String())
	ErrUserIsDisabled    = status.Errorf(codes.PermissionDenied, "reason: %v", v1.ErrorReason_USER_IS_DISABLED.String())
	ErrWrongPassword     = status.Errorf(codes.PermissionDenied, "reason: %v", v1.ErrorReason_WRONG_PASSWORD.String())
	ErrInvalidToken      = status.Errorf(codes.Unauthenticated, "reason: %v", v1.ErrorReason_INVALID_TOKEN.String())
	ErrInvalidSession    = status.Errorf(codes.Unauthenticated, "reason: %v", v1.ErrorReason_INVALID_SESSION.String())
)

func internalErr(reason error) error {
	return status.Errorf(codes.Internal, "reason: %v", reason.Error())
}
