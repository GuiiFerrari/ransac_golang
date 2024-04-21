package ransac

import (
	"math/rand"
	"pransac/geometry"
)

func RandomSample(size int) []int {
	sample := make([]int, 2)
	first_number := rand.Intn(size)
	second_number := rand.Intn(size)
	for second_number == first_number {
		second_number = rand.Intn(size)
	}
	sample[0] = first_number
	sample[1] = second_number
	return sample
}

func ConstructVector(point_a geometry.Point3D, point_b geometry.Point3D) geometry.Vector3D {
	versor := geometry.Point3D{X: point_b.X - point_a.X, Y: point_b.Y - point_a.Y, Z: point_b.Z - point_a.Z}
	mod := versor.Modulus()
	versor.X = versor.X / mod
	versor.Y = versor.Y / mod
	versor.Z = versor.Z / mod
	return geometry.Vector3D{Origin: &point_a, Versor: &versor}
}

type Ransac struct {
	MinimumSamples int
	MaxDistance    float64
	Iterations     int
}

func (r *Ransac) Fit(points geometry.PointCloud) geometry.Vector3D {
	var best_model geometry.Vector3D
	best_score := 0
	for i := 0; i < r.Iterations; i++ {
		sample := RandomSample(len(points.Points))
		point_a := points.GetPoint(sample[0])
		point_b := points.GetPoint(sample[1])
		model := ConstructVector(point_a, point_b)
		score := 0
		for j := 0; j < len(points.Points); j++ {
			point := points.GetPoint(j)
			distance := geometry.DistancePointLine(&point, &model)
			if distance < r.MaxDistance {
				score++
			}
		}
		if score > best_score {
			best_model = model
			best_score = score
		}
	}
	return best_model
}
