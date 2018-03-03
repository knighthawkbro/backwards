package array

import (
	"fmt"
	"log"
)

// Array (Public) - Structure that defines
type Array struct {
	count      int
	size       int
	collection []interface{}
}

// Init (Public) - initializes the array with whatever size is provided, This is what can be overrided by the user.
func (a *Array) Init(capacity int) *Array {
	if capacity < 0 {
		return nil
	}
	a.collection = make([]interface{}, capacity)
	a.size = 0
	return a
}

// New (Public) - Returns an initialized array with default size of 10.
func New() *Array { return new(Array).Init(10) }

// Add (Public) - Adds a new item to the end of the list
func (a *Array) Add(item interface{}) {
	if item == nil {
		log.Println("Cannot store a nil value")
		return
	}
	a.ensureSpace()
	a.collection[a.size] = item
	a.size++
}

// Remove (Public) - Removes the first instance of the item in the list.
func (a *Array) Remove(item interface{}) bool {
	if item == nil {
		log.Println("Cannot remove a nil value")
		return false
	}
	for x := 0; x < a.size; x++ {
		if item == a.collection[x] {
			a.shiftLeft(x)
			a.size--
			return true
		}
	}
	return false
}

// Contains (Public) - Checks to see if item is in array, returns true or false
func (a *Array) Contains(item interface{}) bool {
	for x := 0; x < a.size; x++ {
		if item == a.collection[x] {
			return true
		}
	}
	return false
}

// AddIndex (Public) - adds an item at the index specified.
func (a *Array) AddIndex(index int, item interface{}) {
	if item == nil {
		log.Println("Cannot store nil value")
		return
	}
	err := a.checkIndex(index)
	if err != nil {
		log.Println("index out of bounds")
		return
	}
	a.ensureSpace()

	a.shiftRight(index)
	a.collection[index] = item
	a.size++
}

// Set (Public) - Overwrites the data at a current index, returns old
// value stored.
func (a *Array) Set(index int, item interface{}) interface{} {
	if item == nil {
		log.Println("Cannot store nil value")
		return nil
	}
	err := a.checkIndex(index)
	if err != nil {
		log.Println("index out of bounds")
		return nil
	}

	removed := a.collection[index]
	a.collection[index] = item
	return removed
}

// RemoveIndex (Public) - Returns the item at the index and overwrites
// the item by shifting all newer items leftwards.
func (a *Array) RemoveIndex(index int) interface{} {
	err := a.checkIndex(index)
	if err != nil {
		log.Println("index out of bounds")
		return nil
	}
	removed := a.collection[index]
	a.shiftLeft(index)

	a.size--
	a.collection[a.size] = nil
	return removed
}

// Get (Public) - Returns the value in the given index.
func (a *Array) Get(index int) interface{} {
	if a.size == 0 {
		return nil
	}
	err := a.checkIndex(index)
	if err != nil {
		log.Println("index out of bounds")
		return nil
	}
	return a.collection[index]
}

// IndexOf (Public) - traverses the list forwards to find the first
// instance of an item. Returns -1 if nothing is found.
func (a *Array) IndexOf(item interface{}) int {
	for x := 0; x < a.size; x++ {
		if item == a.collection[x] {
			return x
		}
	}
	return -1
}

// LastIndexOf (Public) - Traverse the list backwards to find the last
// instance of an item and returns the index. Returns -1 if nothing
// is found.
func (a *Array) LastIndexOf(item interface{}) int {
	for x := a.size - 1; x >= 0; x-- {
		if item == a.collection[x] {
			return x
		}
	}
	return -1
}

// Size (Public) - returns the size of the Array
func (a *Array) Size() int {
	return a.size
}

// String (Public) - formats the array when fmt.Print is called.
func (a *Array) String() string {
	if a.size == 0 {
		return "[ ]"
	}
	s := "[ "
	for x := 0; x < a.size; x++ {
		s += fmt.Sprintf("%v ", a.collection[x])
	}
	return s + "]"
}

// HasNext (Public) - iterates over the list and returns the node data
func (a *Array) HasNext() bool {
	return a.count < a.size
}

// Next (Public) - returns the current node and moves to the next node
func (a *Array) Next() func() interface{} {
	if a.HasNext() {
		return func() interface{} {
			if a.count >= a.size {
				log.Println("Off end of list")
				return nil
			}
			a.count++
			return a.collection[a.count-1]
		}
	}
	return func() interface{} {
		log.Println("Off end of list")
		return nil
	}
}

// NextIndex (Public) - returns the current index.
func (a *Array) NextIndex() int {
	return a.count
}

// HasPrevious (Public) - Returns true if count is greater than zero
func (a *Array) HasPrevious() bool {
	return a.count > 0
}

// Previous (Public) - A closure that returns the previous
// node data each time it is called.
func (a *Array) Previous() func() interface{} {
	if a.count == 0 {
		a.count = a.size
	}
	if a.HasPrevious() {
		return func() interface{} {
			if a.count == 0 {
				log.Println("Off end of list")
				return nil
			}
			a.count--
			return a.collection[a.count]
		}
	}
	return func() interface{} {
		log.Println("Off end of list")
		return nil
	}
}

// PreviousIndex (Public) - Returns the current count
func (a *Array) PreviousIndex() int {
	return a.count
}

// ensureSpace (Private) - Sees if the size and capacity of the array are the same. If so,
// It creates a new array with double the capacity and overwrites the old array with a new
// array, then clears the new array for the GC.
func (a *Array) ensureSpace() {
	if a.size == cap(a.collection) {
		new := new(Array).Init(cap(a.collection) * 2)
		new.size = a.size
		for i := 0; i < a.size; i++ {
			new.collection[i] = a.collection[i]
		}
		*a = *new
		new = nil
	}
}

// checkIndex (Private) -
func (a *Array) checkIndex(index int) error {
	if index < 0 || index >= a.size {
		return fmt.Errorf("index outside of list")
	}
	return nil
}

// shiftLeft (Private) - Moves all the items left after index (Destructive)
func (a *Array) shiftLeft(index int) {
	for i := index; i < a.size-1; i++ {
		a.collection[i] = a.collection[i+1]
	}
}

// shiftRight (Private) - Moves all the items to the right after the index (non-destructive)
func (a *Array) shiftRight(index int) {
	for i := a.size; i > index; i-- {
		a.collection[i] = a.collection[i-1]
	}
}
