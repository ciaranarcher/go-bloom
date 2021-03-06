package bloom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	filter := NewFilter(1024)
	filter.LoadItems([]string{"the", "quick", "brown", "fox"})
	assert.True(t, filter.Query("the"))
	assert.True(t, filter.Query("quick"))
	assert.True(t, filter.Query("brown"))
	assert.True(t, filter.Query("fox"))

	assert.False(t, filter.Query("missing"))
	assert.False(t, filter.Query(""))
	assert.False(t, filter.Query("fog"))
}
