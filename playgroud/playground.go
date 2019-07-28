package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ConcatStrs(s1 string, s2 string) string {
	ps1 := &s1
	fmt.Println(ps1)
	// return s1 + " " + s2
	return strings.Join([]string{s1, s2}, " ")
}

func Slice() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))

	sliced := nums[:3]
	fmt.Println(sliced)
	fmt.Println("len:", len(sliced))
	fmt.Println("cap:", cap(sliced))
}

func CopySlice(src []int) []int {
	dest := make([]int, len(src))
	copy(dest, src)
	return dest
}

func CopySlice2(src []int) []int {
	dest := append([]int(nil), src...)
	return dest
}

func InsertSlice(src []int, ele int, i int) []int {
	src = append(src[:i+1], src[i:]...)
	src[i] = ele
	return src
}

func InsertSlice2(src []int, ele int, i int) []int {
	src = append(src, ele)
	copy(src[i+1:], src[i:])
	src[i] = ele
	return src
}

func InsertSlice3(src []int, sl []int, i int) []int {
	src = append(src, sl...)
	copy(src[i+len(sl):], src[i:])
	copy(src[i:], sl)
	return src
}

func DeleteSlice(src []int, i int) []int {
	return append(src[:i], src[i+1:]...)
}

func DeleteSlice2(src []int, i int) []int {
	// input: [1 2 3]
	copy(src[i:], src[i+1:])
	// src: [1 3 3]
	src[len(src)-1] = 0
	// src: [1 3 0]
	src = src[:len(src)-1]
	// src: [1 3]
	return src
}

func CutSlice(src []int, i int, j int) []int {
	// input: [1 2 3 4 5], 1, 2
	copy(src[i:], src[j:])
	// src: [1 4 5 4 5]
	for k, n := len(src)-j+i, len(src); k < n; k++ {
		src[len(src)-1-i] = 0
	}
	// src: [1 4 5 0 0]
	src = src[:len(src)-j+i]
	// src: [1 4 5]
	return src
}

func StringToNumber() {
	var i int
	var k int64
	var f float64
	var s string
	var s2 string
	var err error
	i, err = strconv.Atoi("350")
	k, err = strconv.ParseInt("cc7fdd", 16, 32)
	f, err = strconv.ParseFloat("3.14", 64)
	s = strconv.Itoa(340)
	s2 = strconv.FormatInt(1340277, 16)
	fmt.Println(i, k, f, s, s2, err)
}

func PlayMap() {
	// m := make(map[string]int)
	m := map[string]string{}
	m["hello"] = "world"
	v, ok := m["hello"]
	fmt.Println(v, ok)
	m["go"] = "lang"
	delete(m, "hello")
	_, prs := m["hello"]
	fmt.Println(prs)

	kvs := map[string]string{
		"hello": "wolrd",
	}
	num := 0
	for k, v := range kvs {
		num += 1
		fmt.Println(k, v)
	}
	fmt.Println(num)
}

func Sum(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func CountChar(s string, codeCount map[rune]int) {
	for _, r := range s {
		codeCount[r]++
	}
}

func hasDupeRune(s string) bool {
	runeSet := map[rune]struct{}{}
	for _, r := range s {
		if _, prs := runeSet[r]; prs {
			return true
		}
		runeSet[r] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(ConcatStrs("Hello", "World"))
	StringToNumber()
}
