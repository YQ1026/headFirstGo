package magazine

type Address struct {
	City       string
	PostalCode int
}

type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}
