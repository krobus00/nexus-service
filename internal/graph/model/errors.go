package model

import "github.com/vektah/gqlparser/v2/gqlerror"

var (
	ErrInternal = gqlerror.Errorf("internal server error")
)
