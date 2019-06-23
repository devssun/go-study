package main

import "fmt"

func main() {
	// slice 예제
	var slice []int

	// 선언 1
	// slice에 값을 넣어주기 위해 make() 를 사용한다
	// 이렇게 작성하면 0으로 초기화 된다
	slice = make([]int, 5)

	// 선언2
	// 배열과 같은 방식으로 {}를 사용하여 초기화할 수 있음
	slice2 := []int{1, 2, 3, 4, 5}

	fmt.Println("slice init")

	fmt.Println("Hello, Slice", slice)
	fmt.Println("Hello, Slice 2 ", slice2)

	fmt.Println("---------------------------")
	fmt.Println("slice = reference type")

	// slice = reference type
	// 두 변수는 내부적으로 같은 객체를 참조하고 있음
	slice = slice2
	slice[2] = 50

	fmt.Println("Slice : ", slice)
	fmt.Println("Slice2 : ", slice2)

	fmt.Println("---------------------------")
	fmt.Println("copy 예제")

	// 값만 복사하기
	slice3 := make([]int, 5)
	copy(slice3, slice2)
	slice3[2] = 0

	fmt.Println("Slice2 : ", slice2)
	fmt.Println("Slice3 : ", slice3)

	fmt.Println(slice3[2:3])
}
