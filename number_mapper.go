package bloom

import (
	"errors"
	"math"
)

// NumberMapper holds the details of how to map numbers
type NumberMapper struct {
	inRangeStart  float64
	inRangeEnd    float64
	outRangeStart float64
	outRangeEnd   float64
}

// NewNumberMapper creates a new NumberMapper
func NewNumberMapper(inRngStrt, inRngEnd, outRngStrt, outRngEnd float64) NumberMapper {
	return NumberMapper{
		inRangeStart:  inRngStrt,
		inRangeEnd:    inRngEnd,
		outRangeStart: outRngStrt,
		outRangeEnd:   outRngEnd,
	}
}

// Map maps a number into the configured range
func (nm *NumberMapper) Map(num float64) (float64, error) {
	if num < nm.inRangeStart || num > nm.inRangeEnd {
		return -1, errors.New("number out of range")
	}

	return nm.mapNum(num), nil
}

func (nm *NumberMapper) mapNum(x float64) float64 {
	y := (x-nm.inRangeStart)*(nm.outRangeEnd-nm.outRangeStart)/(nm.inRangeEnd-nm.inRangeStart) + nm.outRangeStart
	return round(y)
}

func round(num float64) float64 {
	return float64(int(num + math.Copysign(0.5, num)))
}
