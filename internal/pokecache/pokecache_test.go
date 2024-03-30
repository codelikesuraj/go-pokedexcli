package pokecache

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCases []struct {
	inputKey string
	inputVal []byte
}

func TestCreateCache(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)

	assert.NotNil(t, cache.cache)
}

func TestAddCache(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)

	require.NotNil(t, cache.cache)

	testCases := testCases{
		{"key1", []byte("val1")},
		{"key2", []byte("val2")},
	}

	for _, testCase := range testCases {
		cache.Add(testCase.inputKey, testCase.inputVal)
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			assert.Contains(t, cache.cache, testCase.inputKey)
		})
	}
}

func TestGetCache(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)

	require.NotNil(t, cache.cache)

	testCases := testCases{
		{"key1", []byte("val1")},
		{"key2", []byte("val2")},
	}

	for _, testCase := range testCases {
		cache.Add(testCase.inputKey, testCase.inputVal)
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			val, ok := cache.Get(testCase.inputKey)

			assert.True(t, ok)
			assert.Equal(t, val, testCase.inputVal)
		})
	}
}

func TestReap(t *testing.T) {
	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	testCases := testCases{
		{"key1", []byte("val1")},
		{"key2", []byte("val2")},
	}

	for _, testCase := range testCases {
		cache.Add(testCase.inputKey, testCase.inputVal)
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case %v before deadline", i), func(t *testing.T) {
			_, ok := cache.Get(testCase.inputKey)

			assert.True(t, ok, "Should not have been reaped")
		})
	}

	time.Sleep(interval + (100 * time.Millisecond))

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case %v after deadline", i), func(t *testing.T) {
			_, ok := cache.Get(testCase.inputKey)

			assert.False(t, ok, "Should have been reaped")
		})
	}
}
