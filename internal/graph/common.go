package graph

import (
	"context"
	"fmt"

	"github.com/krobus00/nexus-service/internal/constant"
)

func getUserIDFromCtx(ctx context.Context) string {
	ctxUserID := ctx.Value(constant.KeyUserIDCtx)

	userID := fmt.Sprintf("%v", ctxUserID)
	if userID == "" {
		return constant.GuestID
	}
	return userID
}

func getTokenIDFromCtx(ctx context.Context) string {
	ctxTokenID := ctx.Value(constant.KeyTokenIDCtx)

	tokenID := fmt.Sprintf("%v", ctxTokenID)
	return tokenID
}
