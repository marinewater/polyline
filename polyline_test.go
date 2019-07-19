package polyline_test

import (
	"github.com/marinewater/polyline"
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {

	encodeTests := []struct {
		name      string
		Points    []polyline.Point
		Precision uint32
		Expected  string
	}{
		{
			name:      "empty string (precision: 5)",
			Points:    []polyline.Point{},
			Precision: 5,
			Expected:  "",
		},
		{
			name: "single point (precision: 5)",
			Points: []polyline.Point{
				{
					Latitude:  -79.448,
					Longitude: -179.9832104,
				},
			},
			Precision: 5,
			Expected:  "~d|cN`~oia@",
		},
		{
			name: "multiple points (precision: 5)",
			Points: []polyline.Point{
				{
					Latitude:  38.5,
					Longitude: -120.2,
				},
				{
					Latitude:  40.7,
					Longitude: -120.95,
				},
				{
					Latitude:  43.252,
					Longitude: -126.453,
				},
			},
			Precision: 5,
			Expected:  "_p~iF~ps|U_ulLnnqC_mqNvxq`@",
		},
		{
			name:      "empty string (precision: 6)",
			Points:    []polyline.Point{},
			Precision: 6,
			Expected:  "",
		},
		{
			name: "multiple points (precision: 6)",
			Points: []polyline.Point{
				{
					Latitude:  48.208771,
					Longitude: 16.372572,
				},
				{
					Latitude:  48.210133,
					Longitude: 16.374164,
				},
				{
					Latitude:  48.210495,
					Longitude: 16.373436,
				},
			},
			Precision: 6,
			Expected:  "ewl}zAwthf^ctAobBsUnl@",
		},
		{
			name: "multiple points with minimal distance (precision: 6)",
			Points: []polyline.Point{
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
				{
					Latitude:  -37.472687,
					Longitude: -72.357526,
				},
				{
					Latitude:  -37.472165,
					Longitude: -72.357484,
				},
				{
					Latitude:  -37.472273,
					Longitude: -72.355672,
				},
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
			},
			Precision: 6,
			Expected:  "pfdnfAjic_iCsK~}Es_@sAvEgpBne@cjB",
		},
		{
			name: "same point twice (precision: 5)",
			Points: []polyline.Point{
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
			},
			Precision: 5,
			Expected:  "p|ucFfsrxL??",
		},
		{
			name: "same point twice (precision: 6)",
			Points: []polyline.Point{
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
			},
			Precision: 6,
			Expected:  "pfdnfAjic_iC??",
		},
	}

	for _, tt := range encodeTests {

		t.Run(tt.name, func(t *testing.T) {
			if result := polyline.Encode(tt.Points, tt.Precision); result != tt.Expected {
				t.Errorf("Expected: \"%s\", got: \"%s\"", tt.Expected, result)
			}
		})

	}

}

func TestEncode5(t *testing.T) {

	encodeTests := []struct {
		name     string
		Points   []polyline.Point
		Expected string
	}{
		{
			name:     "empty string",
			Points:   []polyline.Point{},
			Expected: "",
		},
		{
			name: "single point",
			Points: []polyline.Point{
				{
					Latitude:  -79.448,
					Longitude: -179.9832104,
				},
			},
			Expected: "~d|cN`~oia@",
		},
		{
			name: "multiple points",
			Points: []polyline.Point{
				{
					Latitude:  38.5,
					Longitude: -120.2,
				},
				{
					Latitude:  40.7,
					Longitude: -120.95,
				},
				{
					Latitude:  43.252,
					Longitude: -126.453,
				},
			},
			Expected: "_p~iF~ps|U_ulLnnqC_mqNvxq`@",
		},
		{
			name: "same point twice",
			Points: []polyline.Point{
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
			},
			Expected: "p|ucFfsrxL??",
		},
	}

	for _, tt := range encodeTests {

		t.Run(tt.name, func(t *testing.T) {
			if result := polyline.Encode5(tt.Points); result != tt.Expected {
				t.Errorf("Expected: %s, got: %s", tt.Expected, result)
			}
		})

	}

}

func TestEncode6(t *testing.T) {

	encodeTests := []struct {
		name     string
		Points   []polyline.Point
		Expected string
	}{
		{
			name:     "empty string",
			Points:   []polyline.Point{},
			Expected: "",
		},
		{
			name: "multiple points",
			Points: []polyline.Point{
				{
					Latitude:  48.208771,
					Longitude: 16.372572,
				},
				{
					Latitude:  48.210133,
					Longitude: 16.374164,
				},
				{
					Latitude:  48.210495,
					Longitude: 16.373436,
				},
			},
			Expected: "ewl}zAwthf^ctAobBsUnl@",
		},
		{
			name: "multiple points with minimal distance",
			Points: []polyline.Point{
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
				{
					Latitude:  -37.472687,
					Longitude: -72.357526,
				},
				{
					Latitude:  -37.472165,
					Longitude: -72.357484,
				},
				{
					Latitude:  -37.472273,
					Longitude: -72.355672,
				},
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
			},
			Expected: "pfdnfAjic_iCsK~}Es_@sAvEgpBne@cjB",
		},
		{
			name: "same point twice",
			Points: []polyline.Point{
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
			},
			Expected: "pfdnfAjic_iC??",
		},
	}

	for _, tt := range encodeTests {

		t.Run(tt.name, func(t *testing.T) {
			if result := polyline.Encode6(tt.Points); result != tt.Expected {
				t.Errorf("Expected: %s, got: %s", tt.Expected, result)
			}
		})

	}

}

func TestDecode(t *testing.T) {

	decodeTests := []struct {
		name      string
		Polyline  string
		Precision uint32
		Expected  []polyline.Point
	}{
		{
			name:      "empty string (precision: 5)",
			Polyline:  "",
			Precision: 5,
			Expected:  nil,
		},
		{
			name:      "not a polyline (precision: 5)",
			Polyline:  "a",
			Precision: 5,
			Expected:  nil,
		},
		{
			name:      "single point (precision: 5)",
			Polyline:  "~d|cN`~oia@",
			Precision: 5,
			Expected: []polyline.Point{
				{
					Latitude:  -79.448,
					Longitude: -179.98321,
				},
			},
		},
		{
			name:      "multiple points (precision: 5)",
			Polyline:  "_p~iF~ps|U_ulLnnqC_mqNvxq`@",
			Precision: 5,
			Expected: []polyline.Point{
				{
					Latitude:  38.5,
					Longitude: -120.2,
				},
				{
					Latitude:  40.7,
					Longitude: -120.95,
				},
				{
					Latitude:  43.252,
					Longitude: -126.453,
				},
			},
		},
		{
			name:      "empty string (precision: 6)",
			Polyline:  "",
			Precision: 6,
			Expected:  nil,
		},
		{
			name:      "not a polyline (precision: 6)",
			Polyline:  "a",
			Precision: 6,
			Expected:  nil,
		},
		{
			name:      "multiple points (precision: 6)",
			Polyline:  "ewl}zAwthf^ctAobBsUnl@",
			Precision: 6,
			Expected: []polyline.Point{
				{
					Latitude:  48.208771,
					Longitude: 16.372572,
				},
				{
					Latitude:  48.210133,
					Longitude: 16.374164,
				},
				{
					Latitude:  48.210495,
					Longitude: 16.373436,
				},
			},
		},
		{
			name:      "multiple points with minimal distance (precision: 6)",
			Polyline:  "pfdnfAjic_iCsK~}Es_@sAvEgpBne@cjB",
			Precision: 6,
			Expected: []polyline.Point{
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
				{
					Latitude:  -37.472687,
					Longitude: -72.357526,
				},
				{
					Latitude:  -37.472165,
					Longitude: -72.357484,
				},
				{
					Latitude:  -37.472273,
					Longitude: -72.355672,
				},
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
			},
		},
		{
			name:      "same point twice (precision: 5)",
			Polyline:  "p|ucFfsrxL??",
			Precision: 5,
			Expected: []polyline.Point{
				{
					Latitude:  -37.47289,
					Longitude: -72.35396,
				},
				{
					Latitude:  -37.47289,
					Longitude: -72.35396,
				},
			},
		},
		{
			name:      "same point twice (precision: 6)",
			Polyline:  "pfdnfAjic_iC??",
			Precision: 6,
			Expected: []polyline.Point{
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
			},
		},
	}

	for _, tt := range decodeTests {

		t.Run(tt.name, func(t *testing.T) {
			if result := polyline.Decode(tt.Polyline, tt.Precision); !reflect.DeepEqual(result, tt.Expected) {
				t.Errorf("Expected: %v, got: %v", tt.Expected, result)
			}
		})

	}

}

func TestDecode5(t *testing.T) {

	decodeTests := []struct {
		name     string
		Polyline string
		Expected []polyline.Point
	}{
		{
			name:     "empty string",
			Polyline: "",
			Expected: nil,
		},
		{
			name:     "not a polyline",
			Polyline: "a",
			Expected: nil,
		},
		{
			name:     "single point",
			Polyline: "~d|cN`~oia@",
			Expected: []polyline.Point{
				{
					Latitude:  -79.448,
					Longitude: -179.98321,
				},
			},
		},
		{
			name:     "multiple points",
			Polyline: "_p~iF~ps|U_ulLnnqC_mqNvxq`@",
			Expected: []polyline.Point{
				{
					Latitude:  38.5,
					Longitude: -120.2,
				},
				{
					Latitude:  40.7,
					Longitude: -120.95,
				},
				{
					Latitude:  43.252,
					Longitude: -126.453,
				},
			},
		},
		{
			name:     "same point twice",
			Polyline: "p|ucFfsrxL??",
			Expected: []polyline.Point{
				{
					Latitude:  -37.47289,
					Longitude: -72.35396,
				},
				{
					Latitude:  -37.47289,
					Longitude: -72.35396,
				},
			},
		},
	}

	for _, tt := range decodeTests {

		t.Run(tt.name, func(t *testing.T) {
			if result := polyline.Decode5(tt.Polyline); !reflect.DeepEqual(result, tt.Expected) {
				t.Errorf("Expected: %v, got: %v", tt.Expected, result)
			}
		})

	}

}

func TestDecode6(t *testing.T) {

	decodeTests := []struct {
		name     string
		Polyline string
		Expected []polyline.Point
	}{
		{
			name:     "empty string",
			Polyline: "",
			Expected: nil,
		},
		{
			name:     "not a polyline",
			Polyline: "a",
			Expected: nil,
		},
		{
			name:     "multiple points",
			Polyline: "ewl}zAwthf^ctAobBsUnl@",
			Expected: []polyline.Point{
				{
					Latitude:  48.208771,
					Longitude: 16.372572,
				},
				{
					Latitude:  48.210133,
					Longitude: 16.374164,
				},
				{
					Latitude:  48.210495,
					Longitude: 16.373436,
				},
			},
		},
		{
			name:     "multiple points with minimal distance",
			Polyline: "pfdnfAjic_iCsK~}Es_@sAvEgpBne@cjB",
			Expected: []polyline.Point{
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
				{
					Latitude:  -37.472687,
					Longitude: -72.357526,
				},
				{
					Latitude:  -37.472165,
					Longitude: -72.357484,
				},
				{
					Latitude:  -37.472273,
					Longitude: -72.355672,
				},
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
			},
		},
		{
			name:     "same point twice",
			Polyline: "pfdnfAjic_iC??",
			Expected: []polyline.Point{
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
				{
					Latitude:  -37.472889,
					Longitude: -72.353958,
				},
			},
		},
	}

	for _, tt := range decodeTests {

		t.Run(tt.name, func(t *testing.T) {
			if result := polyline.Decode6(tt.Polyline); !reflect.DeepEqual(result, tt.Expected) {
				t.Errorf("Expected: %v, got: %v", tt.Expected, result)
			}
		})

	}

}

func BenchmarkEncode(b *testing.B) {

	points := []polyline.Point{
		{
			Latitude:  -37.472889,
			Longitude: -72.353958,
		},
		{
			Latitude:  -37.472687,
			Longitude: -72.357526,
		},
		{
			Latitude:  -37.472165,
			Longitude: -72.357484,
		},
		{
			Latitude:  -37.472273,
			Longitude: -72.355672,
		},
		{
			Latitude:  -37.472889,
			Longitude: -72.353958,
		},
	}

	for n := 0; n < b.N; n++ {
		polyline.Encode(points, 6)
	}

}

func BenchmarkEncode5(b *testing.B) {

	points := []polyline.Point{
		{
			Latitude:  -37.472889,
			Longitude: -72.353958,
		},
		{
			Latitude:  -37.472687,
			Longitude: -72.357526,
		},
		{
			Latitude:  -37.472165,
			Longitude: -72.357484,
		},
		{
			Latitude:  -37.472273,
			Longitude: -72.355672,
		},
		{
			Latitude:  -37.472889,
			Longitude: -72.353958,
		},
	}

	for n := 0; n < b.N; n++ {
		polyline.Encode5(points)
	}

}

func BenchmarkEncode6(b *testing.B) {

	points := []polyline.Point{
		{
			Latitude:  -37.472889,
			Longitude: -72.353958,
		},
		{
			Latitude:  -37.472687,
			Longitude: -72.357526,
		},
		{
			Latitude:  -37.472165,
			Longitude: -72.357484,
		},
		{
			Latitude:  -37.472273,
			Longitude: -72.355672,
		},
		{
			Latitude:  -37.472889,
			Longitude: -72.353958,
		},
	}

	for n := 0; n < b.N; n++ {
		polyline.Encode6(points)
	}

}

func BenchmarkDecode(b *testing.B) {

	for n := 0; n < b.N; n++ {
		polyline.Decode("pfdnfAjic_iCsK~}Es_@sAvEgpBne@cjB", 6)
	}

}

func BenchmarkDecode5(b *testing.B) {

	for n := 0; n < b.N; n++ {
		polyline.Decode5("p|ucFfsrxLg@hUgBGTiJzBuI")
	}

}

func BenchmarkDecode6(b *testing.B) {

	for n := 0; n < b.N; n++ {
		polyline.Decode6("pfdnfAjic_iCsK~}Es_@sAvEgpBne@cjB")
	}

}
