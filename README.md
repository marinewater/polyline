# polyline
[![Build Status](https://travis-ci.org/marinewater/polyline.svg?branch=master)](https://travis-ci.org/marinewater/polyline)

Go package to encode/decode polylines in "Encoded Polyline Algorithm Format"

## Types
### Point
```go
type Point struct {
	Latitude  float64
	Longitude float64
}
```
Point is a coordinate of a point on the polyline

## Methods
### Encode
 ```go
Encode(points []Point, precision uint32) string
```
Encode encodes coordinates to the "Encoded Polyline Algorithm Format"

More info: https://developers.google.com/maps/documentation/utilities/polylinealgorithm

points: points of the polyline

precision: usually 5 or 6; Google's original algorithm uses 5 digits of decimal precision,
which is accurate to about a meter. A precision of 6 gives you an accuracy of about 10cm
more info: https://mapzen.com/blog/polyline-precision/

### Decode
```go
Decode(polyline string, precision uint32) []Point
```
Decode decodes coordinates to the "Encoded Polyline Algorithm Format"

More info: https://developers.google.com/maps/documentation/utilities/polylinealgorithm

polyline: polyline string in "Encoded Polyline Algorithm Format"

precision: usually 5 or 6; Google's original algorithm uses 5 digits of decimal precision,
which is accurate to about a meter. A precision of 6 gives you an accuracy of about 10cm
more info: https://mapzen.com/blog/polyline-precision/

### Encode5
```go
Encode5(points []Point) string
```
Encode5 is a short call for Encode with precision set to 5

Accuracy is about one meter

### Encode6
```go
Encode6(points []Point) string
```
Encode6 is a short call for Encode with precision set to 6

Accuracy is about ten centimeters

### Decode5
```go
Decode5(polyline string) []Point
```
Decode5 is a short call for Decode with precision set to 5

Accuracy is about one meter

### Decode6
```go
Decode6(polyline string) []Point
```
Decode6 is a short call for Decode with precision set to 6

Accuracy is about ten centimeters
