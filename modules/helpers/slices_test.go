package helpers

import "fmt"

func ExampleDiffrence() {
	slice1 := []string{"Rachid", "Lafriakh", "foo", "bar"}
	slice2 := []string{"Rachid", "Lafriakh"}

	diff := Difference(slice2, slice1)

	fmt.Println(diff)
	// Output:
}
