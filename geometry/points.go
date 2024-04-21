package geometry

import "math"

type Point2D struct {
	X float64
	Y float64
}

func (p Point2D) Sum(point Point2D) Point2D {
	return Point2D{X: p.X + point.X, Y: p.Y + point.Y}
}

func (p Point2D) Multiply(scalar float64) Point2D {
	return Point2D{X: p.X * scalar, Y: p.Y * scalar}
}

func (p Point2D) Modulus() float64 {
	return math.Sqrt(InnerProduct2d(p, p))
}

func InnerProduct2d(point_a Point2D, point_b Point2D) float64 {
	return point_a.X*point_b.X + point_a.Y*point_b.Y
}

type Point3D struct {
	X float64
	Y float64
	Z float64
}

func (p Point3D) Sum(point Point3D) Point3D {
	return Point3D{X: p.X + point.X, Y: p.Y + point.Y, Z: p.Z + point.Z}
}

func (p Point3D) Multiply(scalar float64) Point3D {
	return Point3D{X: p.X * scalar, Y: p.Y * scalar, Z: p.Z * scalar}
}

func (p Point3D) Modulus() float64 {
	return math.Sqrt(InnerProduct3d(p, p))
}

func InnerProduct3d(point_a Point3D, point_b Point3D) float64 {
	return point_a.X*point_b.X + point_a.Y*point_b.Y + point_a.Z*point_b.Z
}
