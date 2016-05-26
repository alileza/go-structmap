package structmap

import "fmt"

type StructAddress struct {
	Address    string
	PostalCode int64
}

type Person struct {
	Name    string
	Age     int
	Address StructAddress
	Married bool
	Height  float32
	Weight  float64
}

func main() {
	a := &Person{}

	fmt.Println(StructToMap(*a))
	fmt.Println(StructToMap(*a, true))
}
