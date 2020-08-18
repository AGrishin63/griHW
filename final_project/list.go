package main

import "time"

type ListItem struct {
	// Place your code here
	Value []time.Time // значение
	Next  *ListItem   // следующий элемент
	Prev  *ListItem   // предыдущий элемент
}

type List struct {
	// Place your code here
	Size  int       // Размер
	First *ListItem // первый элемент
	Last  *ListItem // последний элемент
}

func NewList() *List {
	return &List{
		Size:  0,
		First: nil,
		Last:  nil,
	}
}

func (lst List) Len() int {
	return lst.Size
}

func (lst *List) Front() *ListItem {
	return lst.First
}

func (lst *List) Back() *ListItem {
	return lst.Last
}

func (lst *List) PushFront(v []time.Time) *ListItem {
	it := &ListItem{
		v,
		nil,
		nil,
	}
	if lst.Len() != 0 {
		it.Prev = lst.Front()
		lst.First.Next = it
	} else {
		lst.Last = it
	}
	lst.First = it
	lst.Size++

	return it
}

func (lst *List) PushBack(v []time.Time) *ListItem {
	it := &ListItem{
		v,
		nil,
		nil,
	}
	if lst.Len() != 0 {
		it.Next = lst.Last
		lst.Last.Prev = it
	} else {
		lst.First = it
	}
	lst.Last = it
	lst.Size++

	return it
}

func (lst *List) Remove(i *ListItem) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	lst.Size--
	i = nil
}

func (lst *List) MoveToFront(i *ListItem) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		return
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		lst.Last = i.Next
	}
	i.Next = nil
	i.Prev = lst.First
	lst.First.Next = i
	lst.First = i
}

func (lst *List) PrintList() {
	itemPtr := lst.First
	for j := 0; j < lst.Size; j++ {
		itemPtr = itemPtr.Prev
	}
}
