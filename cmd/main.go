package main

import (
	"fmt"
	"pransac/geometry"
)

func main() {
	p1 := geometry.Point3D{X: 2.0, Y: 3.0, Z: 4.0}
	p2 := geometry.Point3D{X: 2.0, Y: 3.0, Z: 4.0}
	p3 := p1.Sum(p2)
	fmt.Println(p1)
	fmt.Println(p3)
	p1 = p1.Multiply(2)
	fmt.Println(p1)
}
