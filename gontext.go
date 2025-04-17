package gontext

import "context"

func Value[T any](ctx context.Context, key any) (T, bool) {
	value := ctx.Value(key)
	if v, ok := value.(T); ok {
		return v, true
	}

	return *new(T), false
}

func ValueOrDefault[T any](ctx context.Context, key any, defaultValue T) T {
	if v, ok := Value[T](ctx, key); ok {
		return v
	}

	return defaultValue
}

func ValueOrZero[T any](ctx context.Context, key any) T {
	v, _ := Value[T](ctx, key)

	return v
}
