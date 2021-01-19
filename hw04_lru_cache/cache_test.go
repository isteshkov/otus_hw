package hw04_lru_cache //nolint:golint,stylecheck
import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c, err := NewCache(10)
		require.NoError(t, err)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c, err := NewCache(5)
		require.NoError(t, err)

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
		t.Run("out of capacity", func(t *testing.T) {
			c, err := NewCache(3)
			require.NoError(t, err)

			var key1, key2, key3, key4, key5 Key
			key1, key2, key3, key4, key5 = "aaa", "bbb", "ccc", "ddd", "eee"

			wasInCache := c.Set(key1, 1)
			require.False(t, wasInCache)

			wasInCache = c.Set(key2, 2)
			require.False(t, wasInCache)

			wasInCache = c.Set(key3, 3)
			require.False(t, wasInCache)

			v, ok := c.Get(key1)
			require.True(t, ok)
			require.Equal(t, 1, v)
			v, ok = c.Get(key2)
			require.True(t, ok)
			require.Equal(t, 2, v)
			v, ok = c.Get(key3)
			require.True(t, ok)
			require.Equal(t, 3, v)

			wasInCache = c.Set(key4, 4)
			require.False(t, wasInCache)
			_, ok = c.Get(key1)
			require.False(t, ok)

			wasInCache = c.Set(key5, 5)
			require.False(t, wasInCache)
			_, ok = c.Get(key2)
			require.False(t, ok)
		})
		t.Run("rarely used", func(t *testing.T) {
			c, err := NewCache(3)
			require.NoError(t, err)

			var key1, key2, rareKey3, key4 Key
			key1, key2, rareKey3, key4 = "aaa", "bbb", "ccc", "ddd"
			val1, val2, val3, val4 := 1, 2, 3, 4

			wasInCache := c.Set(key1, val1)
			require.False(t, wasInCache)

			wasInCache = c.Set(key2, val2)
			require.False(t, wasInCache)

			wasInCache = c.Set(rareKey3, val3)
			require.False(t, wasInCache)

			v, ok := c.Get(key1)
			require.True(t, ok)
			require.Equal(t, val1, v)

			v, ok = c.Get(rareKey3)
			require.True(t, ok)
			require.Equal(t, val3, v)

			v, ok = c.Get(key2)
			require.True(t, ok)
			require.Equal(t, val2, v)

			newVal1 := 77
			wasInCache = c.Set(key1, newVal1)
			require.True(t, wasInCache)

			newVal2 := 88
			wasInCache = c.Set(key2, newVal2)
			require.True(t, wasInCache)

			// key1 used 2 times
			// key2 used 2 times
			// key3 used 1 times

			wasInCache = c.Set(key4, val4)
			require.False(t, wasInCache)
			// key3 pushed out from cache
			_, ok = c.Get(rareKey3)
			require.False(t, ok)
		})
	})
}

func TestCacheMultithreading(t *testing.T) {
	c, err := NewCache(10)
	require.NoError(t, err)
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
