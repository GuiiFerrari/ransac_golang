package geometry

func DistancePointLine(point *Point3D, line *Vector3D) float64 {
	// First, calculate the vector from a point of the line to the point of interest

	auxiliar_point := Point3D{
		X: point.X - line.Origin.X,
		Y: point.Y - line.Origin.Y,
		Z: point.Z - line.Origin.Z,
	}
	cross_product := CrossProduct3D(line, &Vector3D{Origin: &auxiliar_point, Versor: line.Versor})
	return cross_product.Modulus()
}
