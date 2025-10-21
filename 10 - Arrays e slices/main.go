package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("09 - Arrays e slices")

	//Array
	fmt.Println("------------------ antes de popular ------------")
	var array1 [10]int
	fmt.Println(array1)
	fmt.Println("------------------ depois de popular ------------")
	array1[0] = 10
	fmt.Println(array1)
	fmt.Println("-----------------Passando por inferencia-------------------")
	array2 := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array2)
	fmt.Println("---------------------------------")
	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)
	fmt.Println("---------------------------------")

	//slice
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(slice)
	slice = append(slice, 22)
	fmt.Println(slice)
	slice2 := array2[1:3]
	fmt.Println(slice2)
	fmt.Println("---------------------------------")
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array1))
	fmt.Println("---------------------------------")

	//Arrays internos
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
