package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
	Key   Key
}

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}
	prev := i.Prev
	next := i.Next
	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	} else {
		l.back = prev
	}
	first := l.front
	first.Prev = i
	i.Prev = nil
	i.Next = first
	l.front = i
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := new(ListItem)
	newItem.Value = v
	if l.len == 0 {
		l.back, l.front = newItem, newItem
	} else {
		newItem.Next = l.front
		l.front.Prev = newItem
		l.front = newItem
	}
	l.len++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}
	if l.len == 0 {
		l.back, l.front = newItem, newItem
	} else {
		newItem.Prev = l.back
		l.back.Next = newItem
		l.back = newItem
	}
	l.len++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = l.front.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = l.back.Prev
	}
	i.Next = nil
	i.Prev = nil
	l.len--
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func NewList() List {
	return new(list)
}
