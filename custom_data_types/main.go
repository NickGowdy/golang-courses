package main

import (
	"custom_data_types/datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("James", "Wilson", organization.NewEuropeanUnionIdentifier("23u8234", "Germany"))
	p.SetTwitterHandler(organization.TwitterHandler("@jam_wils"))
	fmt.Println(p.FullName())
	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectUrl())
	println(p.ID())
}
