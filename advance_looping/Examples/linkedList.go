package Examples

import (
	"fmt"
	"math/rand"
	"time"
)

type linkedList struct {
	value   int
	pointer *linkedList
}

func StartLinkedList() {
	rand.Seed(time.Now().UnixNano())

	list := createList(10, nil)

	for ; list != nil; list = list.pointer {
		fmt.Printf("%v/n", list.value)
	}
}

func createList(num int, temp *linkedList) *linkedList {
	if temp == nil {
		temp = &linkedList{rand.Intn(100), nil}
		num--
	}

	tempList := temp

	for i := 0; i < num; i++ {
		t := &linkedList{rand.Intn(100), nil}
		tempList.pointer = t
		tempList = tempList.pointer
	}
	return temp
}
