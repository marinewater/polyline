package polyline_test

import (
	"github.com/marinewater/polyline"
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {

	encodeTests := []struct {
		Points    []polyline.Point
		Precision uint32
		Expected  string
	}{
		{
			Points:    []polyline.Point{},
			Precision: 5,
			Expected:  "",
		},
		{
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
			Points:    []polyline.Point{},
			Precision: 6,
			Expected:  "",
		},
		{
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
	}

	for _, tt := range encodeTests {

		if result := polyline.Encode(tt.Points, tt.Precision); result != tt.Expected {
			t.Errorf("Expected: %s, got: %s", tt.Expected, result)
		}
	}

}

func TestEncode5(t *testing.T) {

	encodeTests := []struct {
		Points   []polyline.Point
		Expected string
	}{
		{
			Points:   []polyline.Point{},
			Expected: "",
		},
		{
			Points: []polyline.Point{
				{
					Latitude:  -79.448,
					Longitude: -179.9832104,
				},
			},
			Expected: "~d|cN`~oia@",
		},
		{
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
	}

	for _, tt := range encodeTests {

		if result := polyline.Encode5(tt.Points); result != tt.Expected {
			t.Errorf("Expected: %s, got: %s", tt.Expected, result)
		}
	}

}

func TestEncode6(t *testing.T) {

	encodeTests := []struct {
		Points   []polyline.Point
		Expected string
	}{
		{
			Points:   []polyline.Point{},
			Expected: "",
		},
		{
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
	}

	for _, tt := range encodeTests {

		if result := polyline.Encode6(tt.Points); result != tt.Expected {
			t.Errorf("Expected: %s, got: %s", tt.Expected, result)
		}
	}

}

func TestDecode(t *testing.T) {

	decodeTests := []struct {
		Polyline  string
		Precision uint32
		Expected  []polyline.Point
	}{
		{
			Polyline:  "",
			Precision: 5,
			Expected:  []polyline.Point{},
		},
		{
			Polyline:  "a",
			Precision: 5,
			Expected:  []polyline.Point{},
		},
		{
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
			Polyline:  "",
			Precision: 6,
			Expected:  []polyline.Point{},
		},
		{
			Polyline:  "a",
			Precision: 6,
			Expected:  []polyline.Point{},
		},
		{
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
	}

	for _, tt := range decodeTests {

		if result := polyline.Decode(tt.Polyline, tt.Precision); !reflect.DeepEqual(result, tt.Expected) {
			t.Errorf("Expected: %v, got: %v", tt.Expected, result)
		}

	}

}

func TestDecode5(t *testing.T) {

	decodeTests := []struct {
		Polyline string
		Expected []polyline.Point
	}{
		{
			Polyline: "",
			Expected: []polyline.Point{},
		},
		{
			Polyline: "a",
			Expected: []polyline.Point{},
		},
		{
			Polyline: "~d|cN`~oia@",
			Expected: []polyline.Point{
				{
					Latitude:  -79.448,
					Longitude: -179.98321,
				},
			},
		},
		{
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
	}

	for _, tt := range decodeTests {

		if result := polyline.Decode5(tt.Polyline); !reflect.DeepEqual(result, tt.Expected) {
			t.Errorf("Expected: %v, got: %v", tt.Expected, result)
		}

	}

}

func TestDecode6(t *testing.T) {

	decodeTests := []struct {
		Polyline string
		Expected []polyline.Point
	}{
		{
			Polyline: "",
			Expected: []polyline.Point{},
		},
		{
			Polyline: "a",
			Expected: []polyline.Point{},
		},
		{
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
	}

	for _, tt := range decodeTests {

		if result := polyline.Decode6(tt.Polyline); !reflect.DeepEqual(result, tt.Expected) {
			t.Errorf("Expected: %v, got: %v", tt.Expected, result)
		}

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
