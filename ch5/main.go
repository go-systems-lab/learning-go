package main

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) error {
	return nil
}

func main() {
	// using a struct to simulate named parameters
	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		Age:       24,
		LastName:  "Smith",
	})
}
