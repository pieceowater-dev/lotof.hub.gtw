package middleware

import (
	"app/internal/core/graph/model"
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
)

type contextKey string

const TokenContextKey contextKey = "token"

func AuthDirective(ctx context.Context, next graphql.Resolver, verifyTokenFunc func(string) (bool, *model.User, error)) (any, error) {
	raw := graphql.GetRequestContext(ctx).Headers.Get("Authorization")
	if raw == "" {
		return nil, errors.New("unauthorized: token is missing")
	}

	// "Bearer <token>"
	if len(raw) > 7 && raw[:7] == "Bearer " {
		raw = raw[7:]
	}

	valid, user, err := verifyTokenFunc(raw)
	if err != nil || !valid || user == nil {
		return nil, errors.New("unauthorized: token is invalid or user not found")
	}

	newCtx := context.WithValue(ctx, TokenContextKey, user)

	return next(newCtx)
}
