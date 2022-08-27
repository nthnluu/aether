package aether

import (
	"context"
	"testing"
)

func TestGetFullMethodNameFromContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), FullMethodNameCtxKey, "asdf")
	methodName := GetFullMethodNameFromContext(ctx)

	if methodName != "asdf" {
		t.Errorf("expected method name 'asdf', got '%v'\n", methodName)
	}

	defer func() {
		if err := recover(); err != nil {
			t.Errorf("GetFullMethodNameFromContext resulted in panic %v", err)
		}
	}()
}
