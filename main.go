package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func unusedFn_with_looooooooong_name_and_empty_return() {
	return
}

func main() {

	var intsArray []int
	var firstArgument int
	ia := &intsArray
	fa := &firstArgument
	var err error

	task165d([]int{1, 2, 3, 4, 5})

	// set flags
	exeTask108 := flag.Bool("108", false, "execute task 108(int)")
	exeTask331a := flag.Bool("331a", false, "execute task 331a(int)")
	exeTask165d := flag.Bool("165d", false, "execute task 165d([]int)")
	flag.Parse()

	//check args
	if len(os.Args) < 2 {
		fmt.Println("Expect at least one flag with args:\n-108 (int)\n-331a (int)\n-165d ([]int)")
	} else if len(os.Args) > 2 {

		*fa, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: arg type")
			return
		}
		argsArray := os.Args[2:]
		for _, s := range argsArray {
			num, err := strconv.Atoi(s)
			if err == nil {
				*ia = append(*ia, num)
			}
		}
	} else {
		fmt.Println("Error: Empty Args.")
		return
	}

	// choose flag
	if len(os.Args) > 0 {

		switch {
		case *exeTask108:
			fmt.Println(task108(firstArgument))
			return
		case *exeTask331a:
			fmt.Println(task331a(firstArgument))
			return
		case *exeTask165d:
			fmt.Println(task165d(intsArray))
			return
		}
	}
}

// повертає найменше число у вигляді 2^n що є більшим від заданого.
// обробляються лише цілі додатні числа.
func task108(x int) int {
	if x < 2 {
		return 0
	}
	res := 1
	pow := 1
	for res <= x {
		res *= 2
		pow++
	}
	return res
}

// повертає представлення числа n у виді трьох чисел, сума квадратів яких дорівнює n.
func task331a(n int) (int, int, int) {
	var flag bool
	var a, b, c int
	for x := 0; x*x < n; x++ {
		if flag {
			break
		}
		for y := x; y*y < n; y++ {
			if flag {
				break
			}
			for z := y; z*z < n; z++ {
				if flag {
					break
				}
				if x*x+y*y+z*z == n {
					// fmt.Println("Success!")
					// fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)
					a, b, c = x, y, z
					flag = true
				}
			}
		}
	}
	if !flag {
		// fmt.Printf("Fail, no solution for  %d.\n", n)
		return a, b, c
	}
	// fmt.Println(a, b, c)
	return a, b, c
}

// створення нового зрізу до першого від'ємного числа.
// вхід в рекурсивну функцію.
func task165d(arr []int) []int {
	var destArr []int
	var newArr []int
	for _, i := range arr {
		if i >= 0 {
			newArr = append(newArr, i)
		} else {
			break
		}
	}
	return task165dRecursive(newArr, len(newArr), &destArr)
}

// рекурсивна функція для виведення добутку
func task165dRecursive(arr []int, customLen int, destinationArr *[]int) []int {
	sum := 1
	for _, i := range arr[:customLen] {
		sum *= i
	}
	*destinationArr = append(*destinationArr, sum)
	if customLen > 1 {
		task165dRecursive(arr, customLen-1, destinationArr)
	}
	return *destinationArr
}
