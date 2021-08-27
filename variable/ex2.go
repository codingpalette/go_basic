package main

import "fmt"

func main()  {
	var minimumWage int = 10 // 변수 선언 및 초기화
	var workingHour int = 10 // 변수 선언 및 초기화

	// 변수 선언 및 초기화
	var income int = minimumWage + workingHour

	// 변수 출력
	fmt.Println(minimumWage, workingHour, income)
}
