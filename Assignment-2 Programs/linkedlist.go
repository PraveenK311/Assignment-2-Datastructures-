package main

import (
    "fmt"
    "container/list"
)

func main() {
    fmt.Println("Go Linked Lists Tutorial")
    mylist := list.New()
    mylist.PushBack(1)
    mylist.PushFront(2)
    // we now have a linked list with '1' at the back of the list
    // and '2' at the front of the list.
}