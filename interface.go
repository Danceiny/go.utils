package go_utils

import (
	"fmt"
	"reflect"
	"strconv"
)

// todo: support more type
func TransferInterfaces(items *[]interface{}, t reflect.Kind) {
	var v interface{}
	for i, item := range *items {
		switch t {
		case reflect.Int:
			v = Cast2Int(item)
			break
		case reflect.String:
			v = Cast2String(item)
			break
		}
		(*items)[i] = v
	}
}

func Cast2Int(v interface{}) int {
	var ret int
	var err error
	switch v.(type) {
	case int:
		return v.(int)
	case int64:
		return int(v.(int64))
	case int32:
		return int(v.(int32))
	default:
		ret, err = strconv.Atoi(fmt.Sprint(v))
		if err == nil {
			return ret
		}
	}
	if err != nil {
		return 0
	} else {
		return ret
	}
}

func Cast2String(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	default:
		return fmt.Sprint(v)
	}
}
