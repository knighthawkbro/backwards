package main

import (
	"backwards/array"
	"backwards/list"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// dblinklist (Private) - This is the interface that array/list implements
// for this lab.
type dblinklist interface {
	Add(item interface{})
	Remove(item interface{}) bool
	Size() int
	String() string

	AddIndex(index int, item interface{})
	Get(index int) interface{}
	RemoveIndex(index int) interface{}
	Set(index int, item interface{}) interface{}
	IndexOf(item interface{}) int

	Next() func() interface{}
	HasNext() bool
	NextIndex() int
	Previous() func() interface{}
	HasPrevious() bool
	PreviousIndex() int
}

func main() {
	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning driver function as an array...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	arr := array.New()
	driver(arr)
	arr = nil
	arr = array.New()
	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning backwards function as an array...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	backwards(arr)

	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning driver function as a list...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	lst := list.New()
	driver(lst)
	lst = nil
	lst = list.New()
	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning backwards function as a list...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	backwards(lst)
}

// backwards (private) - Reads the message.txt file, store all the
// words in a doubly-linked list, reads the list with an iterator
// backwards skipping everyother word.
func backwards(words dblinklist) {
	file, err := os.Open("message.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, word := range strings.Split(scanner.Text(), " ") {
			word = reg.ReplaceAllString(word, "")
			words.Add(strings.ToLower(word))
		}
	}
	iter := words.Previous()
	for x := 1; x <= words.Size(); x++ {
		if (x % 2) == 0 {
			fmt.Println(iter())
		} else {
			iter()
		}
	}
}

// driver (Private) - Performs the basic test to try each function
// defined in the dblinklist interface.
func driver(words dblinklist) {
	structures := []string{"bag", "set", "queue", "stack"}
	fmt.Println("\nhere's our list")
	for _, word := range structures {
		words.Add(word)
	}
	fmt.Println(words)

	// removing specific item
	fmt.Println("\nremoving bag")
	words.Remove("bag")
	fmt.Println(words)
	fmt.Println("removing stack")
	words.Remove("stack")
	fmt.Println(words)
	fmt.Println("removing queue")
	words.Remove("queue")
	fmt.Println(words)
	fmt.Println("removing set")
	words.Remove("set")
	fmt.Println(words)
	fmt.Println("trying to remove set again")
	words.Remove("set")
	fmt.Println(words)

	fmt.Println("\nadding items back in")
	for _, word := range structures {
		words.Add(word)
	}
	fmt.Println(words)

	fmt.Println("\nremoving item at first index")
	fmt.Println(words.RemoveIndex(0))
	fmt.Println(words)
	fmt.Println("removing item at last index")
	fmt.Println(words.RemoveIndex(2))
	fmt.Println(words)
	fmt.Println("removing item at index 1")
	fmt.Println(words.RemoveIndex(1))
	fmt.Println(words)
	fmt.Println("removing item at index 0")
	fmt.Println(words.RemoveIndex(0))
	fmt.Println(words)

	fmt.Println("\nadding items in once more")
	for _, word := range structures {
		words.Add(word)
	}
	fmt.Println(words)
	// reverse iterator

	fmt.Println("\nprinting indices and items in reverse order")
	iter := words.Previous()
	for words.HasPrevious() {
		fmt.Println(words.PreviousIndex(), " ", iter())
	}
	fmt.Println("Trying to remove item from empty list")
	iter()
}
