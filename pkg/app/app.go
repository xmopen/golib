// Package app
// Create  2023-03-11 18:53:42 by zhenxinma.
package app

import (
	"runtime/debug"

	"gitee.com/zhenxinma/golib/pkg/utils"
	"gitee.com/zhenxinma/golib/pkg/xlogging"
	"github.com/gin-gonic/gin"
)

const (
	GinContextXlogKey = "xlog"
	GinContextReqID   = "reqid"
)

// PanicRecover 捕获panic.
func PanicRecover(xlog *xlogging.Entry) {
	if err := recover(); err != nil {
		xlog.Errorf("app panic stack:%+v", string(debug.Stack()))
	}
}

// Log 获取请求上下文log.
func Log(c *gin.Context) *xlogging.Entry {

	xlog, ok := c.Get(GinContextXlogKey)
	if ok {
		return xlog.(*xlogging.Entry)
	}
	reqid, ok := c.Get(GinContextReqID)
	if !ok {
		reqid = utils.UUID()
		c.Set(GinContextReqID, reqid)
	}
	log := xlogging.Tag(reqid.(string))
	c.Set(GinContextXlogKey, log)
	return log
}
