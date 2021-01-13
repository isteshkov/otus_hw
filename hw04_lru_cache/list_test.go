package hw04_lru_cache //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})

	t.Run("push elements", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Front().Next)

		l.PushFront(100) // [100, 10]
		require.Equal(t, 10, l.Front().Next.Value)
		require.Equal(t, 100, l.Front().Next.Prev.Value)

		l.Remove(l.Front())
		l.PushBack(100) // [10, 100]
		require.Equal(t, 100, l.Front().Next.Value)
		require.Equal(t, 10, l.Front().Next.Prev.Value)

		l.PushBack(200) // [10, 100, 200]
		require.Equal(t, 200, l.Back().Value)
		require.Equal(t, 200, l.Back().Prev.Next.Value)
		require.Equal(t, 100, l.Back().Prev.Value)
	})

	t.Run("move to front", func(t *testing.T) {
		l := NewList()

		// one item in list [10]
		l.PushFront(10)
		require.Equal(t, 10, l.Front().Value)
		l.MoveToFront(l.Front())
		require.Equal(t, 10, l.Front().Value)

		// two items in list [10, 100]
		l.PushBack(100)
		require.Equal(t, 100, l.Back().Value)
		l.MoveToFront(l.Back()) // [100, 10]
		require.Equal(t, 100, l.Front().Value)

		// three items and last item move to front
		l.PushBack(200) // [100, 10, 200]
		require.Equal(t, 200, l.Back().Value)
		require.Equal(t, 100, l.Front().Next.Value)
		l.MoveToFront(l.Back()) // [200, 100, 10]
		require.Equal(t, 200, l.Front().Value)
		// move middle to front
		l.MoveToFront(l.Front().Next) // [100, 200, 10]
		require.Equal(t, 100, l.Front().Value)

		// more items and middle item move to front
		l.PushBack(300) // [100, 200, 10, 300]
		require.Equal(t, 300, l.Back().Value)
		l.MoveToFront(l.Back().Prev) // [10, 100, 200, 300]
		require.Equal(t, 10, l.Front().Value)
	})

	t.Run("remove", func(t *testing.T) {
		l := NewList()

		// one item in list [10]
		l.PushFront(10)
		require.Equal(t, 10, l.Front().Value)
		l.Remove(l.Front())
		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())

		// two items in list remove last [10, 100]
		l.PushFront(10)
		l.PushBack(100)
		require.Equal(t, 100, l.Back().Value)
		l.Remove(l.Back()) // [10]
		require.Equal(t, 1, l.Len())
		require.Nil(t, l.Back())

		// three items in list remove middle [100, 10, 200]
		l.PushFront(100)
		l.PushBack(200)
		require.Equal(t, 10, l.Front().Next.Value)
		l.Remove(l.Front().Next) // [100, 200]
		require.Equal(t, 2, l.Len())

		l.PushBack(300)
		l.PushBack(400)
		require.Equal(t, 400, l.Back().Value)
		l.Remove(l.Back())
		require.Equal(t, 3, l.Len())
		require.Equal(t, 300, l.Back().Value)
		require.NotNil(t, l.Back().Prev)
		require.NotNil(t, l.Back().Prev.Next)
	})
}
