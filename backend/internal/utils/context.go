package utils

import "context"

type ContextKey string

const (
	UserIDKey ContextKey = "user_id"
	EmailKey ContextKey = "email"
)

func GetUserID(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value("user_id").(int64)
	return userID, ok
}

func GetUserEmail(ctx context.Context) (string, bool) {
	email, ok := ctx.Value("email").(string)
	return email, ok
}