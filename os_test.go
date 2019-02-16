package go_utils

import (
    "github.com/Danceiny/go.fastjson"
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
    v = GetEnvOrDefault("KEY", fastjson.JSONObject{"KEY": 10})
    o := v.(fastjson.JSONObject)
    vv, ok := o.GetInt("KEY")
    assert.Equal(t, true, ok)
    assert.Equal(t, 10, vv)

    _ = os.Setenv("KEY", "{\"KEY\": 10}")
    v = GetEnvOrDefault("KEY", fastjson.JSONObject{})
    o = v.(fastjson.JSONObject)
    vv, ok = o.GetInt("KEY")
    assert.Equal(t, true, ok)
    assert.Equal(t, 10, vv)

    _ = os.Setenv("KEY", "[{\"KEY\": 10}]")
    v = GetEnvOrDefault("KEY", fastjson.JSONArray{})
    o2 := v.(fastjson.JSONArray)
    o2o := o2.Get(0)
    vv, ok = o2o.GetInt("KEY")
    assert.Equal(t, true, ok)
    assert.Equal(t, 10, vv)
}
