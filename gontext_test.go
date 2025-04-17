package gontext

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func pointerOf[T any](v T) *T {
	return &v
}

func TestValue(t *testing.T) {
	type args struct {
		ctx context.Context
		key any
	}
	tests := []struct {
		name  string
		args  args
		want  *string
		want1 bool
	}{
		{
			name: "success: returns value and true",
			args: args{
				ctx: context.WithValue(context.Background(), "key", pointerOf("value")),
				key: "key",
			},
			want:  pointerOf("value"),
			want1: true,
		},
		{
			name: "success: returns zero and false",
			args: args{
				ctx: context.Background(),
				key: "key",
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Value[*string](tt.args.ctx, tt.args.key)
			assert.Equal(t, tt.want1, got1)
			if got1 {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestValueOrDefault(t *testing.T) {
	type args struct {
		ctx          context.Context
		key          any
		defaultValue *string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "success: returns value",
			args: args{
				ctx:          context.WithValue(context.Background(), "key", pointerOf("value")),
				key:          "key",
				defaultValue: pointerOf("default"),
			},
			want: pointerOf("value"),
		},
		{
			name: "success: returns defaultValue",
			args: args{
				ctx:          context.Background(),
				key:          "key",
				defaultValue: pointerOf("default"),
			},
			want: pointerOf("default"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValueOrDefault(tt.args.ctx, tt.args.key, tt.args.defaultValue)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestValueOrZero(t *testing.T) {
	type args struct {
		ctx context.Context
		key any
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValueOrZero[*string](tt.args.ctx, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValueOrZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
