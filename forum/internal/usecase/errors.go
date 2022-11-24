package usecase

import (
	v1 "github.com/Skijetler/alphinium/forum/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrCategoryNotFound  = status.Errorf(codes.NotFound, "reason: %v", v1.ErrorReason_CATEGORY_NOT_FOUND.String())
	ErrUserNotAuthorized = status.Errorf(codes.Unauthenticated, "reason: %v", v1.ErrorReason_NOT_AUTHORIZED.String())
)

func internalErr(reason error) error {
	return status.Errorf(codes.Internal, "reason: %v", reason.Error())
}
