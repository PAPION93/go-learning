package main

import (
	"fmt"
	"strings"
)

// 기본
func multiply(a, b int) int {
	return a * b
}

// 리턴 참고
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 파라미터 참고
func repeatMe(workds ...string) {
	fmt.Println(workds)
}

func main() {
	fmt.Println(multiply(2, 2))

	totalLengtht, upperName := lenAndUpper("jun")
	fmt.Println(totalLengtht, upperName)

	repeatMe("Jin", "Jun", "Hee")

}
