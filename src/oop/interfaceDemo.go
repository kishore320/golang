package main

import "fmt"

func main() {
	sq := Square{10}
	rect := Rectangle{5, 10}

	shapes := [2]Shape{&sq, &rect}
	for _, sh := range shapes {
		fmt.Println("Shape Name :", getName(sh))
		fmt.Println("Shape Area :", getArea(sh))
	}
}

// Interface is just a collection of method signatures
// To implement interface, just implement the methods defined in the interface
type Shape interface {
	area() int
	getName() string
}

type Square struct {
	side int
}

type Rectangle struct {
	length  int
	breadth int
}

// Square is implementing the area() method
func (sq *Square) area() int {
	return sq.side * sq.side
}

// Square is implementing the getName() method
func (sq *Square) getName() string {
	return "Square"
}

// Rectangle is implementing the area() method
func (rect *Rectangle) area() int {
	return rect.length * rect.breadth
}

// Rectangle is implementing the getName() method
func (rect *Rectangle) getName() string {
	return "Rectangle"
}

func getArea(shape Shape) int {
	return shape.area()
}

func getName(shape Shape) string {
	return shape.getName()
}
