package main

import (
	"fmt"
)

func main() {
	// arr := [5]int{1, 2, 3, 4, 5}
	// slice := arr[3:4]
	// arr[4] = 22
	// fmt.Println(slice)

	// var arr [10]int
	// fmt.Println(arr)
	// arr[0] = 22
	// arr[9] = -23151
	// fmt.Printf("Vash massiv %d, dlina massiva %d", arr, len(arr))

	// arr := [...]int{1, 2, 34}
	// fmt.Printf("Vash massiv %d, dlina massiva %d", arr, len(arr))
	// var arr2 [10][10]string
	// arr2[0] = [10]string{"1", "2", "1", "2", "1", "2", "1", "2", "1", "2"}
	// arr2[1] = [10]string{"1", "2", "1", "2", "1", "2", "1", "2", "1", "2"}
	// arr2[2] = [10]string{"1", "2", "1", "2", "1", "2", "1", "2", "1", "2"}
	// arr2[3] = [10]string{"1", "2", "1", "2", "1", "2", "1", "2", "1", "2"}
	// arr2[4] = [10]string{"1", "2", "1", "2", "1", "2", "1", "2", "1", "2"}
	// arr2[5] = [10]string{"1", "2", "1", "2", "1", "2", "1", "2", "1", "2"}
	// arr2[6] = [10]string{"1", "2", "1", "2", "1", "2", "1", "2", "1", "2"}
	// arr2[7] = [10]string{"1", "2", "1", "2", "1", "2", "1", "2", "1", "2"}
	// arr2[8] = [10]string{"1", "2", "1", "2", "1", "2", "1", "2", "1", "2"}
	// arr2[9] = [10]string{"1", "2", "1", "2", "1", "2", "1", "2", "1", "2"}
	// fmt.Println(arr2)

	// a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// s1 := a[3:8]
	// s2 := s1[0:2]
	// fmt.Println(a, s1, s2)

	// a[4] = 22
	// fmt.Println(a, s1, s2)

	// fmt.Println(len(s2))
	// fmt.Println(cap(s2))
	// s2[0] = 9
	// fmt.Println(a, s1, s2)
	// s2 = append(s2, 11, 33, 44, 55, 66, 77, 88, 99, 111)

	// fmt.Println(a, s1, s2)
	// fmt.Println(len(s2))
	// fmt.Println(cap(s2))
	// s2 = append(s2, 11, 33, 44, 55, 66, 77, 88, 99, 111, 11, 33, 44, 55, 66, 77, 88, 99, 111, 222, 333)

	// fmt.Println(a, s1, s2)
	// fmt.Println(len(s2))
	// fmt.Println(cap(s2))

	// sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// sl2 := append(sl[:3], sl[4:]...)
	// fmt.Println(sl2)
	// fmt.Println(cap(sl2))

	// sl := make([]int, 0, 0)

	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, v := range sl {
		fmt.Printf("Индекс: %d, а его значение: %d \n", i, v)
	}

}
