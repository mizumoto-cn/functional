package types_test

import (
	"testing"

	"github.com/mizumoto-cn/functional/types"
	assert "github.com/stretchr/testify/assert"
)

// Test for types.Uint64 type

func TestNewUint(t *testing.T) {
	cases := []uint64{
		uint64(0),
		uint64(1),
	}
	for _, x := range cases {
		assert.True(t, types.NewUint(x).Value() == x)
	}
}

func TestNewUint64(t *testing.T) {
	cases := []uint64{
		uint64(0),
		uint64(1),
	}
	for _, x := range cases {
		assert.True(t, types.NewUint64(x).Value() == x)
	}
}

func TestUint64Compare(t *testing.T) {
	cases := []struct {
		x, y uint64
		want types.Ordering
	}{
		{uint64(0), uint64(0), types.EQ},
		{uint64(0), uint64(1), types.LT},
		{uint64(1), uint64(0), types.GT},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint64(c.x).Compare(types.NewUint64(c.y)) == c.want)
	}
}

func TestUint64LT(t *testing.T) {
	cases := []struct {
		x, y uint64
		want bool
	}{
		{uint64(0), uint64(0), false},
		{uint64(0), uint64(1), true},
		{uint64(1), uint64(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint64(c.x).LT(types.NewUint64(c.y)) == c.want)
	}
}

func TestUint64GT(t *testing.T) {
	cases := []struct {
		x, y uint64
		want bool
	}{
		{uint64(0), uint64(0), false},
		{uint64(0), uint64(1), false},
		{uint64(1), uint64(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint64(c.x).GT(types.NewUint64(c.y)) == c.want)
	}
}

func TestUint64EQ(t *testing.T) {
	cases := []struct {
		x, y uint64
		want bool
	}{
		{uint64(0), uint64(0), true},
		{uint64(0), uint64(1), false},
		{uint64(1), uint64(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint64(c.x).EQ(types.NewUint64(c.y)) == c.want)
	}
}

func TestUint64LE(t *testing.T) {
	cases := []struct {
		x, y uint64
		want bool
	}{
		{uint64(0), uint64(0), true},
		{uint64(0), uint64(1), true},
		{uint64(1), uint64(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint64(c.x).LE(types.NewUint64(c.y)) == c.want)
	}
}

func TestUint64GE(t *testing.T) {
	cases := []struct {
		x, y uint64
		want bool
	}{
		{uint64(0), uint64(0), true},
		{uint64(0), uint64(1), false},
		{uint64(1), uint64(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint64(c.x).GE(types.NewUint64(c.y)) == c.want)
	}
}

func TestUint64String(t *testing.T) {
	cases := []struct {
		x    uint64
		want string
	}{
		{uint64(0), "0"},
		{uint64(1), "1"},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint64(c.x).String() == c.want)
	}
}
