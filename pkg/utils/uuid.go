// Package utils
// Create  2023-03-11 19:04:50 by zhenxinma.
package utils

import (
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// UUID 获取随机UUID.
func UUID() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return strconv.Itoa(int(time.Now().Unix()))
	}
	return strings.ReplaceAll(uuid.String(), "-", "")
}
