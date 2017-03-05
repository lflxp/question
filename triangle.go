package workspace
/*
Write a function that determines the type of a triangle. It should take the lengths of the triangle's three sides as input, and return whether the triangle is equilateral, isosceles or scalene. Also consider tests, documentation and/or errors.
*/
import (
	"fmt"
)
//To determine whether the triangle composition conditions
//1.The maximum edge is less than the sum of the other two sides
//2.Any side is greater than 0
func IsOk(a,b,c int) bool {
	var rs bool
	if a <= 0 || b <=0 || c <= 0 {
		fmt.Println("Less than or equal to 0")
		return false
	}
	max := a
	num := 1
	if b > max {
		max = b
		num = 2
	}
	if c > max {
		max = c
		num = 3
	}
	switch num {
	case 1:
		if a < b+c {
			rs = true
		} else {
			rs = false
		}
	case 2:
		if b < a+c {
			rs = true
		} else {
			rs = false
		}
	case 3:
		if c < a+b {
			rs = true
		} else {
			rs = false
		}
	}
	return rs
}
//Making use of the uniqueness of the dictionary to carry out the three side long and short repetition
func Triangles(a,b,c int) string {
	tmp := map[int]int{}
	//Determine whether you can form a triangle
	if IsOk(a,b,c) {
		tmp[a] = a
		tmp[b] = b
		tmp[c] = c
	}
	//Repeat the number of edges according to the length of the dictionary
	switch len(tmp) {
	case 1:
		return "equilateral"
	case 2:
		return "isosceles"
	case 3:
		return "scalene"
	default:
		return "illegal triangles"
	}
}