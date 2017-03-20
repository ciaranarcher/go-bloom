package bloom

import "errors"

// NumberMapper holds the details of how to map numbers
type NumberMapper struct {
	inRangeStart  int32
	inRangeEnd    int32
	outRangeStart int32
	outRangeEnd   int32
}

// NewNumberMapper creates a new NumberMapper
func NewNumberMapper(inRngStrt, inRngEnd, outRngStrt, outRngEnd int32) NumberMapper {
	return NumberMapper{
		inRangeStart:  inRngStrt,
		inRangeEnd:    inRngEnd,
		outRangeStart: outRngStrt,
		outRangeEnd:   outRngEnd,
	}
}

// Map maps a number into the configured range
func (nm *NumberMapper) Map(num int32) (int32, error) {
	if num < nm.inRangeStart || num > nm.inRangeEnd {
		return -1, errors.New("number out of range")
	}

	return nm.mapNum(num), nil
}

func (nm *NumberMapper) mapNum(x int32) int32 {
	xAdj := float32(x)
	inRangeStart := float32(nm.inRangeStart)
	inRangeEnd := float32(nm.inRangeEnd)
	outRangeStart := float32(nm.outRangeStart)
	outRangeEnd := float32(nm.outRangeEnd)

	res := (xAdj-inRangeStart)*(outRangeEnd-outRangeStart)/(inRangeEnd-inRangeStart) + outRangeStart
	// return (x-nm.inRangeStart)*(nm.outRangeEnd-nm.outRangeStart)/(nm.inRangeEnd-nm.inRangeStart) + nm.outRangeStart

	return int32(res)
}
