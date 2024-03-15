package main

type shape interface {
	getArea() float64
}

type triangule struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangule{base: 10, height: 10}
	s := square{sideLength: 10}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	println(s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangule) getArea() float64 {
	return 0.5 * t.base * t.height
}
