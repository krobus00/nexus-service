package constant

type ctxKey string

const (
	KeyUserIDCtx  ctxKey = "USERID"
	KeyTokenIDCtx ctxKey = "TOKENID"

	GuestID = string("GUEST")
)
