package types_test

import (
	"testing"

	"github.com/mizumoto-cn/functional/types"
	assert "github.com/stretchr/testify/assert"
)

// Test for types.Int32 type

func TestNewInt32(t *testing.T) {
	cases := []int32{
		int32(0),
		int32(1),
		int32(-1),
	}
	for _, x := range cases {
		assert.True(t, types.NewInt32(x).Value() == x)
	}
}

func TestInt32Compare(t *testing.T) {
	cases := []struct {
		x, y int32
		want types.Ordering
	}{
		{int32(0), int32(0), types.EQ},
		{int32(0), int32(1), types.LT},
		{int32(1), int32(0), types.GT},
		{int32(0), int32(-1), types.GT},
		{int32(-1), int32(0), types.LT},
		{int32(1), int32(-1), types.GT},
		{int32(-1), int32(1), types.LT},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt32(c.x).Compare(types.NewInt32(c.y)) == c.want)
	}
}

func TestInt32LT(t *testing.T) {
	cases := []struct {
		x, y int32
		want bool
	}{
		{int32(0), int32(0), false},
		{int32(0), int32(1), true},
		{int32(1), int32(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt32(c.x).LT(types.NewInt32(c.y)) == c.want)
	}
}

func TestInt32GT(t *testing.T) {
	cases := []struct {
		x, y int32
		want bool
	}{
		{int32(0), int32(0), false},
		{int32(0), int32(1), false},
		{int32(1), int32(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt32(c.x).GT(types.NewInt32(c.y)) == c.want)
	}
}

func TestInt32EQ(t *testing.T) {
	cases := []struct {
		x, y int32
		want bool
	}{
		{int32(0), int32(0), true},
		{int32(0), int32(1), false},
		{int32(1), int32(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt32(c.x).EQ(types.NewInt32(c.y)) == c.want)
	}
}

func TestInt32LE(t *testing.T) {
	cases := []struct {
		x, y int32
		want bool
	}{
		{int32(0), int32(0), true},
		{int32(0), int32(1), true},
		{int32(1), int32(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt32(c.x).LE(types.NewInt32(c.y)) == c.want)
	}
}

func TestInt32GE(t *testing.T) {
	cases := []struct {
		x, y int32
		want bool
	}{
		{int32(0), int32(0), true},
		{int32(0), int32(1), false},
		{int32(1), int32(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt32(c.x).GE(types.NewInt32(c.y)) == c.want)
	}
}

func TestInt32String(t *testing.T) {
	cases := []struct {
		x    int32
		want string
	}{
		{int32(0), "0"},
		{int32(1), "1"},
		{int32(-1), "-1"},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt32(c.x).String() == c.want)
	}
}
