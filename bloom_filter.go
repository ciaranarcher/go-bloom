package bloom

// Filter allows a way to prime and seach a bloom filter
type Filter struct {
	bitmapSize int
	Bitmap     []bool
}

// NewFilter creates a new bloom filter with a bitmap of the passed size
func NewFilter(bitmapSize int) Filter {
	return Filter{
		bitmapSize: bitmapSize,
		Bitmap:     make([]bool, bitmapSize, bitmapSize),
	}
}

// LoadItems will prime the bloom filter with data
func (b *Filter) LoadItems(words []string) {
}

// Query will see if there is a match in the bloom filter for the passed word
func (b *Filter) Query(word string) bool {
	return false
}
