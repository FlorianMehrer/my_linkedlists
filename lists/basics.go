package lists

import (
	"fmt"
	"linkedlists/elements"
)

func New() *List {
	return &List{head: elements.NewEmpty()}
}

func (l *List) IsEmpty() bool {
	return l.head.IsEmpty()
}

func (l *List) Append(value int) {
	l.head.Append(value)
}

func (l *List) String() string {
	return fmt.Sprintf("[%v]", l.head)
}

func (l *List) Length() int {
	return l.head.Length()
}

func (l *List) Contains(value int) bool {
	return l.head.Contains(value)
}
