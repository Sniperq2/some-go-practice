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
	Value interface{} // данные
	Next  *ListItem   // указатель на следующий элемент
	Prev  *ListItem   // указатель на предыдущий элемент
}

type list struct {
	Head *ListItem // указатель на голову
	Tail *ListItem // указатель на хвост
	len  int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.Head // вернём "голову"
}

func (l *list) Back() *ListItem {
	return l.Tail // вернем "хвост"
}

func (l *list) PushFront(v interface{}) *ListItem {
	n := &ListItem{Value: v}

	if l.Head != nil {
		n.Next = l.Head
		l.Head.Prev = n
	}

	l.Head = n // новый элемент становится "головой" списка
	l.len++    // увеличиваем счетчик длины списка
	return l.Head
}

func (l *list) PushBack(v interface{}) *ListItem {
	n := &ListItem{Value: v}

	if l.Head != nil { // если есть "голова" то
		tempPointer := l.Head
		for tempPointer.Next != nil {
			tempPointer = tempPointer.Next
		}
		n.Prev = tempPointer
		tempPointer.Next = n
		l.Tail = n // новый элемент пристегнем к хвосту next
	}
	l.len++ // увеличиваем счетчик длины списка
	return l.Tail
}

func (l *list) Remove(i *ListItem) {
	tempPointer := l.Head
	for tempPointer != i { // пробегаем до нужного элемента
		tempPointer = tempPointer.Next
	}
	tempPointer2 := tempPointer.Prev
	tempPointer2.Next = tempPointer.Next
	tempPointer.Next.Prev = tempPointer2
	l.len-- // уменьшаем счетчик длины списка
}

func (l *list) MoveToFront(i *ListItem) {

}
