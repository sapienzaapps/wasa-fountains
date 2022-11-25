package locationutils

import (
	"encoding/json"
	"testing"
)

func TestParseLatitude(t *testing.T) {
	t.Run("PositiveOk", func(t *testing.T) {
		lat, err := ParseLatitude("12.3456")
		if err != nil {
			t.Fatal("Good latitude rejected with: ", err)
		}
		if lat != 12.3456 {
			t.Fatal("Latitude decode is incorrect, expected: 12.3456 found: ", lat)
		}
	})

	t.Run("NegativeOk", func(t *testing.T) {
		lat, err := ParseLatitude("-12.3456")
		if err != nil {
			t.Fatal("Good latitude rejected with: ", err)
		}
		if lat != -12.3456 {
			t.Fatal("Latitude decode is incorrect, expected: -12.3456 found: ", lat)
		}
	})

	t.Run("LimitsOk", func(t *testing.T) {
		lat, err := ParseLatitude("90")
		if err != nil {
			t.Fatal("Good latitude rejected with: ", err)
		}
		if lat != 90 {
			t.Fatal("Latitude decode is incorrect, expected: 90 found: ", lat)
		}

		lat, err = ParseLatitude("-90")
		if err != nil {
			t.Fatal("Good latitude rejected with: ", err)
		}
		if lat != -90 {
			t.Fatal("Latitude decode is incorrect, expected: -90 found: ", lat)
		}
	})

	t.Run("TooHigh", func(t *testing.T) {
		_, err := ParseLatitude("91")
		if err == nil {
			t.Fatal("Bad latitude (91) accepted")
		}
	})

	t.Run("TooLow", func(t *testing.T) {
		_, err := ParseLatitude("-91")
		if err == nil {
			t.Fatal("Bad latitude (-91) accepted")
		}
	})

	t.Run("Empty", func(t *testing.T) {
		_, err := ParseLatitude("")
		if err == nil {
			t.Fatal("Bad empty value accepted")
		}
	})

	t.Run("String", func(t *testing.T) {
		_, err := ParseLatitude("asdds12")
		if err == nil {
			t.Fatal("String value accepted")
		}
	})

	t.Run("InvalidNumber", func(t *testing.T) {
		_, err := ParseLatitude("12,34")
		if err == nil {
			t.Fatal("Invalid value accepted")
		}
	})
}

func TestParseLatitudeFromJSON(t *testing.T) {
	type teststruct struct {
		Latitude Latitude `json:"latitude"`
	}
	t.Run("PositiveOk", func(t *testing.T) {
		var obj teststruct
		err := json.Unmarshal([]byte(`{"latitude": 12.3456}`), &obj)
		switch {
		case err != nil:
			t.Fatal("Good latitude rejected with: ", err)
		case !obj.Latitude.IsValid():
			t.Fatal("Good latitude rejected as invalid")
		case obj.Latitude != 12.3456:
			t.Fatal("Latitude decode is incorrect, expected: 12.3456 found: ", obj.Latitude)
		}
	})

	t.Run("NegativeOk", func(t *testing.T) {
		var obj teststruct
		err := json.Unmarshal([]byte(`{"latitude": -12.3456}`), &obj)
		switch {
		case err != nil:
			t.Fatal("Good latitude rejected with: ", err)
		case !obj.Latitude.IsValid():
			t.Fatal("Good latitude rejected as invalid")
		case obj.Latitude != -12.3456:
			t.Fatal("Latitude decode is incorrect, expected: -12.3456 found: ", obj.Latitude)
		}
	})

	t.Run("LimitsOk", func(t *testing.T) {
		var obj teststruct
		err := json.Unmarshal([]byte(`{"latitude": 90}`), &obj)
		switch {
		case err != nil:
			t.Fatal("Good latitude rejected with: ", err)
		case !obj.Latitude.IsValid():
			t.Fatal("Good latitude rejected as invalid")
		case obj.Latitude != 90:
			t.Fatal("Latitude decode is incorrect, expected: 90 found: ", obj.Latitude)
		}

		err = json.Unmarshal([]byte(`{"latitude": -90}`), &obj)
		switch {
		case err != nil:
			t.Fatal("Good latitude rejected with: ", err)
		case !obj.Latitude.IsValid():
			t.Fatal("Good latitude rejected as invalid")
		case obj.Latitude != -90:
			t.Fatal("Latitude decode is incorrect, expected: -90 found: ", obj.Latitude)
		}
	})

	t.Run("TooHigh", func(t *testing.T) {
		var obj teststruct
		err := json.Unmarshal([]byte(`{"latitude": 91}`), &obj)
		if err == nil {
			t.Fatal("Bad latitude (91) accepted")
		}
	})

	t.Run("TooLow", func(t *testing.T) {
		var obj teststruct
		err := json.Unmarshal([]byte(`{"latitude": -91}`), &obj)
		if err == nil {
			t.Fatal("Bad latitude (-91) accepted")
		}
	})

	t.Run("String", func(t *testing.T) {
		var obj teststruct
		err := json.Unmarshal([]byte(`{"latitude": ""}`), &obj)
		if err == nil {
			t.Fatal("String value accepted")
		}
	})

	t.Run("InvalidNumber", func(t *testing.T) {
		var obj teststruct
		err := json.Unmarshal([]byte(`{"latitude": 12,34}`), &obj)
		if err == nil {
			t.Fatal("Invalid value accepted")
		}
	})
}

func TestParseLongitude(t *testing.T) {
	t.Run("PositiveOk", func(t *testing.T) {
		lng, err := ParseLongitude("12.3456")
		if err != nil {
			t.Fatal("Good longitude rejected with: ", err)
		}
		if lng != 12.3456 {
			t.Fatal("Longitude decode is incorrect, expected: 12.3456 found: ", lng)
		}
	})

	t.Run("NegativeOk", func(t *testing.T) {
		lng, err := ParseLongitude("-12.3456")
		if err != nil {
			t.Fatal("Good longitude rejected with: ", err)
		}
		if lng != -12.3456 {
			t.Fatal("Longitude decode is incorrect, expected: -12.3456 found: ", lng)
		}
	})

	t.Run("LimitsOk", func(t *testing.T) {
		lng, err := ParseLongitude("180")
		if err != nil {
			t.Fatal("Good longitude rejected with: ", err)
		}
		if lng != 180 {
			t.Fatal("Longitude decode is incorrect, expected: 180 found: ", lng)
		}

		lng, err = ParseLongitude("-180")
		if err != nil {
			t.Fatal("Good longitude rejected with: ", err)
		}
		if lng != -180 {
			t.Fatal("Longitude decode is incorrect, expected: -180 found: ", lng)
		}
	})

	t.Run("TooHigh", func(t *testing.T) {
		_, err := ParseLongitude("181")
		if err == nil {
			t.Fatal("Bad longitude (181) accepted")
		}
	})

	t.Run("TooLow", func(t *testing.T) {
		_, err := ParseLongitude("-181")
		if err == nil {
			t.Fatal("Bad longitude (-181) accepted")
		}
	})

	t.Run("Empty", func(t *testing.T) {
		_, err := ParseLongitude("")
		if err == nil {
			t.Fatal("Bad empty value accepted")
		}
	})

	t.Run("String", func(t *testing.T) {
		_, err := ParseLongitude("asdds12")
		if err == nil {
			t.Fatal("String value accepted")
		}
	})

	t.Run("InvalidNumber", func(t *testing.T) {
		_, err := ParseLongitude("12,34")
		if err == nil {
			t.Fatal("Invalid value accepted")
		}
	})
}

func TestParseLongitudeFromJSON(t *testing.T) {
	type teststruct struct {
		Longitude Longitude `json:"longitude"`
	}
	t.Run("PositiveOk", func(t *testing.T) {
		var obj teststruct
		err := json.Unmarshal([]byte(`{"longitude": 12.3456}`), &obj)
		switch {
		case err != nil:
			t.Fatal("Good longitude rejected with: ", err)
		case !obj.Longitude.IsValid():
			t.Fatal("Good longitude rejected as invalid")
		case obj.Longitude != 12.3456:
			t.Fatal("Latitude decode is incorrect, expected: 12.3456 found: ", obj.Longitude)
		}
	})

	t.Run("NegativeOk", func(t *testing.T) {
		var obj teststruct
		err := json.Unmarshal([]byte(`{"longitude": -12.3456}`), &obj)
		switch {
		case err != nil:
			t.Fatal("Good longitude rejected with: ", err)
		case !obj.Longitude.IsValid():
			t.Fatal("Good longitude rejected as invalid")
		case obj.Longitude != -12.3456:
			t.Fatal("Latitude decode is incorrect, expected: -12.3456 found: ", obj.Longitude)
		}
	})

	t.Run("LimitsOk", func(t *testing.T) {
		var obj teststruct
		err := json.Unmarshal([]byte(`{"longitude": 180}`), &obj)
		switch {
		case err != nil:
			t.Fatal("Good longitude rejected with: ", err)
		case !obj.Longitude.IsValid():
			t.Fatal("Good longitude rejected as invalid")
		case obj.Longitude != 180:
			t.Fatal("Latitude decode is incorrect, expected: 180 found: ", obj.Longitude)
		}

		err = json.Unmarshal([]byte(`{"longitude": -180}`), &obj)
		switch {
		case err != nil:
			t.Fatal("Good longitude rejected with: ", err)
		case !obj.Longitude.IsValid():
			t.Fatal("Good longitude rejected as invalid")
		case obj.Longitude != -180:
			t.Fatal("Latitude decode is incorrect, expected: -180 found: ", obj.Longitude)
		}
	})

	t.Run("TooHigh", func(t *testing.T) {
		var obj teststruct
		err := json.Unmarshal([]byte(`{"longitude": 181}`), &obj)
		if err == nil {
			t.Fatal("Bad longitude (181) accepted")
		}
	})

	t.Run("TooLow", func(t *testing.T) {
		var obj teststruct
		err := json.Unmarshal([]byte(`{"longitude": -181}`), &obj)
		if err == nil {
			t.Fatal("Bad longitude (-181) accepted")
		}
	})

	t.Run("String", func(t *testing.T) {
		var obj teststruct
		err := json.Unmarshal([]byte(`{"longitude": ""}`), &obj)
		if err == nil {
			t.Fatal("String value accepted")
		}
	})

	t.Run("InvalidNumber", func(t *testing.T) {
		var obj teststruct
		err := json.Unmarshal([]byte(`{"longitude": 12,34}`), &obj)
		if err == nil {
			t.Fatal("Invalid value accepted")
		}
	})
}
