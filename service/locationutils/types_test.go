package locationutils

import "testing"

func TestLatitude_IsValid(t *testing.T) {
	t.Run("Ok", func(t *testing.T) {
		lat := Latitude(12.3456)
		if !lat.IsValid() {
			t.Fatal("Valid latitude rejected: ", lat)
		}
	})

	t.Run("TooHigh", func(t *testing.T) {
		lat := Latitude(91)
		if lat.IsValid() {
			t.Fatal("Too high latitude accepted: ", lat)
		}
	})

	t.Run("TooLow", func(t *testing.T) {
		lat := Latitude(-91)
		if lat.IsValid() {
			t.Fatal("Too low latitude accepted: ", lat)
		}
	})
}

func TestLongitude_IsValid(t *testing.T) {
	t.Run("Ok", func(t *testing.T) {
		lng := Longitude(12.3456)
		if !lng.IsValid() {
			t.Fatal("Valid longitude rejected: ", lng)
		}
	})

	t.Run("TooHigh", func(t *testing.T) {
		lng := Longitude(181)
		if lng.IsValid() {
			t.Fatal("Too high longitude accepted: ", lng)
		}
	})

	t.Run("TooLow", func(t *testing.T) {
		lng := Longitude(-181)
		if lng.IsValid() {
			t.Fatal("Too low longitude accepted: ", lng)
		}
	})
}

func TestPoint2D_IsZero(t *testing.T) {
	t.Run("Ok", func(t *testing.T) {
		p := Point2D{0, 0}
		if !p.IsZero() {
			t.Fatal("Zero point not recognized")
		}
	})

	t.Run("NotZero", func(t *testing.T) {
		p := Point2D{1, 2}
		if p.IsZero() {
			t.Fatal("Zero point when no zero: ", p)
		}
	})
}

func TestPoint2D_IsValid(t *testing.T) {
	t.Run("OkZero", func(t *testing.T) {
		p := Point2D{0, 0}
		if !p.IsValid() {
			t.Fatal("Zero point not valid")
		}
	})

	t.Run("Ok", func(t *testing.T) {
		p := Point2D{12, 34}
		if !p.IsValid() {
			t.Fatal("Valid point rejected")
		}
	})

	t.Run("Invalid_Partial", func(t *testing.T) {
		p := Point2D{12, 200}
		if p.IsValid() {
			t.Fatal("Invalid point accepted")
		}
	})

	t.Run("Invalid_Full", func(t *testing.T) {
		p := Point2D{300, 200}
		if p.IsValid() {
			t.Fatal("Invalid point accepted")
		}
	})
}

func TestBoundingBox_Contains(t *testing.T) {
	t.Run("Ok", func(t *testing.T) {
		p := Point2D{
			Latitude:  12,
			Longitude: 34,
		}
		box := BoundingBox{
			A: Point2D{Latitude: 10, Longitude: 20},
			B: Point2D{Latitude: 20, Longitude: 40},
		}
		if !box.Contains(p) {
			t.Fatal("Box contains p, but .Contains thinks otherwise")
		}
	})
	t.Run("Ok_borderline", func(t *testing.T) {
		p := Point2D{
			Latitude:  12,
			Longitude: 34,
		}
		box := BoundingBox{
			A: Point2D{Latitude: 12, Longitude: 20},
			B: Point2D{Latitude: 20, Longitude: 34},
		}
		if !box.Contains(p) {
			t.Fatal("Box contains p, but .Contains thinks otherwise")
		}
	})
	t.Run("Outside_latitude_high", func(t *testing.T) {
		p := Point2D{
			Latitude:  12,
			Longitude: 34,
		}
		box := BoundingBox{
			A: Point2D{Latitude: 5, Longitude: 20},
			B: Point2D{Latitude: 10, Longitude: 40},
		}
		if box.Contains(p) {
			t.Fatal("Box does not contains p, but .Contains thinks otherwise")
		}
	})
	t.Run("Outside_latitude_low", func(t *testing.T) {
		p := Point2D{
			Latitude:  12,
			Longitude: 34,
		}
		box := BoundingBox{
			A: Point2D{Latitude: 20, Longitude: 20},
			B: Point2D{Latitude: 30, Longitude: 40},
		}
		if box.Contains(p) {
			t.Fatal("Box does not contains p, but .Contains thinks otherwise")
		}
	})
	t.Run("Outside_longitude_high", func(t *testing.T) {
		p := Point2D{
			Latitude:  12,
			Longitude: 34,
		}
		box := BoundingBox{
			A: Point2D{Latitude: 10, Longitude: 20},
			B: Point2D{Latitude: 20, Longitude: 30},
		}
		if box.Contains(p) {
			t.Fatal("Box does not contains p, but .Contains thinks otherwise")
		}
	})
	t.Run("Outside_longitude_low", func(t *testing.T) {
		p := Point2D{
			Latitude:  12,
			Longitude: 34,
		}
		box := BoundingBox{
			A: Point2D{Latitude: 10, Longitude: 40},
			B: Point2D{Latitude: 20, Longitude: 50},
		}
		if box.Contains(p) {
			t.Fatal("Box does not contains p, but .Contains thinks otherwise")
		}
	})
}
