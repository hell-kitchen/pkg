package errcodes

import "fmt"

func ExampleFromString() {
	err1 := FromString("unknown status")
	fmt.Println(err1.Error(), err1.Code())

	err2 := FromString(string(CodeIngredientAlreadyExists))
	fmt.Println(err2.Error(), err2.Code())

	// Output:
	// Unknown status IT00
	// ingredient already exists IG09
}
