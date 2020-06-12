package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *ListItem                  // первый Item
	Back() *ListItem                   // последний Item
	PushFront(v interface{}) *ListItem // добавить значение в начало
	PushBack(v interface{}) *ListItem  // добавить значение в конец
	Remove(i *ListItem)                // удалить элемент
	MoveToFront(i *ListItem)           // поместить элемент в начало
	PrintList()
}

type ListItem struct {
	Value interface{} // значение
	Next  *ListItem   // следующий элемент
	Prev  *ListItem   // предыдущий элемент
}

type list struct {
	Size  int       // Размер
	First *ListItem // первый элемент
	Last  *ListItem // последний элемент
}

func NewList() List {
	return new(list)
}

func (lst list) Len() int {
	return lst.Size
}

func (lst *list) Front() *ListItem {
	return lst.First
}

func (lst *list) Back() *ListItem {
	return lst.Last
}

func (lst *list) PushFront(v interface{}) *ListItem {
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

func (lst *list) PushBack(v interface{}) *ListItem {
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

func (lst *list) Remove(i *ListItem) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	lst.Size--
	i = nil
}

func (lst *list) MoveToFront(i *ListItem) {
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

func (lst *list) PrintList() {
	itemPtr := lst.First
	for j := 0; j < lst.Size; j++ {
		itemPtr = itemPtr.Prev
	}
}
