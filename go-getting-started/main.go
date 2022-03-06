package main

import (
	"fmt"

	"github.com/pluralsight/webservice/models"
)

func main() {
	u := models.User{
		ID:        1,
		FirstName: "Nick",
		LastName:  "Gowdy",
	}

	fmt.Println(u)
}

// const (
// 	first = iota + 6
// 	second
// )

// const (
// 	third = iota
// 	fourth
// )

func slice() {

	// slice := []int{1, 2, 3}
	// fmt.Println(slice)

	// slice = append(slice, 4, 42, 27)
	// fmt.Println(slice)

	// s2 := slice[1:]
	// s3 := slice[:2]
	// s4 := slice[1:2]
	// fmt.Println(s2, s3, s4)

	// fmt.Println(first, second, third, fourth)

	// arr := [3]int{1, 2, 3}

	// slice := arr[:]

	// arr[1] = 42
	// slice[2] = 27

	// fmt.Println(arr, slice)
}

func mapCode() {
	// m := map[string]int{"foo": 42}
	// fmt.Println(m)
	// fmt.Println(m["foo"])

	// m["foo"] = 27
	// fmt.Println(m)

	// delete(m, "foo")
	// fmt.Println(m)
}

func structCode() {
	// type user struct {
	// 	ID        int
	// 	FirstName string
	// 	LastName  string
	// }

	// var u user
	// u.ID = 1
	// u.FirstName = "Nick"
	// u.LastName = "Gowdy"

	// u2 := user{ID: 2,
	// 	FirstName: "Victoria",
	// 	LastName:  "Pinto Pane"}
	// fmt.Println(u, u2)
}
