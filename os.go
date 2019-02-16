package go_utils

import "os"

func GetEnvOrDefault(k string, dv interface{}) interface{} {
    v, ok := os.LookupEnv(k)
    if !ok {
        return dv
    } else {
        return v
    }
}
