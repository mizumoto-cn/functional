package types_test

import (
	"testing"

	"github.com/mizumoto-cn/functional/types"
	assert "github.com/stretchr/testify/assert"
)

// Test for types.Int64 type

func TestNewInt(t *testing.T) {
	cases := []int64{
		int64(0),
		int64(1),
		int64(-1),
	}
	for _, x := range cases {
		assert.True(t, types.NewInt(x).Value() == x)
	}
}

func TestNewInt64(t *testing.T) {
	cases := []int64{
		int64(0),
		int64(1),
		int64(-1),
	}
	for _, x := range cases {
		assert.True(t, types.NewInt64(x).Value() == x)
	}
}

func TestInt64Compare(t *testing.T) {
	cases := []struct {
		x, y int64
		want types.Ordering
	}{
		{int64(0), int64(0), types.EQ},
		{int64(0), int64(1), types.LT},
		{int64(1), int64(0), types.GT},
		{int64(0), int64(-1), types.GT},
		{int64(-1), int64(0), types.LT},
		{int64(1), int64(-1), types.GT},
		{int64(-1), int64(1), types.LT},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt64(c.x).Compare(types.NewInt64(c.y)) == c.want)
	}
}

func TestInt64LT(t *testing.T) {
	cases := []struct {
		x, y int64
		want bool
	}{
		{int64(0), int64(0), false},
		{int64(0), int64(1), true},
		{int64(1), int64(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt64(c.x).LT(types.NewInt64(c.y)) == c.want)
	}
}

func TestInt64GT(t *testing.T) {
	cases := []struct {
		x, y int64
		want bool
	}{
		{int64(0), int64(0), false},
		{int64(0), int64(1), false},
		{int64(1), int64(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt64(c.x).GT(types.NewInt64(c.y)) == c.want)
	}
}

func TestInt64EQ(t *testing.T) {
	cases := []struct {
		x, y int64
		want bool
	}{
		{int64(0), int64(0), true},
		{int64(0), int64(1), false},
		{int64(1), int64(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt64(c.x).EQ(types.NewInt64(c.y)) == c.want)
	}
}

func TestInt64LE(t *testing.T) {
	cases := []struct {
		x, y int64
		want bool
	}{
		{int64(0), int64(0), true},
		{int64(0), int64(1), true},
		{int64(1), int64(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt64(c.x).LE(types.NewInt64(c.y)) == c.want)
	}
}

func TestInt64GE(t *testing.T) {
	cases := []struct {
		x, y int64
		want bool
	}{
		{int64(0), int64(0), true},
		{int64(0), int64(1), false},
		{int64(1), int64(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt64(c.x).GE(types.NewInt64(c.y)) == c.want)
	}
}

func TestInt64String(t *testing.T) {
	cases := []struct {
		x    int64
		want string
	}{
		{int64(0), "0"},
		{int64(1), "1"},
		{int64(-1), "-1"},
	}
	for _, c := range cases {
		assert.True(t, types.NewInt64(c.x).String() == c.want)
	}
}
