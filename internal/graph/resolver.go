//go:generate go run github.com/99designs/gqlgen generate
package graph

import (
	"errors"

	authPB "github.com/krobus00/auth-service/pb/auth"
	productPB "github.com/krobus00/product-service/pb/product"
	storagePB "github.com/krobus00/storage-service/pb/storage"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authClient    authPB.AuthServiceClient
	storageClient storagePB.StorageServiceClient
	productClient productPB.ProductServiceClient
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

func (r *Resolver) InjectStorageClient(client storagePB.StorageServiceClient) error {
	if client == nil {
		return errors.New("invalid storage client")
	}
	r.storageClient = client
	return nil
}

func (r *Resolver) InjectProductClient(client productPB.ProductServiceClient) error {
	if client == nil {
		return errors.New("invalid product client")
	}
	r.productClient = client
	return nil
}
