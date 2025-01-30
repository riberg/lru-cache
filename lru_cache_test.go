package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	t.Run("Add and Get", func(t *testing.T) {
		cache := NewLRUCache(2).(*LRUCacheImpl)

		assert.True(t, cache.Add("1", "one"))
		value, ok := cache.Get("1")
		assert.True(t, ok)
		assert.Equal(t, "one", value)

		assert.False(t, cache.Add("1", "new")) // Добавляем существующий ключ
	})

	t.Run("Remove", func(t *testing.T) {
		cache := NewLRUCache(2).(*LRUCacheImpl)

		cache.Add("1", "one")

		assert.True(t, cache.Remove("1"))
		assert.False(t, cache.Remove("none"))
	})

	t.Run("Capacity overflow", func(t *testing.T) {
		cache := NewLRUCache(2).(*LRUCacheImpl)

		cache.Add("1", "one")
		cache.Add("2", "two")
		cache.Add("3", "three") // Должен вытеснить "1"

		assert.Equal(t, 2, cache.list.Len())
		_, ok := cache.Get("1")
		assert.False(t, ok)

		cache.Get("2")         // Повышаем приоритет
		cache.Add("4", "four") // Должен вытеснить "3"
		_, ok = cache.Get("3")
		assert.False(t, ok)
	})
}
