package hw04_lru_cache //nolint:golint,stylecheck

<<<<<<< HEAD
<<<<<<< HEAD
type list interface {
	// Place your code here
}
=======
// type list interface {
// 	// Place your code here
// 	Len() int                          // длина списка
// 	Front() *listItem                  // первый Item
// 	Back() *listItem                   // последний Item
// 	PushFront(v interface{}) *listItem // добавить значение в начало
// 	PushBack(v interface{}) *listItem  // добавить значение в конец
// 	Remove(i *listItem)                // удалить элемент
// 	MoveToFront(i *listItem)           // поместить элемент в начало
// 	PrintList()
// }
>>>>>>> d2e4e2d... HW4 is completed

=======
>>>>>>> ac6863b... HW4 is completed
type ListItem struct {
	// Place your code here
<<<<<<< HEAD
=======
	Value interface{} // значение
	Next  *ListItem   // следующий элемент
	Prev  *ListItem   // предыдущий элемент
>>>>>>> d2e4e2d... HW4 is completed
}

type List struct {
	// Place your code here
<<<<<<< HEAD
}
<<<<<<< HEAD
=======
=======
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

func (lst *List) PushFront(v interface{}) *ListItem {
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

func (lst *List) PushBack(v interface{}) *ListItem {
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

<<<<<<< HEAD
>>>>>>> eb10d20... HW4 is completed
func (lst *list) PrintList() {
=======
func (lst *List) PrintList() {
>>>>>>> a006b92... HW4 is completed
	itemPtr := lst.First
	for j := 0; j < lst.Size; j++ {
		itemPtr = itemPtr.Prev
	}
<<<<<<< HEAD
>>>>>>> b254570... HW4 is completed

func NewList() List {
	return &list{}
=======
>>>>>>> ac6863b... HW4 is completed
}
