package main

import (
	"fmt"
	"task01/customer"
	"errors"
)

func example02() {
	user := customer.New("John", 25, 322, 0, true)

	user.CalcDiscount = func() (int, error) {
		if !user.Discount {
			return 0, errors.New("discount not available")
		}
		result := customer.DefaultDiscount - user.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil 
	}

	discount, err := user.CalcDiscount()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("user discount: %+v\n", discount)
}
