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

	if l.Tail == nil {
		l.Tail = n
	}

	l.len++ // увеличиваем счетчик длины списка
	return n
}

func (l *list) PushBack(v interface{}) *ListItem {
	n := &ListItem{Value: v}

	n.Prev = l.Tail // указатель в новом элементе указывает на хвост

	if l.Tail != nil {
		l.Tail.Next = n // новый элемент пристегнем к хвосту next
	}

	l.Tail = n

	if l.Head == nil { // если нет "головы" то пристегнем элемент к ней
		l.Head = n
	}

	l.len++ // увеличиваем счетчик длины списка
	return n
}

func (l *list) Remove(i *ListItem) {
	next := i.Next // у элемента в списке есть уазатель на следующий элемент
	prev := i.Prev // и предыдущий

	if prev != nil {
		prev.Next = next
	}

	if next != nil {
		next.Prev = prev
	}

	l.len-- // уменьшаем счетчик длины списка
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.Head {
		return
	} else if i == l.Tail {
		l.Tail.Prev.Next = nil
		l.Tail = l.Tail.Prev
	} else {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	i.Prev = nil
	i.Next = l.Head
	l.Head.Prev = i
	l.Head = i
}
