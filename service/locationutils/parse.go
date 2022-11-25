package locationutils

import (
	"errors"
	"strconv"
)

// ParseLatitude parses the string as latitude and checks the data validity (latitude range)
func ParseLatitude(str string) (Latitude, error) {
	latitude, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	} else if latitude < -90 || latitude > 90 {
		return 0, errors.New("invalid value range for latitude")
	}
	return Latitude(latitude), nil
}

// ParseLongitude parses the string as longitude and checks the data validity (longitude range)
func ParseLongitude(str string) (Longitude, error) {
	longitude, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	} else if longitude < -180 || longitude > 180 {
		return 0, errors.New("invalid value range for longitude")
	}
	return Longitude(longitude), nil
}

func (lat *Latitude) UnmarshalJSON(data []byte) error {
	latitude, err := ParseLatitude(string(data))
	if err != nil {
		return err
	}
	*lat = latitude
	return nil
}

func (lng *Longitude) UnmarshalJSON(data []byte) error {
	longitude, err := ParseLongitude(string(data))
	if err != nil {
		return err
	}
	*lng = longitude
	return nil
}

// ParseLatitudeToFloat parses the string as float64 and checks the data validity for latitude range
func ParseLatitudeToFloat(str string) (float64, error) {
	lat, err := ParseLatitude(str)
	return float64(lat), err
}

// ParseLongitudeToFloat parses the string as float64 and checks the data validity for longitude range
func ParseLongitudeToFloat(str string) (float64, error) {
	lng, err := ParseLongitude(str)
	return float64(lng), err
}
