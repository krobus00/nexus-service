package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"

	"github.com/krobus00/nexus-service/internal/graph/model"
	"github.com/krobus00/storage-service/pb/storage"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FindObjectByID is the resolver for the findObjectByID field.
func (r *queryResolver) FindObjectByID(ctx context.Context, id string) (*model.Object, error) {
	userID := getUserIDFromCtx(ctx)

	object, err := r.storageClient.GetObjectByID(ctx, &storage.GetObjectByIDRequest{
		UserId:   userID,
		ObjectId: id,
	})

	e, ok := status.FromError(err)
	if !ok {
		return nil, model.ErrInternal
	}

	switch e.Code() {
	case codes.OK:
	case codes.FailedPrecondition, codes.NotFound, codes.Unauthenticated:
		return nil, gqlerror.Errorf(e.Message())
	default:
		return nil, model.ErrInternal
	}

	return &model.Object{
		ID:         object.GetId(),
		FileName:   object.GetFileName(),
		Type:       object.GetType(),
		SignedURL:  object.GetSignedUrl(),
		ExpiredAt:  object.GetExpiredAt(),
		IsPublic:   object.GetIsPublic(),
		UploadedBy: object.GetUploadedBy(),
		CreatedAt:  object.GetCreatedAt(),
	}, nil
}
