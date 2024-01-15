// Package xlogging
// Create  2023-03-11 16:08:50 by zhenxinma.
package xlogging

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	FIELDLOGTAG = "tag"
)

type Level = logrus.Level
type Logger = logrus.Logger

var globalDefaultLogger *Logger

type Entry = logrus.Entry

// LogFormat 日志格式化.
type LogFormat struct {
	TimestampFormat string
	PrintCaller     bool
}

func init() {
	globalDefaultLogger = logrus.New()
	globalDefaultLogger.SetLevel(logrus.TraceLevel)
	// 添加钩子函数,暂时保留.
	//out := make([]io.Writer, 0)
	//out = append(out, os.Stdout)
	//if dir, err := os.Getwd(); err == nil {
	//	out = append(out, NewLogFileWriter(dir+"/var/", defaultLogFileSize))
	//} else {
	//	fmt.Printf("xlogging init pwd err:%+v\n", err)
	//}

	//globalDefaultLogger.SetOutput(io.MultiWriter(out...))
	formatter := &LogFormat{
		TimestampFormat: "2006-01-02 15:04:05.999",
		PrintCaller:     true,
	}
	globalDefaultLogger.SetFormatter(formatter)

	globalDefaultLogger.SetOutput(os.Stdout)
	globalDefaultLogger.SetReportCaller(true)
}

func (l *LogFormat) Format(entry *logrus.Entry) ([]byte, error) {

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	b.WriteString(entry.Time.Format(l.TimestampFormat))
	if tag, ok := entry.Data[FIELDLOGTAG]; ok {
		b.WriteString(" [")
		b.WriteString(tag.(string))
		b.WriteString("] ")
	}
	b.WriteByte('[')
	b.WriteString(strings.ToUpper(entry.Level.String()))
	b.WriteString("] ")
	if entry.Message != "" {
		b.WriteString(entry.Message)
	}
	// 文件名以及行数.
	if l.PrintCaller {
		b.WriteString(fmt.Sprintf(" [%s:%d]", entry.Caller.File, entry.Caller.Line))
	}
	b.WriteByte('\n')
	return b.Bytes(), nil
}

// Tag 对外提供.
func Tag(tag string) *Entry {
	return globalDefaultLogger.WithFields(logrus.Fields{FIELDLOGTAG: tag})
}
