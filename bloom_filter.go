package bloom

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

const minInt32 = -2147483648
const maxInt32 = 2147483647

// Filter allows a way to prime and seach a bloom filter
type Filter struct {
	bitmapSize   int32
	bitmap       []bool
	numberMapper NumberMapper
}

// NewFilter creates a new bloom filter with a bitmap of the passed size
func NewFilter(bitmapSize int32) Filter {
	return Filter{
		bitmapSize:   bitmapSize,
		bitmap:       make([]bool, bitmapSize, bitmapSize),
		numberMapper: NewNumberMapper(minInt32, maxInt32, 0, bitmapSize-1),
	}
}

// LoadItems will prime the bloom filter with data
func (b *Filter) LoadItems(words []string) {
	for _, word := range words {
		hashedWord := hashWord(word)
		wordAsInts := getIntsFromHash(hashedWord)
		mappedInts := b.mapToRange(wordAsInts)
		b.updateBitmap(mappedInts)
	}
}

// Query will see if there is a match in the bloom filter for the passed word
func (b *Filter) Query(word string) bool {
	hashedWord := hashWord(word)
	wordAsInts := getIntsFromHash(hashedWord)
	mappedInts := b.mapToRange(wordAsInts)

	for _, n := range mappedInts {
		if b.bitmap[n] == false {
			return false
		}
	}
	return true
}

func (b *Filter) mapToRange(wordAsInts []int32) []int32 {
	mappedInts := []int32{}

	for _, wordInt := range wordAsInts {
		mappedNum, err := b.numberMapper.Map(wordInt * 1.0)
		if err != nil {
			fmt.Println(err)
		}
		mappedInts = append(mappedInts, mappedNum)
	}
	return mappedInts
}

func hashWord(word string) []byte {
	res := sha256.Sum256([]byte(word))
	return res[:] // Return a slice rather than a fixed array
}

func (b *Filter) updateBitmap(mappedInts []int32) {
	for _, n := range mappedInts {
		b.bitmap[n] = true
	}
}

func getIntsFromHash(hash []byte) []int32 {
	return []int32{
		asInt(hash[:4]),
		asInt(hash[4:8]),
		asInt(hash[8:12]),
		asInt(hash[12:16]),
		asInt(hash[16:20]),
		asInt(hash[20:24]),
		asInt(hash[24:28]),
		asInt(hash[28:32]),
	}
}

func asInt(data []byte) int32 {
	var n int32
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, &n)

	if err != nil {
		fmt.Println(err)
	}

	return n
}
