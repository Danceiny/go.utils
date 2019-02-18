package go_utils

import (
    "github.com/stretchr/testify/assert"
    "reflect"
    "testing"
)

func TestTransferInterfaces(t *testing.T) {
    var1 := []interface{}{10, "20"}
    TransferInterfaces(&var1, reflect.Int)
    assert.Equal(t, 10, var1[0])
    assert.Equal(t, 20, var1[1])

    var2 := []interface{}{10, "20"}
    TransferInterfaces(&var2, reflect.String)
    assert.Equal(t, "10", var2[0])
    assert.Equal(t, "20", var2[1])
}
