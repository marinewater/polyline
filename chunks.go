package polyline

import (
	"math"
)

type chunks []int32

// Parse splices an integer into chunks
func (c *chunks) Parse(element int32) {

	c.slice(element, 5)

}

// ParseLine converts and splices string into integer chunks
func (c *chunks) ParseLine(line string) {

	chunkSlice := []int32{}

	for index, letter := range line {
		elementInt := int32(letter) - 63
		if index != len(line)-1 {
			elementInt = elementInt & 31
		}

		chunkSlice = append(chunkSlice, elementInt)
	}

	*c = chunkSlice

}

// String returns the chunks as polyline in base64
func (c *chunks) String() string {

	c.or()
	s := ""
	for i := range *c {
		s += string((*c)[i])
	}
	return s

}

// Coordinate converts integer chunks into a single coordinate
func (c *chunks) Coordinate(precision uint32) float64 {
	resultInt := int32(0)

	for i, element := range *c {
		resultInt += element << uint32(i*5)
	}

	if resultInt&1 == 1 {
		resultInt = ^resultInt
	}
	resultInt = resultInt >> 1

	return float64(resultInt) / math.Pow10(int(precision))
}

// slice splits elements into group of "length" bits
func (c *chunks) slice(element int32, length int) {

	chunkSlice := []int32{}

	bitMask := int32(31)

	for i := 0; int32(math.Pow(2, float64(i))) <= element; i += 5 {
		group := (element >> uint(i)) & bitMask
		chunkSlice = append(chunkSlice, group)
	}

	*c = chunkSlice

}

// or sets the 6th bit to 1 for every chunk except the last one
// as indicator bit for coordinate boundaries.
// It also adds 63 (decimal) to every group to ensure it is in
// ASCII range
func (c *chunks) or() {

	for i := range *c {
		if i < len(*c)-1 {
			(*c)[i] = (*c)[i] | 0x20
		}
		(*c)[i] += 63
	}

}
