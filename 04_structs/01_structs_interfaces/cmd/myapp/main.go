package main

import (
	"errors"
	"fmt"
	"myapp/internal"
)

const DEFAULT_DISCOUNT = 500

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)

	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("discount not available")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}

	price := 1000 // Пример цены товара
    finalPrice, err := internal.CalcPrice(*cust, price)
    if err != nil {
        fmt.Printf("Error calculating price: %v\n", err)
    } else {
        fmt.Printf("Final price: %d\n", finalPrice)
    }
}
