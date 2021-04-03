package main

import "fmt"

func main() {

	// // step := 2
	// for step := 8; step > 0; step-- {
	// 	fmt.Println(step)
	// }

	str := "" // "Hello"
	//空串不执行循环
	for k, v := range str {
		fmt.Println(k, "=", uint8(v))
	}

	var myArr [10]string
	myArr[0] = "hello"
	myArr[1] = "hellotffg"

	for k, v := range myArr {
		fmt.Println(k, "=", v)
	}

	arr := [...]int{1, 4, 7, 9}
	fmt.Println("Int Arrary length=", len(arr))
	bb := arr[1:2:2]
	fmt.Println("Slice of Int Arrary:", bb)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum1 := 0
	cur := 0
	for cur < 10 {
		sum1 += cur
		cur++
	}
	fmt.Println(sum1)

}
