package iterator

import "fmt"

// Collection Collection interface
type Collection interface {
	CreateIterator() Iterator
	AddItem(item string)
}

// StringCollection String collection implementation
type StringCollection struct {
	items []string
}

func (c *StringCollection) CreateIterator() Iterator {
	return &StringIterator{collection: c, index: -1}
}

func (c *StringCollection) AddItem(item string) {
	c.items = append(c.items, item)
}

// Iterator Iterator interface
type Iterator interface {
	Next() bool
	Current() string
}

// StringIterator String iterator implementation
type StringIterator struct {
	collection *StringCollection
	index      int
}

func (i *StringIterator) Next() bool {
	i.index++
	return i.index < len(i.collection.items)
}

func (i *StringIterator) Current() string {
	if i.index < len(i.collection.items) {
		return i.collection.items[i.index]
	}
	return ""
}

// Usage example
func TraverseCollection(collection Collection) {
	iterator := collection.CreateIterator()
	for iterator.Next() {
		fmt.Println(iterator.Current())
	}
}
