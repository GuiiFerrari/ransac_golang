package geometry

type PointCloud struct {
	Points []Point3D
}

func (p *PointCloud) AddPoint(point Point3D) {
	p.Points = append(p.Points, point)
}

func (p *PointCloud) AddPoints(points []Point3D) {
	p.Points = append(p.Points, points...)
}

func (p *PointCloud) AddPointCloud(point_cloud PointCloud) {
	p.Points = append(p.Points, point_cloud.Points...)
}

func (p *PointCloud) RemovePoint(index int) {
	p.Points = append(p.Points[:index], p.Points[index+1:]...)
}

func (p *PointCloud) RemovePoints(indexes []int) {
	for i := 0; i < len(indexes); i++ {
		p.RemovePoint(indexes[i])
	}
}
func (p *PointCloud) GetPoint(index int) Point3D {
	return p.Points[index]
}
