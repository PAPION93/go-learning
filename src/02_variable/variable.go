package main

import (
	"fmt"
	"math"
	_ "time" // 사용하지않는 패키지 에러방지
	"unicode/utf8"
	"unsafe"
)

func main() {

	var a int
	a = 10

	var b float32
	b = math.Pi

	var b2 float64
	b2 = math.Pi

	var c string
	c = "Hello"

	fmt.Println(a, b, b2, c)

	// 빠른 선언
	age := 10
	name := "jun"
	fmt.Println(age, name)

	var f = 1
	fmt.Println(f)

	// 다수개 선언
	var x, y int = 31, 50
	age, name = 3, "gopher"
	fmt.Println(x, y, age, name)

	x, y = 31, 51
	fmt.Println(x, y)

	// 사용하지 않는 변수 컴파일 에러 방지
	var _ int = a

	// 변수의 크기 구하기
	var num int
	fmt.Println(unsafe.Sizeof(num))

	// String
	var s1 string = `Hello
World`
	var s2 string = "한글"

	// String 길이 구하기
	fmt.Println(s1)
	fmt.Println(len(s1)) // 11
	fmt.Println(len(s2)) // 6
	fmt.Println(utf8.RuneCountInString(s2)) // 2

	// 문자열 연산
	fmt.Println(s1 == s2)
	fmt.Println(s1 + s2)
	fmt.Println(s1 + s2 + "반갑습니다")
	fmt.Printf("%c\n", s1[0]) // H

}