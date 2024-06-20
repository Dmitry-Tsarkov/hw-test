package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3).(*lruCache)

		c.Set("a", 1)
		c.Set("b", 2)
		c.Set("c", 3)

		assert.Equal(t, 3, c.queue.Len())
		assert.Equal(t, 3, len(c.items))

		c.Set("d", 4)

		_, found := c.Get("a")
		assert.False(t, found)

		value, found := c.Get("b")
		assert.True(t, found)
		assert.Equal(t, 2, value)

		value, found = c.Get("c")
		assert.True(t, found)
		assert.Equal(t, 3, value)

		value, found = c.Get("d")
		assert.True(t, found)
		assert.Equal(t, 4, value)
	})

	t.Run("old element", func(t *testing.T) {
		c := NewCache(3).(*lruCache)

		c.Set("a", 1)
		c.Set("b", 2)
		c.Set("c", 3)

		assert.Equal(t, 3, c.queue.Len())
		assert.Equal(t, 3, len(c.items))

		c.Set("c", 10)
		c.Set("a", 20)

		c.Set("c", 200)
		c.Set("a", 300)

		c.Set("c", 2000)
		c.Set("a", 3000)

		assert.Equal(t, 3, c.queue.Len())
		assert.Equal(t, 3, len(c.items))

		c.Set("d", 4)

		_, found := c.Get("b")
		assert.False(t, found)

		value, found := c.Get("a")
		assert.True(t, found)
		assert.Equal(t, 3000, value)

		value, found = c.Get("c")
		assert.True(t, found)
		assert.Equal(t, 2000, value)

		value, found = c.Get("d")
		assert.True(t, found)
		assert.Equal(t, 4, value)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
