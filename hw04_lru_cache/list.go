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
}

type list struct {
	List
	Head *ListItem
	Tail *ListItem
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

	if l.Head == nil { // если "головы" нет, значит и списка нет
		l.Head = n // добавим новый элемент к списку
	} else { // если есть "голова" то
		temp := l.Head     // сохраняем предыдущую "голову" в temp
		l.Head = n         // новый элемент становится "головой" списка
		n.Next = temp.Prev // меняем указатель prev на next нового списка
	}
	l.len++ // увеличиваем счетчик длины списка
	return l.Head
}

func (l *list) PushBack(v interface{}) *ListItem {
	n := &ListItem{Value: v}
	if l.Head == nil { // если "головы" нет, значит и списка нет
		l.Head = n // добавим новый элемент к списку
	} else { // если есть "голова" то
		l.Tail.Next = n // новый элемент пристегнем к хвосту next
	}

	l.len++ // увеличиваем счетчик длины списка

	return l.Tail
}

func (l *list) Remove(i *ListItem) {

}

func (l *list) MoveToFront(i *ListItem) {

}
