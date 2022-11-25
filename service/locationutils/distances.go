package locationutils

import "math"

const (
	EarthRadiusKm = 6377.830272

	metersPerUnit = 1105.74
	latLngChange  = 0.01
)

// Squaring takes a point as the center point of a circle with range = `distance`.
// It returns the boundaries of the square that inscribes such circle.
func Squaring(point Point2D, distance float64) BoundingBox {
	offset := (distance * 1000 / metersPerUnit) * latLngChange

	return BoundingBox{
		A: Point2D{
			Latitude:  point.Latitude - Latitude(offset),
			Longitude: point.Longitude - Longitude(offset),
		},
		B: Point2D{
			Latitude:  point.Latitude + Latitude(offset),
			Longitude: point.Longitude + Longitude(offset),
		},
	}
}

// IsPointWithinRadius checks if two points are within `radius` distance from
// each other
func IsPointWithinRadius(p1 Point2D, p2 Point2D, radius float64) bool {
	distance := HaversineDistance(p1, p2)
	return distance <= radius
}

// HaversineDistance implementation of the haversine formula to get the distance between two points
func HaversineDistance(p1 Point2D, p2 Point2D) float64 {
	lat1 := toRad(float64(p1.Latitude))
	lat2 := toRad(float64(p2.Latitude))
	latDisplacement := toRad(float64(p2.Latitude - p1.Latitude))
	lonDisplacement := toRad(float64(p2.Longitude - p1.Longitude))

	a := math.Sin(latDisplacement/2)*math.Sin(latDisplacement/2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Sin(lonDisplacement/2)*math.Sin(lonDisplacement/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return c * EarthRadiusKm
}

// toRad converts degree to rad
func toRad(angle float64) float64 {
	return angle * math.Pi / 180
}
