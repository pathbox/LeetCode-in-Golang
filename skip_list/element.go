package skiplist

func (element *Element) Next() *Element {
	return element.next[0]
}

func (element *Element) NextLevel(level int) *Element {
	if level >= len(element.next) || level < 0 {
		panic("invalid argument to NextLevel")
	}

	return element.next[level]
}

func (element *Element) Key() interface{} {
	return element.Key
}