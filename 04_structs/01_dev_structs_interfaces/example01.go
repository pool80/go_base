package main

import (
	"fmt"
	"task01/customer"
)

func example01() {
	var firstUser customer.Customer
	firstUser.Name = "Alex"
	firstUser.Debt = 3000

	secondUser := customer.Customer{
		Name: "Bob",
		Debt: 2000,
	}

	thirdUser := new(customer.Customer)
	thirdUser.Name = "Alice"
	thirdUser.Debt = 0

	fourthUser := customer.New("John", 25, 322, 0, false)

    fmt.Printf("First user: %+v\n", firstUser)
    fmt.Printf("Second user: %+v\n", secondUser)
    fmt.Printf("Third user: %+v\n", thirdUser)
    fmt.Printf("Fourth user: %+v\n", fourthUser)
}
