package pokecache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	tests := map[string]struct {
		key   string
		value string
	}{
		"empty key": {
			key:   "",
			value: "test_val",
		},
		"non empty value": {
			key:   "test_key",
			value: "test_val",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			cache := NewCache()
			cache.Add(tc.key, []byte(tc.value))
			res, ok := cache.Get(tc.key)
			assert.Equal(t, []byte(tc.value), res)
			assert.True(t, ok)
		})
	}
}

func TestCache_TimeoutFlush(t *testing.T) {
	cache := NewCache()
	cache.Add("test", []byte("test_value"))
	time.Sleep(refreshInterval + time.Second)
	res, ok := cache.Get("test")
	assert.Nil(t, res)
	assert.False(t, ok)
}
