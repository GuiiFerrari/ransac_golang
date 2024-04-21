package geometry

type Vector3D struct {
	Origin *Point3D
	Versor *Point3D
}

func CrossProduct3D(vector_a *Vector3D, vector_b *Vector3D) *Point3D {
	return &Point3D{
		X: vector_a.Versor.Y*vector_b.Versor.Z - vector_a.Versor.Z*vector_b.Versor.Y,
		Y: vector_a.Versor.Z*vector_b.Versor.X - vector_a.Versor.X*vector_b.Versor.Z,
		Z: vector_a.Versor.X*vector_b.Versor.Y - vector_a.Versor.Y*vector_b.Versor.X,
	}
}

func NormVector3D(point_a *Point3D) *Point3D {
	mod := point_a.Modulus()
	return &Point3D{X: point_a.X / mod, Y: point_a.Y / mod, Z: point_a.Z / mod}
}

func AngleBetweenVectors3D(point_a *Point3D, point_b *Point3D) float64 {
	mod_a := point_a.Modulus()
	mod_b := point_b.Modulus()
	return InnerProduct3d(*point_a, *point_b) / (mod_a * mod_b)
}

func (l *Vector3D) PointAt(t float64) *Point3D {
	return &Point3D{
		X: l.Origin.X + t*l.Versor.X,
		Y: l.Origin.Y + t*l.Versor.Y,
		Z: l.Origin.Z + t*l.Versor.Z}
}
