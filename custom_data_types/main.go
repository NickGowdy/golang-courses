package main

import (
	"custom_data_types/datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("James", "Wilson")
	p.SetTwitterHandler("@jam_wils")
	fmt.Println(p.FullName())
	fmt.Println(p.TwitterHandler())
	println(p.Id())
}
