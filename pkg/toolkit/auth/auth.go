package auth

import (
	"context"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/metadata"
)

const (
	contextKey = "userid"
)

func WithUserId(ctx context.Context, value uuid.UUID) context.Context {
	return context.WithValue(ctx, contextKey, value)
}

func UserIdFromContext(ctx context.Context) uuid.UUID {
	md, _ := metadata.FromIncomingContext(ctx)
	if val := md[contextKey]; val != nil {
		return uuid.Must(uuid.FromString(val[0]))
	}

	if val := ctx.Value(contextKey); val != nil {
		if uuidVal, ok := val.(string); ok {
			return uuid.Must(uuid.FromString(uuidVal))
		}
	}

	return uuid.Nil
}
