package types_test

import (
	"testing"

	"github.com/mizumoto-cn/functional/types"
	assert "github.com/stretchr/testify/assert"
)

// Test for types.Uint16 type

func TestNewUint16(t *testing.T) {
	cases := []uint16{
		uint16(0),
		uint16(1),
	}
	for _, x := range cases {
		assert.True(t, types.NewUint16(x).Value() == x)
	}
}

func TestUint16Compare(t *testing.T) {
	cases := []struct {
		x, y uint16
		want types.Ordering
	}{
		{uint16(0), uint16(0), types.EQ},
		{uint16(0), uint16(1), types.LT},
		{uint16(1), uint16(0), types.GT},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint16(c.x).Compare(types.NewUint16(c.y)) == c.want)
	}
}

func TestUint16LT(t *testing.T) {
	cases := []struct {
		x, y uint16
		want bool
	}{
		{uint16(0), uint16(0), false},
		{uint16(0), uint16(1), true},
		{uint16(1), uint16(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint16(c.x).LT(types.NewUint16(c.y)) == c.want)
	}
}

func TestUint16GT(t *testing.T) {
	cases := []struct {
		x, y uint16
		want bool
	}{
		{uint16(0), uint16(0), false},
		{uint16(0), uint16(1), false},
		{uint16(1), uint16(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint16(c.x).GT(types.NewUint16(c.y)) == c.want)
	}
}

func TestUint16EQ(t *testing.T) {
	cases := []struct {
		x, y uint16
		want bool
	}{
		{uint16(0), uint16(0), true},
		{uint16(0), uint16(1), false},
		{uint16(1), uint16(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint16(c.x).EQ(types.NewUint16(c.y)) == c.want)
	}
}

func TestUint16LE(t *testing.T) {
	cases := []struct {
		x, y uint16
		want bool
	}{
		{uint16(0), uint16(0), true},
		{uint16(0), uint16(1), true},
		{uint16(1), uint16(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint16(c.x).LE(types.NewUint16(c.y)) == c.want)
	}
}

func TestUint16GE(t *testing.T) {
	cases := []struct {
		x, y uint16
		want bool
	}{
		{uint16(0), uint16(0), true},
		{uint16(0), uint16(1), false},
		{uint16(1), uint16(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint16(c.x).GE(types.NewUint16(c.y)) == c.want)
	}
}

func TestUint16String(t *testing.T) {
	cases := []struct {
		x    uint16
		want string
	}{
		{uint16(0), "0"},
		{uint16(1), "1"},
	}
	for _, c := range cases {
		assert.True(t, types.NewUint16(c.x).String() == c.want)
	}
}
