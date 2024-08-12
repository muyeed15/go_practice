package arrayslices

import "fmt"

func RunSlices() {

	// slices without capacity
	var slc1 = []int{10, 30, 40}

	fmt.Println(slc1)

	// replace
	slc1[1] = 20

	// append
	slc1 = append(slc1, 50)

	fmt.Println(slc1)

	// slice range
	range1 := slc1[1:3]
	range2 := slc1[2:]
	range3 := slc1[:3]
	range4 := append(range2, 80)
	fmt.Println(range1, range2, range3, range4)

}
