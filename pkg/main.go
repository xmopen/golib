package main

import (
	"fmt"
	"reflect"

	"gitee.com/zhenxinma/golib/pkg/xlogging"
)

type Student struct {
	Age  string
	Name string
}

func Format(data interface{}) error {

	vf := reflect.ValueOf(data).Elem()
	tf := reflect.TypeOf(data).Elem()
	for i := 0; i < tf.NumField(); i++ {
		field := tf.Field(i)
		if field.Name == "Age" {
			f := vf.Field(i)
			f.SetInt(100)
		}
	}

	return nil
}

type Obj struct {
	name string
}

func (o *Obj) String() string {
	return fmt.Sprintf("{%s:%s}", "name", o.name)
}

func main() {
	xlogging.Tag("test tag").Infof("test:%v", "a")
}
