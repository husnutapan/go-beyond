package main

import "fmt"

type Coordinate struct {
	xCoordinate int
	yCoordinate int
}

func Compare(c Coordinate) bool {
	return c.xCoordinate > c.yCoordinate
}

func IsAboveFunc(c Coordinate, y int) bool {
	return c.yCoordinate > y
}

func (c *Coordinate) ChangeValue(s1, s2 int) {
	c.xCoordinate = c.xCoordinate + s1
	c.yCoordinate = c.yCoordinate + s2

}

func main() {
	c := Coordinate{100, 200}

	fmt.Println("Result: ", Compare(c))

	fmt.Println("Result: ", IsAboveFunc(c, 150))

	c.ChangeValue(10, 20)
	fmt.Println("Result (changed Value): ", c.xCoordinate, c.yCoordinate)
}
