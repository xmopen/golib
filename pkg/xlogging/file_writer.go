// Package xlogging
// Create  2023-03-11 16:50:54 by zhenxinma.
package xlogging

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

// 默认日志大小为64M.
const defaultLogFileSize = 1 * 1024 * 1024 * 64

// FileWriter 自定义FileWriter.
type FileWriter struct {
	Path string // 文件路径.
	Size int64  // 超出size大小则进行文件分割.

	logFile     *logFile
	writeBuffer *bufio.Writer
}

type logFile struct {
	file     *os.File
	fileSize int64 // 文件大小.
}

// NewLogFileWriter FileWriter构造函数.
func NewLogFileWriter(path string, size int64) *FileWriter {

	var (
		err      error
		tempPath = path
	)
	if tempPath == "" {
		if path, err = os.Getwd(); err != nil {
			panic(err)
		}
		tempPath += "/log.log"
	}
	if !strings.HasSuffix(tempPath, ".log") {
		tempPath += "log.log"
	}
	if size <= 0 {
		size = defaultLogFileSize
	}

	logFileWriter := &FileWriter{
		Path:    tempPath,
		Size:    size,
		logFile: &logFile{},
	}

	if _, err := os.Stat(logFileWriter.Path); err != nil && errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(logFileWriter.Path)
		if err != nil {
			panic(err)
		}
	}

	file, err := os.Open(logFileWriter.Path)
	if err != nil {
		panic(err)
	}
	logFileWriter.logFile.file = file
	logFileWriter.writeBuffer = bufio.NewWriter(logFileWriter)

	if fileInfo, err := os.Stat(logFileWriter.Path); err == nil {
		// 文件分割.
		if fileInfo.Size() >= logFileWriter.Size {
			tempPath := logFileWriter.Path
			if strings.HasSuffix(tempPath, ".log") {
				temp := strings.Split(tempPath, "/")
				temp = temp[0 : len(temp)-1]
				temp = append(temp, strconv.Itoa(int(time.Now().Unix())))
				tempPath = strings.Join(temp, "/")
				tempPath += ".log"
			}
			return NewLogFileWriter(tempPath, size)
		}
		// 初始化文件打开大小.
		//atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&logFileWriter.logFile.fileSize)), unsafe.Pointer(&size))
		atomic.StoreInt64(&logFileWriter.logFile.fileSize, fileInfo.Size())
	}

	return logFileWriter
}

func (f *FileWriter) getFile() *os.File {

	fileInfo, err := os.Stat(f.Path)
	if err != nil {
		panic(err)
	}
	fileInfo.Size()
	return nil
}

// reset 重置文件.
func (f *FileWriter) reset() {
	// TODO: 待扩展.
}

// Write 用于logrus写入文件.
func (f *FileWriter) Write(p []byte) (n int, err error) {

	if int64(len(p))+f.logFile.fileSize > f.Size {
		f.reset()
	}
	defer func() {
		atomic.StoreInt64(&f.logFile.fileSize, int64(len(p)))
	}()

	return f.writeBuffer.Write(p)
}
