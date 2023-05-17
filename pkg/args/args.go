package args

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// FormatArgs2Struct 解析Args为结构体,参数必须是指针.
func FormatArgs2Struct(data interface{}) error {

	argsDataMap := parseArgs2Map()

	vf := reflect.ValueOf(data)
	tf := reflect.TypeOf(data)
	if vf.Kind() != reflect.Ptr {
		return fmt.Errorf("FormatArgs2Struct args not is pointer")
	}
	vf = vf.Elem()
	tf = tf.Elem()
	for i := 0; i < vf.NumField(); i++ {
		field := tf.Field(i)
		name := field.Name
		nameValue, ok := argsDataMap[name]
		if !ok {
			continue
		}

		vf.Field(i).Set(reflect.ValueOf(nameValue))
	}
	return nil
}

// parseArgs2Map 解析os.Args为Map.
func parseArgs2Map() map[string]interface{} {

	dataMap := make(map[string]interface{})
	osArgs := os.Args[1:]
	index := 0
	for {

		key := osArgs[index]
		// 不支持仅仅只有一个K的参数.
		if !strings.HasPrefix(key, "-") {
			index++
			continue
		}

		index++
		value := osArgs[index]
		dataMap[key[1:]] = value
		index++
		if index >= len(osArgs) {
			break
		}
	}

	data, _ := json.Marshal(dataMap)
	fmt.Printf("dataMap:%s", string(data))
	return dataMap
}
