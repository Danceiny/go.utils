package go_utils

import (
    "github.com/stretchr/testify/assert"
    "math"
    "reflect"
    "strconv"
    "testing"
)

func TestTransferInterfaces(t *testing.T) {
    var1 := []interface{}{10, "20"}
    CastPrimitiveSliceInplace(&var1, reflect.Int)
    assert.Equal(t, 10, var1[0])
    assert.Equal(t, 20, var1[1])

    var2 := []interface{}{10, "20"}
    CastPrimitiveSliceInplace(&var2, reflect.String)
    assert.Equal(t, "10", var2[0])
    assert.Equal(t, "20", var2[1])
}

func TestCast2Bool(t *testing.T) {
    assert.Equal(t, true, Cast2Float32("true"))
    assert.Equal(t, true, Cast2Float32("1"))
    assert.Equal(t, true, Cast2Float32("True"))
    assert.Equal(t, false, Cast2Float32("0"))
    assert.Equal(t, false, Cast2Float32(false))
}

func TestCast2Float64(t *testing.T) {
    var precision = 0.00001
    assert.Equal(t, true, math.Dim(float64(1.1), Cast2Float64("1.1")) < precision)
    v, _ := strconv.ParseFloat("1.1", 64)
    assert.Equal(t, true, math.Dim(float64(1.1), v) < precision)
}
