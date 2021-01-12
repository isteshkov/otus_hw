package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(i *listItem)
	MoveToFront(i *listItem)
}

type listItem struct {
	Value interface{}
	Next  *listItem
	Prev  *listItem
}

type list struct {
	firstItem *listItem
	lastItem  *listItem
	length    int
}

func NewList() List {
	return &list{}
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *listItem {
	return l.firstItem
}

func (l *list) Back() *listItem {
	return l.lastItem
}

func (l *list) PushFront(v interface{}) *listItem {
	defer func() { l.length++ }()

	// list is empty
	if l.firstItem == nil {
		l.firstItem = &listItem{
			Value: v,
			Next:  nil,
			Prev:  nil,
		}

		return l.firstItem
	}

	// one item in list
	if l.lastItem == nil {
		l.lastItem = l.firstItem
		l.firstItem = &listItem{
			Value: v,
			Next:  l.lastItem,
			Prev:  nil,
		}
		l.lastItem.Prev = l.firstItem
		return l.firstItem
	}

	newNode := &listItem{
		Value: v,
		Next:  l.firstItem,
		Prev:  nil,
	}

	l.firstItem.Prev = newNode
	l.firstItem = newNode

	return l.firstItem
}

func (l *list) PushBack(v interface{}) *listItem {
	defer func() { l.length++ }()

	// list is empty
	if l.firstItem == nil {
		l.firstItem = &listItem{
			Value: v,
			Next:  nil,
			Prev:  nil,
		}

		return l.firstItem
	}

	// one item in list
	if l.lastItem == nil {
		l.lastItem = &listItem{
			Value: v,
			Next:  nil,
			Prev:  l.firstItem,
		}
		l.firstItem.Next = l.lastItem

		return l.lastItem
	}

	newNode := &listItem{
		Value: v,
		Next:  nil,
		Prev:  l.lastItem,
	}

	l.lastItem.Next = newNode
	l.lastItem = newNode

	return l.lastItem
}

func (l *list) Remove(i *listItem) {
	defer func() { l.length-- }()

	// removing first element
	if l.firstItem == i {
		if l.length == 1 {
			l.firstItem = nil
			return
		}
		l.firstItem = l.firstItem.Next
		l.firstItem.Prev = nil
		return
	}

	// removing last element
	if l.lastItem == i {
		if l.length == 2 {
			l.lastItem = nil
			l.firstItem.Next = nil
			return
		}
		l.lastItem = l.lastItem.Prev
		l.lastItem.Next = nil
		return
	}

	// removing other
	i.Prev.Next, i.Next.Prev = i.Next, i.Prev
}

func (l *list) MoveToFront(i *listItem) {
	// if one item in list or try to move first item
	if l.length < 2 || i == l.firstItem {
		return
	}

	// if more than two items in list (3 and more)
	if l.length > 2 {
		if i == l.lastItem {
			l.firstItem.Prev, l.lastItem.Next = l.lastItem, l.firstItem
			l.firstItem = l.lastItem
			l.lastItem = l.lastItem.Prev
			l.lastItem.Next = nil
			return
		}
		// fmt.Printf("move i: %+v -> prev(%+v) [%+v] next(%+v)", i, i.Prev, i.Value, i.Next)
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
		i.Prev = nil
		i.Next, l.firstItem.Prev = l.firstItem, i
		l.firstItem = i
		return
	}

	l.swapEdges()
}

func (l *list) swapEdges() {
	l.firstItem.Value, l.lastItem.Value = l.lastItem.Value, l.firstItem.Value
	l.firstItem.Prev, l.firstItem.Next = l.lastItem.Next, l.lastItem.Prev
	l.lastItem.Next, l.lastItem.Prev = l.firstItem.Prev, l.firstItem.Next
}
