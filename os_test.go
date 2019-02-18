package go_utils

import (
    "github.com/stretchr/testify/assert"
    "os"
    "testing"
)

func TestGetEnvOrDefault(t *testing.T) {
    var v interface{}
    _ = os.Setenv("KEY", "10")

    v = GetEnvOrDefault("KEY", 0)
    assert.Equal(t, 10, v)
    _ = os.Unsetenv("KEY")
    v = GetEnvOrDefault("KEY", 100)
    assert.Equal(t, 100, v)
    v = GetEnvOrDefault("KEY", 0)
    assert.Equal(t, 0, v)
    v = GetEnvOrDefault("KEY", "")
    assert.Equal(t, "", v)

}
