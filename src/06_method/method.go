package main

//Rect - struct 정의
type Rect struct {
	width, height int
}

//Rect의 area() 메소드
func (r Rect) area() int {
	r.width++
	return r.width * r.height
}

// 포인터 Receiver
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area() //메서드 호출
	println(area)

	aread := rect.area() //메서드 호출
	println(aread)

	rectb := Rect{10, 20}
	areab := rectb.area2() //메서드 호출
	println(rectb.width, areab) // 11 220 출력

	areac := rectb.area2() //메서드 호출
	println(rectb.width, areac) // 12 240 출력
}
