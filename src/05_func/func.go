package main

func main() {
	//변수 add 에 익명함수 할당
	add := func(i int, j int) int {
		return i + j
	}

	// add 함수 전달
	r1 := calc(add, 10, 20)
	println(r1)

	// 직접 첫번째 파라미터에 익명함수를 정의함
	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)

	// Closure
	next := nextValue()
	println(next()) // 1
	println(next()) // 2
	println(next()) // 3

	anotherNext := nextValue()
	println(anotherNext()) // 1 다시 시작
	println(anotherNext()) // 2

}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
