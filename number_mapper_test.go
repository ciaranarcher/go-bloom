package bloom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberMapperBadRange(t *testing.T) {
	numMapper := NewNumberMapper(1000, 2000, 1, 100)
	t.Log(numMapper)
	_, err := numMapper.Map(500)

	assert.Error(t, err)
}

func TestNumberMapperGoodMidRange(t *testing.T) {
	numMapper := NewNumberMapper(1000, 2000, 1, 100)
	t.Log(numMapper)
	result, err := numMapper.Map(1500)

	assert.NoError(t, err)
	assert.Equal(t, 51.0, result, "should map the middle of the range correctly")
}

func TestNumberMapperGoodBottomRange(t *testing.T) {
	numMapper := NewNumberMapper(1000, 2000, 1, 100)
	t.Log(numMapper)
	result, err := numMapper.Map(1000)

	assert.NoError(t, err)
	assert.Equal(t, 1.0, result, "should map the bottom of the range correctly")
}

func TestNumberMapperGoodTopRange(t *testing.T) {
	numMapper := NewNumberMapper(1000, 2000, 1, 100)
	t.Log(numMapper)
	result, err := numMapper.Map(2000)

	assert.NoError(t, err)
	assert.Equal(t, 100.0, result, "should map the bottom of the range correctly")
}
