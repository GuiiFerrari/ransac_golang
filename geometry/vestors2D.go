package geometry

type Vector2D struct {
	Origin *Point2D
	Versor *Point2D
}

func (l *Vector2D) PointAt(t float64) Point2D {
	return Point2D{
		X: l.Origin.X + t*l.Versor.X,
		Y: l.Origin.Y + t*l.Versor.Y}
}

func AngleBetweenVectors2d(vector_a *Vector2D, vector_b *Vector2D) float64 {
	mod_a := vector_a.Versor.Modulus()
	mod_b := vector_b.Versor.Modulus()
	return InnerProduct2d(*vector_a.Versor, *vector_b.Versor) / (mod_a * mod_b)
}
