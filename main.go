package main

func main() {}

// RectangleArea calculates the area of a rectangle.
func RectangleArea(length, width int) int {
	return length * width
}

// RectanglePerimeter calculates the perimeter of a rectangle.
func RectanglePerimeter(length, width int) int {
	return 2 * (length + width)
}

// SquareArea calculates the area of a square.
func SquareArea(side int) int {
	return side * side
}

// SquarePerimeter calculates the perimeter of a square.
func SquarePerimeter(side int) int {
	return 4 * side
}
