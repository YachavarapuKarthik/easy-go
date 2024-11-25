package main

type Shape struct {
	length  int
	breadth int
}

type Rect interface {
	area() int
}

// using interface

func (s Shape) area() int {
	return s.length * s.breadth
}

func imp() int {
	c:= Shape{length: 10, breadth:20 }
	return c.area()
}
