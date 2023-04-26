package pizza

const TaskQueueName = "pizza-tasks"

type Address struct {
	Line1      string
	Line2      string
	City       string
	State      string
	PostalCode string
}

type Customer struct {
	CustomerID int
	Name       string
	Email      string
	Phone      string
}

type Pizza struct {
	Description string
	Price       int
}

type PizzaOrder struct {
	OrderNumber string
	Customer    Customer
	Items       []Pizza
	IsDelivery  bool
	Address     Address
}

type Distance struct {
	Kilometers int
}

type Bill struct {
	CustomerID  int
	OrderNumber string
	Description string
	Amount      int
}

type OrderConfirmation struct {
	OrderNumber        string
	Status             string
	ConfirmationNumber string
	BillingTimestamp   int64
	Amount             int
}
