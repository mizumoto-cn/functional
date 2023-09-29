package types_test

import (
	"testing"

	"github.com/mizumoto-cn/functional/types"
	assert "github.com/stretchr/testify/assert"
)

// Test for types.Uint32 type

func TestNewUint32(t *testing.T) {
	cases := []uint32{
		uint32(0),
		uint32(1),
	}
	for _, x := range cases {
		assert.True(t, types.NewUint32(x).Value() == x)
	}
}

func TestUint32Compare(t *testing.T) {
	cases := []struct {
		x, y uint32
		want types.Ordering
	}{
		{uint32(0), uint32(0), types.EQ},
		{uint32(0), uint32(1), types.LT},
		{uint32(1), uint32(0), types.GT},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint32(c.x).Compare(types.NewUint32(c.y)) == c.want)
	}
}

func TestUint32LT(t *testing.T) {
	cases := []struct {
		x, y uint32
		want bool
	}{
		{uint32(0), uint32(0), false},
		{uint32(0), uint32(1), true},
		{uint32(1), uint32(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint32(c.x).LT(types.NewUint32(c.y)) == c.want)
	}
}

func TestUint32GT(t *testing.T) {
	cases := []struct {
		x, y uint32
		want bool
	}{
		{uint32(0), uint32(0), false},
		{uint32(0), uint32(1), false},
		{uint32(1), uint32(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint32(c.x).GT(types.NewUint32(c.y)) == c.want)
	}
}

func TestUint32EQ(t *testing.T) {
	cases := []struct {
		x, y uint32
		want bool
	}{
		{uint32(0), uint32(0), true},
		{uint32(0), uint32(1), false},
		{uint32(1), uint32(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint32(c.x).EQ(types.NewUint32(c.y)) == c.want)
	}
}

func TestUint32LE(t *testing.T) {
	cases := []struct {
		x, y uint32
		want bool
	}{
		{uint32(0), uint32(0), true},
		{uint32(0), uint32(1), true},
		{uint32(1), uint32(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint32(c.x).LE(types.NewUint32(c.y)) == c.want)
	}
}

func TestUint32GE(t *testing.T) {
	cases := []struct {
		x, y uint32
		want bool
	}{
		{uint32(0), uint32(0), true},
		{uint32(0), uint32(1), false},
		{uint32(1), uint32(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint32(c.x).GE(types.NewUint32(c.y)) == c.want)
	}
}

func TestUint32String(t *testing.T) {
	cases := []struct {
		x    uint32
		want string
	}{
		{uint32(0), "0"},
		{uint32(1), "1"},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint32(c.x).String() == c.want)
	}
}
