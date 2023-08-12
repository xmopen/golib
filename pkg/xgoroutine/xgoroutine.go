// Package xgoroutine 封装安全goroutine.
package xgoroutine

import (
	"runtime/debug"

	"github.com/xmopen/golib/pkg/xlogging"
)

// SafeGoroutine recover goroutine.
func SafeGoroutine(fn func(), xlogs ...*xlogging.Entry) {
	var xlog *xlogging.Entry
	if len(xlogs) == 0 {
		xlog = xlogging.Tag("safe.goroutine")
	} else {
		xlog = xlogs[0]
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				// 这里没有打印出来error日志.
				xlog.Errorf("panic stack:[%+v]", string(debug.Stack()))
			}
		}()
		fn()
	}()
}
