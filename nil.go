package go_utils

import "reflect"

func InterfaceHasNilValue(actual interface{}) bool {
    value := reflect.ValueOf(actual)
    kind := value.Kind()
    nilable := kind == reflect.Ptr ||
        kind == reflect.Slice ||
        kind == reflect.Chan ||
        kind == reflect.Map ||
        kind == reflect.Func

    // Careful: reflect.Value.IsNil() will panic unless it's an interface, chan, map, func, slice, or ptr
    // Reference: http://golang.org/pkg/reflect/#Value.IsNil
    return nilable && value.IsNil()
}
