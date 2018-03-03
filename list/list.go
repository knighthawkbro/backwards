package list

import (
	"fmt"
	"log"
)

// node (Private) - Defines the structure for each individual node in a linked list
type node struct {
	data interface{} // Value of Node
	next *node       // Pointer to the next Node
	prev *node       // pointer to the previous node
	list *List       // Pointer to the list it is attached to
}

// nextNode (Private) - Returns the next node in the list
func (n node) nextNode() *node {
	// returns nil if there is not list AND if the pointer to the next
	// node is the same as the head's next node there for there is next node
	if next := n.next; n.list != nil && next != &n.list.head {
		return next
	}
	return nil
}

// nextNode (Private) - Returns the next node in the list
func (n node) prevNode() *node {
	// returns nil if there is not list AND if the pointer to the next
	// node is the same as the head's next node there for there is next node
	if prev := n.prev; n.list != nil && prev != &n.list.head {
		return prev
	}
	return nil
}

// List (Public) - The container for all the linked nodes in a set
type List struct {
	head    node  // the begining node
	tail    node  // the end of the node
	current *node // the current for the iterator
	count   int   // the count for the iterator
	size    int   // size of the list
}

// init (Private) - Generates a linked list with Size=0 and head pointing to itself
func (l *List) init() *List {
	l.head.next = &l.head
	l.tail.prev = &l.head
	l.size = 0
	return l
}

// New (Public) - Returns an initialized list.
func New() *List { return new(List).init() }

// Add (Public) - Adds item to list
func (l *List) Add(item interface{}) {
	if item == nil {
		log.Println("cannot add nil values to list")
		return
	}
	new := &node{data: item, list: l}
	prev := l.head.next
	prev.prev = new
	new.next = prev
	new.prev = &l.head
	l.head.next = new
	if l.size == 1 {
		l.tail.prev = prev
	}
	l.size++
}

// Remove (Public) - removes item from list, returns true if item found in list
func (l *List) Remove(item interface{}) bool {
	if item == nil {
		log.Println("cannot add nil values to list")
		return false
	}
	if l.size == 0 {
		return false
	}
	if l.size == 1 {

	}
	current := &l.head
	for x := 0; x < l.size; x++ {
		if current.next.data == item {
			if x != l.size-1 {
				current.next.next.prev = current
				current.next = current.next.next
			} else {
				current.next = nil
				if l.size == 1 {
					current.next = &l.head
				}
				l.tail.prev = current
			}
			l.size--
			return true
		}
		current = current.next
	}
	return false
}

// Contains (Public) - Returns true or false whether an item was contained in the list
func (l *List) Contains(item interface{}) bool {
	for current := l.head.next; current != nil; current = current.nextNode() {
		if current.data == item {
			return true
		}
	}
	return false
}

// AddIndex (public) - adds item at specified index
func (l *List) AddIndex(index int, item interface{}) {
	if item == nil {
		log.Println("Cannot add nil value to list")
		return
	}
	err := l.checkIndex(index)
	if err != nil {
		log.Println(err)
		return
	}
	new := &node{data: item, list: l}
	if index == 0 {
		prev := l.head.next
		prev.prev = new
		new.next = prev
		l.head.next = new
	} else {
		current := l.head.next
		for x := 0; x < index-1; x++ {
			current = current.nextNode()
		}
		new.next = current.next
		new.next.prev = new
		new.prev = current
		current.next = new
	}
	l.size++
}

// Set (Public) - replaces item at specified index with new value,
// returns original value
func (l *List) Set(index int, item interface{}) interface{} {
	if item == nil {
		log.Println("Cannot store a nil value in list")
		return nil
	}
	err := l.checkIndex(index)
	if err != nil {
		log.Println(err)
		return nil
	}
	current := l.head.next
	for x := 0; x < index; x++ {
		current = current.nextNode()
	}
	removed := current.data
	current.data = item
	return removed
}

// RemoveIndex (Public) - removes item at specified index,
// returns removed item
func (l *List) RemoveIndex(index int) interface{} {
	err := l.checkIndex(index)
	if err != nil {
		log.Println(err)
		return nil
	}
	if index == 0 {
		result := l.head.next.data
		l.head.next = l.head.next.next
		if l.size > 1 {
			l.head.next.next.prev = l.head.next
		}
		l.head.next.prev = nil
		l.size--
		return result
	}
	current := l.head.next
	for x := 0; x < index-1; x++ {
		current = current.next
	}
	removed := current.next.data
	current.next = current.next.next
	current.next.prev = current
	l.size--
	return removed
}

// Get (Public) - returns item at specified index
func (l *List) Get(index int) interface{} {
	err := l.checkIndex(index)
	if err != nil {
		log.Println(err)
		return nil
	}
	current := l.head.next
	for x := 0; x < index; x++ {
		current = current.nextNode()
	}
	return current.data
}

// IndexOf (Public) - returns index of specified item,
// returns -1 if item not in list
func (l *List) IndexOf(item interface{}) int {
	current := l.head.next
	for x := 0; x < l.size; x++ {
		if current.data == item {
			return x
		}
		current = current.nextNode()
	}
	return -1
}

// Size (Public) - Returns the size of the list
func (l *List) Size() int {
	return l.size
}

// String (Public) - Allows for the fmt.Print* functions to print the list struct as a string.
func (l *List) String() string {
	if l.size == 0 {
		return "[ ]"
	}
	result := "[ "
	for current := l.head.next; current != nil; current = current.nextNode() {
		result += fmt.Sprintf("%v ", current.data)
	}
	return result + "]"
}

// checkIndex (Private) - Checks to see if the index provided is within the range of indices
// of the list.
func (l *List) checkIndex(index int) error {
	if index < 0 || index >= l.size {
		return fmt.Errorf("index out of range")
	}
	return nil
}

// Previous (Public) - Since this list puts objects in the beginning
// to traverse previously you would need to start from the beginning.
// The count however needs to count down
func (l *List) Previous() func() interface{} {
	if l.current == nil {
		l.count = l.size
		l.current = l.head.next
	}
	if l.HasPrevious() {
		return func() interface{} {
			if l.count < 0 {
				log.Println("Off end of list")
				return nil
			}
			l.count--
			l.current = l.current.next
			return l.current.prev.data
		}
	}
	return func() interface{} {
		log.Println("Off end of list")
		return nil
	}
}

// HasPrevious (Public) - will return true while count is greater than zero.
func (l *List) HasPrevious() bool {
	return l.count > 0
}

// PreviousIndex (Public) - returns the current index which is always +1 from
// the real current
func (l *List) PreviousIndex() int {
	return l.count
}

// Next (Public) - Haven't fixed this function since it is unused.
// I think it has a bug
func (l *List) Next() func() interface{} {
	if l.current == nil || l.current == &l.head {
		l.current = l.tail.prev
		l.count = l.size
	}
	if l.HasNext() {
		return func() interface{} {
			l.count--
			if l.count < 0 {
				log.Println("Off end of list")
				return nil
			}
			l.current = l.current.prev
			return l.current.next.data
		}
	}
	return func() interface{} {
		log.Println("Off end of list")
		return nil
	}
}

// HasNext (Public) - will return true until the count is greater than or equal to
// the size of the list
func (l *List) HasNext() bool {
	return l.count < l.size
}

// NextIndex (Public) - Returns the index for the next index
func (l *List) NextIndex() int {
	return l.count - 1
}
