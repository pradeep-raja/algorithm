// Package list implements double linked list
package list

// Item is an element of a linked list.
type Item struct {
	Value      interface{}
	next, prev *Item
}

// Prev returns previous item
func (e *Item) Prev() *Item {
	return e.prev
}

// Next returns the next item
func (e *Item) Next() *Item {
	return e.next
}

// List represents a doubly linked list.
type List struct {
	first *Item
	last  *Item
	len   int
}

// New returns new list
func New() *List {
	return &List{}
}

// First returns first item in the list
func (l *List) First() *Item {
	return l.first
}

// Last returns last item in array
func (l *List) Last() *Item {
	return l.last
}

// Length returns length
func (l *List) Length() int {
	return l.len
}

// Append new value to the list
func (l *List) Append(v interface{}) {
	item := &Item{Value: v, next: nil, prev: l.last}
	if l.last == nil {
		l.last, l.first = item, item
	} else {
		l.last.next = item
		l.last = item
	}
	l.len++
}

// Prepend value to the list
func (l *List) Prepend(v interface{}) {
	item := &Item{Value: v, next: l.first, prev: nil}
	if l.first == nil {
		l.last, l.first = item, item
	} else {
		l.first.prev = item
		l.first = item
	}
	l.len++
}

// Reverse the list O(n)
func (l *List) Reverse() {
	// Go to each item from last and swap prev and next
	reverse := l.Last()
	for i := 0; i < l.Length(); i++ {
		c := reverse.prev
		reverse.next, reverse.prev = reverse.prev, reverse.next
		reverse = c
	}
	// Make first is last and last is first
	l.first, l.last = l.last, l.first
}

// SubList creates new List using 'from' and 'to'. Result list contains both items.
// If 'from' found after 'to' the list created in reverse order.
// if -1 provided considered as a length of list
func (l *List) SubList(from, to int) *List {
	if from == -1 {
		from = l.len - 1
	}

	if to == -1 {
		to = l.len - 1
	}

	if from > l.len-1 || to > l.len-1 || from < 0 || to < 0 {
		return nil
	}

	list := New()
	if from < to {
		item := l.First()
		for i := 0; i < l.len; i++ {
			if from <= i {
				if to >= i {
					list.Append(item.Value)
				} else {
					break
				}
			}
			item = item.Next()
		}
	} else {
		item := l.Last()
		for i := l.len - 1; i >= 0; i-- {
			if from >= i {
				if to <= i {
					list.Append(item.Value)
				} else {
					break
				}
			}
			item = item.Prev()
		}
	}

	return list
}
