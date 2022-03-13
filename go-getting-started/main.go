package main

import (
	"net/http"

	"github.com/pluralsight/webservice/controllers"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

// r := http.Request{Method: "GET"}

// switch r.Method {
// case "GET":
// 	println("GET request")
// }

// switch r.Method {
// case "POST":
// 	println("POST request")
// }
// panic("Something bad just happened")

// for i := 0; i < 5; i++ {
// 	println(i)

// 	if i == 3 {
// 		break
// 	}
// }

// wellKnownPorts := map[string]int{"http": 80, "https": 443}

// for _, v := range wellKnownPorts {
// 	println(v)
// }

// slice := []int{1, 2, 3}
// for i, v := range slice {
// 	println(i, v)
// }

// u := models.User{
// 	ID:        1,
// 	FirstName: "Nick",
// 	LastName:  "Gowdy",
// }

// fmt.Println(u)

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
