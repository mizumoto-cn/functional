package types_test

import (
	"testing"

	"github.com/mizumoto-cn/functional/types"
	assert "github.com/stretchr/testify/assert"
)

// Test for types.Uint8 type

func TestNewUint8(t *testing.T) {
	cases := []uint8{
		uint8(0),
		uint8(1),
	}
	for _, x := range cases {
		assert.True(t, types.NewUint8(x).Value() == x)
	}
}

func TestUint8Compare(t *testing.T) {
	cases := []struct {
		x, y uint8
		want types.Ordering
	}{
		{uint8(0), uint8(0), types.EQ},
		{uint8(0), uint8(1), types.LT},
		{uint8(1), uint8(0), types.GT},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint8(c.x).Compare(types.NewUint8(c.y)) == c.want)
	}
}

func TestUint8LT(t *testing.T) {
	cases := []struct {
		x, y uint8
		want bool
	}{
		{uint8(0), uint8(0), false},
		{uint8(0), uint8(1), true},
		{uint8(1), uint8(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint8(c.x).LT(types.NewUint8(c.y)) == c.want)
	}
}

func TestUint8GT(t *testing.T) {
	cases := []struct {
		x, y uint8
		want bool
	}{
		{uint8(0), uint8(0), false},
		{uint8(0), uint8(1), false},
		{uint8(1), uint8(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint8(c.x).GT(types.NewUint8(c.y)) == c.want)
	}
}

func TestUint8EQ(t *testing.T) {
	cases := []struct {
		x, y uint8
		want bool
	}{
		{uint8(0), uint8(0), true},
		{uint8(0), uint8(1), false},
		{uint8(1), uint8(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint8(c.x).EQ(types.NewUint8(c.y)) == c.want)
	}
}

func TestUint8LE(t *testing.T) {
	cases := []struct {
		x, y uint8
		want bool
	}{
		{uint8(0), uint8(0), true},
		{uint8(0), uint8(1), true},
		{uint8(1), uint8(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint8(c.x).LE(types.NewUint8(c.y)) == c.want)
	}
}

func TestUint8GE(t *testing.T) {
	cases := []struct {
		x, y uint8
		want bool
	}{
		{uint8(0), uint8(0), true},
		{uint8(0), uint8(1), false},
		{uint8(1), uint8(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint8(c.x).GE(types.NewUint8(c.y)) == c.want)
	}
}

func TestUint8String(t *testing.T) {
	cases := []struct {
		x    uint8
		want string
	}{
		{uint8(0), "0"},
		{uint8(1), "1"},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint8(c.x).String() == c.want)
	}
}
