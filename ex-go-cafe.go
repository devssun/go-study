package main

import "fmt"

func main() {
	// 1
	var cafeName string = "GoCafe"
	var americanoPrice int = 1500
	var lattePrice int = 1700
	var cafePhoneNumber string = "02-428-5544"
	var isOpenToday bool = true

	fmt.Println("카페 이름 : " + cafeName)
	fmt.Println("아메리카노 가격 : ", americanoPrice)
	fmt.Println("라떼 가격 : ", lattePrice)
	fmt.Println("카페 전화번호 : " + cafePhoneNumber)
	fmt.Println("카페 Open 여부 : ", isOpenToday)

	// 2
	var drinkNames = [2]string{"아메리카노", "라떼"}
	var drinkPrices = [2]int{americanoPrice, lattePrice}

	fmt.Println(drinkNames, drinkPrices)

	// 3
	amount := (americanoPrice * 3) + (lattePrice * 2)

	// 4
	isTodayMonday := true
	americanoOrder := 3 // 아메리카노 주문량
	latterOrder := 5    // 라떼 주문량
	totalPrice := 0     // 총 금액

	totalPrice = (americanoPrice * americanoOrder) + (lattePrice * latterOrder)
	// 총 금액이 5000 초과이면 10% 할인
	if totalPrice > 5000 {
		totalPrice = int(float32(totalPrice) * 0.9)
	}
	// 총 금액이 5000 초과이면서 월요일이면 추가 10%
	if amount > 5000 && isTodayMonday {
		totalPrice = int(float32(totalPrice) * 0.9)
	}
	// 같은 음료가 4잔 이상이면 추가 10%
	if americanoOrder >= 4 || latterOrder >= 4 {
		totalPrice = int(float32(totalPrice) * 0.9)
	}

	fmt.Println("최종 가격 : ", totalPrice)

	// 5 use loop
	// drinkNames2 := []string{"Americano", "Latte", "CafeMocha", "GreenTeaLatte", "HotChoco"}
	drinkPrices2 := []int{1500, 1700, 1900, 2300, 2000}

	totalPrice2 := 0

	// for 1
	for i := 0; i < len(drinkPrices2); i++ {
		totalPrice2 += drinkPrices2[i]
	}

	// for 2
	totalPrice2 = 0
	for _, i := range drinkPrices2 {
		totalPrice2 += i
	}

	fmt.Println("VIP 고객 주문 음료 총 금액 : ", totalPrice2)

	// 6
	drinkNames3 := []string{"Americano", "Latte", "CafeMocha", "GreenTeaLatte", "HotChoco", "ChamomileTea", "StrawberrySmoothie", "HotMilk"}
	drinkPrices3 := []int{1500, 1700, 1900, 2300, 2000, 1800, 3500, 1500}

	totalPrice3 := 0
	orderedDrinks := ""

	for i := 0; i < len(drinkNames3); i++ {
		// 2000원 이하 음료만 주문
		if drinkPrices3[i] > 2000 {
			continue
		}
		// 총 금액에 7000 이상이면 주문 끝
		if totalPrice3 >= 7000 {
			break
		}
		orderedDrinks += drinkNames3[i] + ", "
		totalPrice3 += drinkPrices3[i]
	}

	fmt.Println("손님이 주문한 최종 음료 : ", orderedDrinks)
	fmt.Println("손님이 주문한 최종 금액 : ", totalPrice3)

	// 7
	// use function - 카페에 전화를 건 손님에게 인사한 후 , 요일에 따라 마감시간을 안내하는 함수를 만들어 봅시다
	fmt.Println("//----------------------------------------\n")
	greeting()
	greeting2("수요일")
	fmt.Println("//----------------------------------------\n")
}

func greeting() {
	fmt.Println("여기는 GoCafe입니다\n마감시간을 알려드리겠습니다\n")
}

func greeting2(day string) {
	if day == "화요일" || day == "목요일" {
		fmt.Println("마감시간은 9시입니다")
	} else if day == "월요일" || day == "수요일" || day == "금요일" {
		fmt.Println("마감시간은 10시입니다")
	} else if day == "토요일" {
		fmt.Println("마감시간은 8시입니다")
	} else {
		fmt.Println("휴무입니다")
	}
}
