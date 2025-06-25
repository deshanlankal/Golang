package main
import "fmt"

func main() {
	fmt.Println("Wellcome to something hear")
	fmt.Println("This is a simple Go program.")

	var name = "John Doe"
		fmt.Printf("This is a simple Go program %v is the name variable." , name)// This line will not work as expected because fmt.Println does not support formatting like this
	fmt.Println("This is a simple Go program", name, "is the name variable.") // Correct way to print the variable

	const somemoney = 100 // The value of somemoney is constant and cannot be changed
	// somemoney = 200 // This line would cause an error because somemoney// is a constant
	fmt.Println("The amount of money is:", somemoney)

	var remainItems = 50 //When you declare a variable with var, you can assign it a value later
	fmt.Println("Remaining items:", remainItems)
	remainItems = 30 // You can change the value of a variable declared with var
	fmt.Println("Remaining items:", remainItems)



}

