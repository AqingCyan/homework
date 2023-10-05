package main

import "fmt"

func main() {
	testOne := []int{1, 2, 3, 4, 5}
	fmt.Println(removeSliceItemV1(testOne, 1))

	testTwo := []int{1, 2, 3, 4, 5}
	fmt.Println(removeSliceItemV2(testTwo, 1))

	testThree := []int{1, 2, 3, 4, 5}
	fmt.Println(removeSliceItemV3(testThree, 1))

	testFour := []int{1, 2, 3, 4, 5}
	fmt.Println(removeSliceItemV4(testFour, 1))

	testFive := []int{1, 2, 3, 4, 5}
	fmt.Println(removeSliceItemV5(testFive, 1))
}

func removeSliceItemV1(slice []int, index int) []int {
	if index < 0 {
		fmt.Println("index 不能小于 0")
		return slice
	}

	if index >= len(slice) {
		fmt.Println("index 超出 slice 范围")
		return slice
	}

	preSlice := slice[:index]
	tailSlice := slice[index+1:]

	return append(preSlice, tailSlice...)
}

// removeSliceItemV2 上面的 V1 使用append来删除slice的元素可能会导致内存分配和复制
// 在不需要保留原始切片的有序性的情况下，高效的方法是将要删除的元素与最后一个元素交换，然后删除最后一个元素。
func removeSliceItemV2(slice []int, index int) []int {
	if index < 0 {
		fmt.Println("index 不能小于 0")
		return slice
	}

	if index >= len(slice) {
		fmt.Println("index 超出 slice 范围")
		return slice
	}
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// removeSliceItemV3 如果需要保持元素的顺序,则可以使用如下方法,但是其性能可能低于 V2
func removeSliceItemV3(slice []int, index int) []int {
	if index < 0 {
		fmt.Println("index 不能小于 0")
		return slice
	}

	if index >= len(slice) {
		fmt.Println("index 超出 slice 范围")
		return slice
	}

	copy(slice[index:], slice[index+1:])

	return slice[:len(slice)-1]
}

func removeSliceItemV4[T any](slice []T, index int) []T {
	if index < 0 {
		fmt.Println("index 不能小于 0")
		return slice
	}

	if index >= len(slice) {
		fmt.Println("index 超出 slice 范围")
		return slice
	}

	copy(slice[index:], slice[index+1:])

	return slice[:len(slice)-1]
}

func removeSliceItemV5[T any](slice []T, index int) []T {
	if index < 0 {
		fmt.Println("index 不能小于 0")
		return slice
	}

	if index >= len(slice) {
		fmt.Println("index 超出 slice 范围")
		return slice
	}

	copy(slice[index:], slice[index+1:])
	slice = slice[:len(slice)-1]

	// 引入缩容机制
	newSlice := make([]T, len(slice))
	copy(newSlice, slice)
	return newSlice
}
