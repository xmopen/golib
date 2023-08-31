package commonutil

import (
	"crypto/md5"
	"fmt"
)

func MD5(data string) string {
	has := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", has)
}
