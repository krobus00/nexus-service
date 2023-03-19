//go:generate go run github.com/99designs/gqlgen generate
package graph

import (
	"errors"

	authPB "github.com/krobus00/auth-service/pb/auth"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authClient authPB.AuthServiceClient
}

func NewResolver() *Resolver {
	return new(Resolver)
}

func (r *Resolver) InjectAuthClient(client authPB.AuthServiceClient) error {
	if client == nil {
		return errors.New("invalid auth client")
	}
	r.authClient = client
	return nil
}
