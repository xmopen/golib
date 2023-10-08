// Package timeutils 时间工具库.
package timeutils

import (
	"fmt"
	"strings"
	"time"
)

const (
	CNTimeTemplate = "%s年%s月%s日 %s:%s"
)

// StringTimeToCNTime format string time to chian template.
func StringTimeToCNTime(timeStr string) (string, error) {
	if timeStr == "" {
		return "", fmt.Errorf("time param is empty")
	}
	// 2006-01-02 15:04:05
	times := strings.SplitN(timeStr, " ", 2)
	if len(times) != 2 {
		return "", fmt.Errorf("tiem string is illegal ")
	}
	datas := strings.SplitN(times[0], "-", 3)
	if len(datas) != 3 {
		return "", fmt.Errorf("time string is illegal")
	}
	hours := strings.SplitN(times[1], ":", 3)
	if len(hours) != 3 {
		return "", fmt.Errorf("time string is illegal")
	}
	return fmt.Sprintf(CNTimeTemplate, datas[0], datas[1], datas[2], hours[0], hours[1]), nil
}

// NowTime return now time format  2006-01-02 15:04:05
func NowTime() string {
	return time.Now().Format(time.DateTime)
}
