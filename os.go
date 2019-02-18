package go_utils

import (
    "fmt"
    "os"
    "reflect"
    "strconv"
)

// k, env name for `os.LookupEnv` call
// dv, for default return value. some advices(or say notices) may be helpful:
// 1. if do not care default value, you can just pass `nil`, and the function's behavior is just like `os.GetEnv`
// 2. if you do known or do need that the type of returned value is specified,
// you can pass the `zero-value` of this type, like "" for string, 0 for int
// 3. if dv is not `nil`, and the env variable found, will cast the env variable from string type to dv's type
func GetEnvOrDefault(k string, dv interface{}) interface{} {
    v, ok := os.LookupEnv(k)
    if !ok {
        return dv
    } else {
        if dv == nil {
            return v
        }
        var err error
        switch dv.(type) {
        case string:
            return v
        case bool:
            var ret bool
            ret, err = strconv.ParseBool(v)
            if err == nil {
                return ret
            }
            break
        case int:
            var ret int
            ret, err = strconv.Atoi(v)
            if err == nil {
                return ret
            }
            break
        case int64:
            var ret int64
            ret, err = strconv.ParseInt(v, 10, 64)
            if err == nil {
                return ret
            }
            break
        case int32:
            var ret int64
            ret, err = strconv.ParseInt(v, 10, 32)
            if err == nil {
                return ret
            }
            break
        case float64:
            var ret float64
            ret, err = strconv.ParseFloat(v, 64)
            if err == nil {
                return ret
            }
            break
        case float32:
            var ret float64
            ret, err = strconv.ParseFloat(v, 32)
            if err == nil {
                return ret
            }
            break
        default:
            fmt.Printf("type: %v", reflect.TypeOf(dv))
            return v
        }
        if err != nil {
            // TODO: 此处返回值可商榷
            return dv
        }
        return nil
    }
}
