package xlogging

import (
	"testing"
)

func TestTag(t *testing.T) {
	xlog := Tag("test_tag")
	xlog.Infof("test_tag")
}
