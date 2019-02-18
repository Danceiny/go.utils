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

    TransferInterfaces(&var1, reflect.String)
    assert.Equal(t, "10", var1[0])
    assert.Equal(t, "20", var1[1])
}
