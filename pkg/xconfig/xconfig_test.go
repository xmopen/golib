package xconfig

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	config := InitConfig()
	fmt.Println(config.Config().Get("zhenxinma.test"))
}
