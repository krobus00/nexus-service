package graph

import (
	"context"
	"errors"

	"github.com/krobus00/nexus-service/internal/constant"
)

func getUserIDFromCtx(ctx context.Context) (string, error) {
	userID, ok := ctx.Value(constant.KeyUserIDCtx).(string)
	if !ok || userID == "" {
		return "", errors.New("invalid user id")
	}

	return userID, nil
}
