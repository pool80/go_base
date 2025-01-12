package internal

type Customer struct {
	Name         string
	Age          int
	Balance      int
	Debt         int
	Discount     bool
	CalcDiscount func() (int, error)
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		Discount: discount,
	}
}

// реализация функции CalcPrice
func CalcPrice(cust Customer, price int) (int, error) {
    discountAmount, err := cust.CalcDiscount()
    if err != nil {
        return 0, err
    }
    finalPrice := price - discountAmount
    if finalPrice < 0 {
        finalPrice = 0
    }
    return finalPrice, nil
}