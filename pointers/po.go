package pointers

import "fmt"

func changeCar(n string) {
	n = "VW"
}

func changePC(x []string) {
	x[2] = "Lenovo"
}

func changeCarPointer(m *string) {
	*m = "VW"
}

func Run() {
	// group a -> non pointer values (Strings, Ints, Floats, Booleans, Array, Structs)
	car := "BMW"
	fmt.Println(car)
	changeCar(car) // not changing

	/*
	       car           n ------ go creates copy
	   |---------|  |---------| - not changing the original value
	   |  0x001  |  |  0x002  |
	   |---------|  |---------|
	   |   BMW   |  |    VW   |
	   |---------|  |---------|
	*/

	fmt.Println(car)

	// group b -> pointer wrapper values (Slices, Maps, Functions)
	pc := []string{"HP", "Dell", "IBM"}
	fmt.Println(pc)
	changePC(pc) // changing (Directly pointing the values)

	/*
			 |-----------------------------------|
			 pc                                  v
		|---------|       |---------|       |---------|       |---------|
		|  0x001  |       |  0x002  |       |  0x003  |       |  0x004  |
		|---------| ----> |---------|       |---------|       |---------|
		| pointer |       |    HP   |       |   Dell  |       |   IBM   |
		|---------|       |---------|       |---------|       |---------|
			 |													   ^
			 |-----------------------------------------------------|
	*/

	fmt.Println(pc)

	// use pointer to modify group a type value
	memory := &car
	value := *memory
	fmt.Println(memory)
	fmt.Println(value)
	changeCarPointer(memory) // chaning for the memory address value modification

	/*
	        |-----------------------------------|
	        v                                   |
	       car              memory              m ----- carbon copy of the pointer
	   |---------|       |---------|       |---------|
	   |  0x001  |       |  0x002  |       |  0x003  |
	   |---------|       |---------|       |---------|
	   |   BMW   | <---- |  p0x001 |       | p0x001  |
	   |---------|       |---------|       |---------|
	*/

	println(car)

}
