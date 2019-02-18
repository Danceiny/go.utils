package go_utils

import (
    "fmt"
    "reflect"
    "strconv"
    "unsafe"
)

func TransferInterfaceLike(v unsafe.Pointer, dv interface{}) {

}

// todo: support more type
func CastPrimitiveValueInplace(item *interface{}, t reflect.Kind) {
    switch t {
    case reflect.Int:
        *item = Cast2Int(*item)
        break
    case reflect.String:
        *item = Cast2String(*item)
        break
    case reflect.Bool:
        *item = Cast2Bool(*item)
    case reflect.Float64:
        *item = Cast2Float64(*item)
    case reflect.Float32:
        *item = Cast2Float32(*item)
    }
}

func CastPrimitiveSliceInplace(items *[]interface{}, t reflect.Kind) {
    for i, item := range *items {
        CastPrimitiveValueInplace(&item, t)
        (*items)[i] = item
    }
}

func CastPrimitiveSlice(items []interface{}, t reflect.Kind) interface{} {
    CastPrimitiveSliceInplace(&items, t)
    switch t {
    case reflect.Int:
        return toInts(items)
    case reflect.String:
        return toStrings(items)
    case reflect.Bool:
        return toBools(items)
    case reflect.Int64:
        return toInt64s(items)
    case reflect.Float64:
        return toFloat64s(items)
    default:
        return items
    }
}

func toStrings(items []interface{}) []string {
    arr := make([]string, len(items))
    for i, o := range items {
        arr[i] = o.(string)
    }
    return arr
}

func toBools(items []interface{}) []bool {
    arr := make([]bool, len(items))
    for i, o := range items {
        arr[i] = o.(bool)
    }
    return arr
}

func toFloat64s(items []interface{}) []float64 {
    arr := make([]float64, len(items))
    for i, o := range items {
        arr[i] = o.(float64)
    }
    return arr
}

func toInt64s(items []interface{}) []int64 {
    arr := make([]int64, len(items))
    for i, item := range items {
        arr[i] = item.(int64)
    }
    return arr
}

func toInts(items []interface{}) []int {
    arr := make([]int, len(items))
    for i, item := range items {
        arr[i] = item.(int)
    }
    return arr
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
    case uint32:
        return int(v.(uint32))
    case uint64:
        return int(v.(uint64))
    case int8:
        return int(v.(int8))
    case int16:
        return int(v.(int16))
    case uint8:
        return int(v.(uint8))
    case uint16:
        return int(v.(uint16))
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

func Cast2Bool(v interface{}) bool {
    switch v.(type) {
    case bool:
        return v.(bool)
    case string:
        switch v.(string) {
        case "1", "t", "T", "true", "TRUE", "True":
            return true
        case "0", "f", "F", "false", "FALSE", "False":
            return false
        }
    case int, uint, int64, uint64, float64, float32, int8, uint8, int16, uint16, int32, uint32:
        return Cast2Int(v) == 0
    default:
        return true
    }
    return false
}

func Cast2Float64(v interface{}) float64 {
    switch v.(type) {
    case float64:
        return v.(float64)
    case float32:
        return float64(v.(float32))
    case string:
        ret, err := strconv.ParseFloat(v.(string), 64)
        if err == nil {
            return ret
        }
    case int, uint, int64, uint64, int8, uint8, int16, uint16, int32, uint32:
        return float64(Cast2Int(v))
    }
    return 0
}

func Cast2Float32(v interface{}) float32 {
    switch v.(type) {
    case float64:
        return float32(v.(float64))
    case float32:
        return v.(float32)
    case string:
        ret, err := strconv.ParseFloat(v.(string), 32)
        if err == nil {
            return float32(ret)
        }

    case int, uint, int64, uint64, int8, uint8, int16, uint16, int32, uint32:
        return float32(Cast2Int(v))
    }
    return 0
}
