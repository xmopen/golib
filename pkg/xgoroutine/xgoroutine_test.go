package xgoroutine

import (
	"testing"
)

func TestGORecover(t *testing.T) {
	SafeGoroutine(func() {
		panic("zhenxinma panic")
	})
}
