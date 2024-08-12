package arrayslices

import "fmt"

func RunArray() {

	// type with capacity
	var arr1 [3]int = [3]int{20, 25, 30}

	// without type
	var arr2 = [4]string{"Muyeed", "Dipra", "Alam"}

	fmt.Println(arr1, len(arr2))

}
