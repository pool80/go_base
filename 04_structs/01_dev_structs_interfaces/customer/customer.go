package customer

const DefaultDiscount = 10

type Customer struct {
	Name         string
	Age          int
	Balance      int
	Debt         int
	Discount     bool
	CalcDiscount func() (int, error)
}

func New(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		Discount: discount,
	}
}