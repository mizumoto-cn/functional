package types_test

import (
	"testing"

	"github.com/mizumoto-cn/functional/types"
	assert "github.com/stretchr/testify/assert"
)

// Test for types.Int16 type

func TestNewInt16(t *testing.T) {
	cases := []int16{
		int16(0),
		int16(1),
		int16(-1),
	}
	for _, x := range cases {
		assert.True(t, types.NewInt16(x).Value() == x)
	}
}

func TestInt16Compare(t *testing.T) {
	cases := []struct {
		x, y int16
		want types.Ordering
	}{
		{int16(0), int16(0), types.EQ},
		{int16(0), int16(1), types.LT},
		{int16(1), int16(0), types.GT},
		{int16(0), int16(-1), types.GT},
		{int16(-1), int16(0), types.LT},
		{int16(1), int16(-1), types.GT},
		{int16(-1), int16(1), types.LT},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt16(c.x).Compare(types.NewInt16(c.y)) == c.want)
	}
}

func TestInt16LT(t *testing.T) {
	cases := []struct {
		x, y int16
		want bool
	}{
		{int16(0), int16(0), false},
		{int16(0), int16(1), true},
		{int16(1), int16(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt16(c.x).LT(types.NewInt16(c.y)) == c.want)
	}
}

func TestInt16GT(t *testing.T) {
	cases := []struct {
		x, y int16
		want bool
	}{
		{int16(0), int16(0), false},
		{int16(0), int16(1), false},
		{int16(1), int16(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt16(c.x).GT(types.NewInt16(c.y)) == c.want)
	}
}

func TestInt16EQ(t *testing.T) {
	cases := []struct {
		x, y int16
		want bool
	}{
		{int16(0), int16(0), true},
		{int16(0), int16(1), false},
		{int16(1), int16(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt16(c.x).EQ(types.NewInt16(c.y)) == c.want)
	}
}

func TestInt16LE(t *testing.T) {
	cases := []struct {
		x, y int16
		want bool
	}{
		{int16(0), int16(0), true},
		{int16(0), int16(1), true},
		{int16(1), int16(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt16(c.x).LE(types.NewInt16(c.y)) == c.want)
	}
}

func TestInt16GE(t *testing.T) {
	cases := []struct {
		x, y int16
		want bool
	}{
		{int16(0), int16(0), true},
		{int16(0), int16(1), false},
		{int16(1), int16(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt16(c.x).GE(types.NewInt16(c.y)) == c.want)
	}
}

func TestInt16String(t *testing.T) {
	cases := []struct {
		x    int16
		want string
	}{
		{int16(0), "0"},
		{int16(1), "1"},
		{int16(-1), "-1"},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt16(c.x).String() == c.want)
	}
}
