package main

func main() {
	arr := [5]byte{0, 1, 2, 3, 4} // They’ve got a fixed size and store elements of the same type in contiguous memory locations.
	println("arr", &arr)

	for i := range arr {
		println(i, &arr[i])
	}
}

// They’ve got a fixed size and store elements of the same type in contiguous memory locations.
// This means Go can access each element quickly since their addresses are calculated based on the starting address of the array and the element’s index.
// arr 0xc000066733
// 0 0xc000066733
// 1 0xc000066734
// 2 0xc000066735
// 3 0xc000066736
// 4 0xc000066737

// There are a couple of things to notice here:
// 1. The address of the array arr is the same as the address of the first element.
// 2. The address of each element is 1 byte apart from each other because our element type is byte.
