package locationutils

type Latitude float64

func (lat *Latitude) IsValid() bool {
	return lat != nil && -90 <= *lat && *lat <= 90
}

type Longitude float64

func (lng *Longitude) IsValid() bool {
	return lng != nil && -180 <= *lng && *lng <= 180
}

type Point2D struct {
	Latitude  Latitude
	Longitude Longitude
}

func (p Point2D) IsValid() bool {
	return p.Latitude.IsValid() && p.Longitude.IsValid()
}

func (p Point2D) IsZero() bool {
	return p.Latitude == 0 && p.Longitude == 0
}

type BoundingBox struct {
	A Point2D
	B Point2D
}

func (b BoundingBox) Contains(p Point2D) bool {
	return b.A.Latitude <= p.Latitude && p.Latitude <= b.B.Latitude &&
		b.A.Longitude <= p.Longitude && p.Longitude <= b.B.Longitude
}
