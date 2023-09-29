package types_test

import (
	"testing"

	"github.com/mizumoto-cn/functional/types"
	assert "github.com/stretchr/testify/assert"
)

// Test for types.Float32 type
func TestNewFloat32(t *testing.T) {
	cases := []float32{
		float32(0.0001),
		float32(1.0001),
		float32(-1.2233),
	}
	for _, x := range cases {
		assert.True(t, types.NewFloat32(x).Value() == x)
	}
}
