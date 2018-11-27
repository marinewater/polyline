// Package polyline provides functions to encode/decode polylines in "Encoded Polyline Algorithm Format"
package polyline

import (
	"math"
)

// Point is a coordinate of a point on the polyline
type Point struct {
	Latitude  float64
	Longitude float64
}

// Encode encodes coordinates to the "Encoded Polyline Algorithm Format"
// More info: https://developers.google.com/maps/documentation/utilities/polylinealgorithm
// points: points of the polyline
// precision: usually 5 or 6; Google's original algorithm uses 5 digits of decimal precision,
// which is accurate to about a meter. A precision of 6 gives you an accuracy of about 10cm
// more info: https://mapzen.com/blog/polyline-precision/
func Encode(points []Point, precision uint32) string {

	encoded := ""

	latitude := 0.0
	longitude := 0.0

	for _, point := range points {
		polyLatitude := encodeElement(point.Latitude-latitude, precision)
		encoded += polyLatitude

		polyLongitude := encodeElement(point.Longitude-longitude, precision)
		encoded += polyLongitude

		// to conserve space, points only include the offset from the previous point
		latitude = point.Latitude
		longitude = point.Longitude

	}

	return encoded

}

// Encode5 is a short call for Encode with precision set to 5
// Accuracy is about one meter
func Encode5(points []Point) string {
	return Encode(points, 5)
}

// Encode6 is a short call for Encode with precision set to 6
// Accuracy is about ten centimeters
func Encode6(points []Point) string {
	return Encode(points, 6)
}

// Decode decodes coordinates to the "Encoded Polyline Algorithm Format"
// More info: https://developers.google.com/maps/documentation/utilities/polylinealgorithm
// polyline: polyline string in "Encoded Polyline Algorithm Format"
// precision: usually 5 or 6; Google's original algorithm uses 5 digits of decimal precision,
// which is accurate to about a meter. A precision of 6 gives you an accuracy of about 10cm
// more info: https://mapzen.com/blog/polyline-precision/
func Decode(polyline string, precision uint32) []Point {

	group := ""
	coordinates := []float64{}

	for _, letter := range polyline {
		group += string(letter)

		if (int32(letter)-63)&0x20 == 0 {
			coordinates = append(coordinates, decodeElement(group, precision))
			group = ""
		}
	}

	points := []Point{}
	for i := 1; i < len(coordinates); i += 2 {
		points = append(points, Point{
			Latitude:  round(coordinates[i-1], precision),
			Longitude: round(coordinates[i], precision),
		})
	}

	for i := range points {
		if i > 0 {
			points[i].Latitude = round(points[i-1].Latitude+points[i].Latitude, precision)
			points[i].Longitude = round(points[i-1].Longitude+points[i].Longitude, precision)
		}
	}

	return points
}

// Decode5 is a short call for Decode with precision set to 5
// Accuracy is about one meter
func Decode5(polyline string) []Point {
	return Decode(polyline, 5)
}

// Decode6 is a short call for Decode with precision set to 6
// Accuracy is about ten centimeters
func Decode6(polyline string) []Point {
	return Decode(polyline, 6)
}

// encodeElement encodes an coordinate element (i.e. latitude or longitude)
// to the "Encoded Polyline Algorithm Format"
func encodeElement(element float64, precision uint32) string {

	elementInt := int32(math.Round(element * math.Pow10(int(precision))))
	elementInt = elementInt << 1
	if element < 0 {
		elementInt = ^elementInt
	}

	var c chunks
	c.Parse(elementInt)

	return c.String()

}

// decodeElement decodes an coordinate element (i.e. latitude or longitude)
// from the "Encoded Polyline Algorithm Format"
func decodeElement(group string, precision uint32) float64 {

	var c chunks
	c.ParseLine(group)
	return c.Coordinate(precision)
}

// round n to precision digits
func round(n float64, precision uint32) float64 {

	factor := math.Pow10(int(precision))
	return math.Round(n*factor) / factor

}
