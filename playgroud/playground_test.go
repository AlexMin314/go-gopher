package playground

import (
	"fmt"
	"sort"
)

func ExampleSlice() {
	Slice()
	// Output:
	// [1 2 3 4 5]
	// len: 5
	// cap: 5
	// [1 2 3]
	// len: 3
	// cap: 5
}

func ExampleCopy() {
	fmt.Println(CopySlice([]int{1, 2, 3}))
	// Output:
	// [1 2 3]
}
func ExampleCopy2() {
	fmt.Println(CopySlice2([]int{1, 2, 3}))
	// Output:
	// [1 2 3]
}

func ExampleInsert() {
	fmt.Println(InsertSlice([]int{1, 2, 3}, 5, 2))
	fmt.Println(InsertSlice2([]int{1, 2, 3}, 5, 2))
	fmt.Println(InsertSlice3([]int{1, 2, 3}, []int{4, 5}, 2))
	// Output:
	// [1 2 5 3]
	// [1 2 5 3]
	// [1 2 4 5 3]
}

func ExampleDelete() {
	fmt.Println(DeleteSlice([]int{1, 2, 3}, 1))
	fmt.Println(DeleteSlice2([]int{1, 2, 3}, 1))
	// Output:
	// [1 3]
	// [1 3]
}

func ExmapleCut() {
	fmt.Println(CutSlice([]int{1, 2, 3, 4, 5}, 1, 2))
	// Output:
	// [1 4 5]
}

func ExmapleMap() {
	PlayMap()
	// Output:
	// "world" true
	// false
}

func ExampleSum() {
	fmt.Println(Sum([]int{1, 2}...))
	// Output:
	// 3
}

func ExampleCounter() {
	codeCount := map[rune]int{}
	CountChar("hello", codeCount)
	var keys sort.IntSlice
	for key := range codeCount {
		keys = append(keys, int(key))
	}
	sort.Sort(keys)
	for _, key := range keys {
		fmt.Println(string(key), codeCount[rune(key)])
	}
	// Output:
	// e 1
	// h 1
	// l 2
	// o 1
}

func ExampleStruct() {
	LogStruct()
	// Output:
	// {1 2} {1 0} {0 0}

}
