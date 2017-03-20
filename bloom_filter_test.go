package bloom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	filter := NewFilter(1024)
	filter.LoadItems([]string{"the", "quick", "brown", "fox"})
	assert.True(t, filter.Query("the"))
}
