package types_test

import (
	"testing"

	"github.com/mizumoto-cn/functional/types"
	assert "github.com/stretchr/testify/assert"
)

// Test for types.Int8 type

func TestNewInt8(t *testing.T) {
	cases := []int8{
		int8(0),
		int8(1),
		int8(-1),
	}
	for _, x := range cases {
		assert.True(t, types.NewInt8(x).Value() == x)
	}
}

func TestInt8Compare(t *testing.T) {
	cases := []struct {
		x, y int8
		want types.Ordering
	}{
		{int8(0), int8(0), types.EQ},
		{int8(0), int8(1), types.LT},
		{int8(1), int8(0), types.GT},
		{int8(0), int8(-1), types.GT},
		{int8(-1), int8(0), types.LT},
		{int8(1), int8(-1), types.GT},
		{int8(-1), int8(1), types.LT},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt8(c.x).Compare(types.NewInt8(c.y)) == c.want)
	}
}

func TestInt8LT(t *testing.T) {
	cases := []struct {
		x, y int8
		want bool
	}{
		{int8(0), int8(0), false},
		{int8(0), int8(1), true},
		{int8(1), int8(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt8(c.x).LT(types.NewInt8(c.y)) == c.want)
	}
}

func TestInt8GT(t *testing.T) {
	cases := []struct {
		x, y int8
		want bool
	}{
		{int8(0), int8(0), false},
		{int8(0), int8(1), false},
		{int8(1), int8(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt8(c.x).GT(types.NewInt8(c.y)) == c.want)
	}
}

func TestInt8EQ(t *testing.T) {
	cases := []struct {
		x, y int8
		want bool
	}{
		{int8(0), int8(0), true},
		{int8(0), int8(1), false},
		{int8(1), int8(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt8(c.x).EQ(types.NewInt8(c.y)) == c.want)
	}
}

func TestInt8LE(t *testing.T) {
	cases := []struct {
		x, y int8
		want bool
	}{
		{int8(0), int8(0), true},
		{int8(0), int8(1), true},
		{int8(1), int8(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt8(c.x).LE(types.NewInt8(c.y)) == c.want)
	}
}

func TestInt8GE(t *testing.T) {
	cases := []struct {
		x, y int8
		want bool
	}{
		{int8(0), int8(0), true},
		{int8(0), int8(1), false},
		{int8(1), int8(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt8(c.x).GE(types.NewInt8(c.y)) == c.want)
	}
}

func TestInt8String(t *testing.T) {
	cases := []struct {
		x    int8
		want string
	}{
		{int8(0), "0"},
		{int8(1), "1"},
		{int8(-1), "-1"},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt8(c.x).String() == c.want)
	}
}
