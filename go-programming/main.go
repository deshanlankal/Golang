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


	////////////////////////////////////////////////////////////////////////////

	var username string
	var age int
	//let the user deside what is the data type

	username = "Deshan" // Assign a value to the variable
	age = 23
	fmt.Printf("Username: %v Age: %v\n", username, age) // Print the values of the variables

	// Print the data type of username
	fmt.Printf("The data type of username is: %T\n", username)

	// Example: Take user input for a new variable
	var city string
	fmt.Print("Enter your city: ")
	fmt.Scanln(&city)
	fmt.Printf("You live in: %v\n", city)
	// Example of different data types in Go

	// Integer types
	var intVar int = 42
	var int8Var int8 = -8
	var uintVar uint = 100

	// Floating point types
	var float32Var float32 = 3.14
	var float64Var float64 = 2.718281828

	// Boolean type
	var boolVar bool = true

	// String type
	var strVar string = "Hello, Go!"

	// Byte and rune types
	var byteVar byte = 'A'      // alias for uint8
	var runeVar rune = '世'     // alias for int32, supports Unicode

	// Array type
	var arrVar [3]int = [3]int{1, 2, 3}

	// Slice type
	var sliceVar []string = []string{"apple", "banana", "cherry"}

	// Map type
	var mapVar map[string]int = map[string]int{"one": 1, "two": 2}

	// Struct type
	type Person struct {
		Name string
		Age  int
	}
	var personVar Person = Person{Name: "Alice", Age: 30}

	// Print all variables and their types
	fmt.Printf("intVar: %v (%T)\n", intVar, intVar)
	fmt.Printf("int8Var: %v (%T)\n", int8Var, int8Var)
	fmt.Printf("uintVar: %v (%T)\n", uintVar, uintVar)
	fmt.Printf("float32Var: %v (%T)\n", float32Var, float32Var)
	fmt.Printf("float64Var: %v (%T)\n", float64Var, float64Var)
	fmt.Printf("boolVar: %v (%T)\n", boolVar, boolVar)
	fmt.Printf("strVar: %v (%T)\n", strVar, strVar)
	fmt.Printf("byteVar: %v (%T)\n", byteVar, byteVar)
	fmt.Printf("runeVar: %v (%T)\n", runeVar, runeVar)
	fmt.Printf("arrVar: %v (%T)\n", arrVar, arrVar)
	fmt.Printf("sliceVar: %v (%T)\n", sliceVar, sliceVar)
	fmt.Printf("mapVar: %v (%T)\n", mapVar, mapVar)
	fmt.Printf("personVar: %+v (%T)\n", personVar, personVar)

	 
}

