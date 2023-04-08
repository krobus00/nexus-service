package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"

	"github.com/krobus00/auth-service/pb/auth"
	"github.com/krobus00/nexus-service/internal/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.Register) (*model.AuthResponse, error) {
	res, err := r.authClient.Register(ctx, &auth.RegisterRequest{
		FullName: input.FullName,
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	})

	e, ok := status.FromError(err)
	if !ok {
		return nil, model.ErrInternal
	}

	switch e.Code() {
	case codes.OK:
	case codes.FailedPrecondition:
		return nil, gqlerror.Errorf(e.Message())
	default:
		return nil, model.ErrInternal
	}

	return &model.AuthResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.AuthResponse, error) {
	res, err := r.authClient.Login(ctx, &auth.LoginRequest{
		Username: input.Username,
		Password: input.Password,
	})

	e, ok := status.FromError(err)
	if !ok {
		return nil, model.ErrInternal
	}

	switch e.Code() {
	case codes.OK:
	case codes.FailedPrecondition, codes.NotFound:
		return nil, gqlerror.Errorf(e.Message())
	default:
		return nil, model.ErrInternal
	}

	return &model.AuthResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context) (*model.AuthResponse, error) {
	userID := getUserIDFromCtx(ctx)
	tokenID := getTokenIDFromCtx(ctx)

	res, err := r.authClient.RefreshToken(ctx, &auth.RefreshTokenRequest{
		SessionUserId: userID,
		TokenId:       tokenID,
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

	return &model.AuthResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}

// Logout is the resolver for the logout field.
func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	userID := getUserIDFromCtx(ctx)
	tokenID := getTokenIDFromCtx(ctx)

	_, err := r.authClient.Logout(ctx, &auth.LogoutRequest{
		SessionUserId: userID,
		TokenId:       tokenID,
	})

	e, ok := status.FromError(err)
	if !ok {
		return false, model.ErrInternal
	}

	switch e.Code() {
	case codes.OK:
	case codes.FailedPrecondition, codes.NotFound, codes.Unauthenticated:
		return false, gqlerror.Errorf(e.Message())
	default:
		return false, model.ErrInternal
	}

	return true, nil
}

// UserInfo is the resolver for the userInfo field.
func (r *queryResolver) UserInfo(ctx context.Context) (*model.User, error) {
	userID := getUserIDFromCtx(ctx)

	res, err := r.authClient.GetUserInfo(ctx, &auth.GetUserInfoRequest{
		UserId: userID,
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

	return &model.User{
		ID:        res.GetId(),
		FullName:  res.GetFullName(),
		Username:  res.GetUsername(),
		Email:     res.GetEmail(),
		CreatedAt: res.GetCreatedAt(),
		UpdatedAt: res.GetUpdatedAt(),
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
