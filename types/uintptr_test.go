package types_test

import (
	"testing"

	"github.com/mizumoto-cn/functional/types"
	assert "github.com/stretchr/testify/assert"
)

// Test for types.Uintptr type

func TestNewUintptr(t *testing.T) {
	cases := []uintptr{
		uintptr(0),
		uintptr(1),
	}
	for _, x := range cases {
		assert.True(t, types.NewUintptr(x).Value() == x)
	}
}

func TestUintptrCompare(t *testing.T) {
	cases := []struct {
		x, y uintptr
		want types.Ordering
	}{
		{uintptr(0), uintptr(0), types.EQ},
		{uintptr(0), uintptr(1), types.LT},
		{uintptr(1), uintptr(0), types.GT},
	}
	for _, c := range cases {
		assert.True(t, types.NewUintptr(c.x).Compare(types.NewUintptr(c.y)) == c.want)
	}
}

func TestUintptrLT(t *testing.T) {
	cases := []struct {
		x, y uintptr
		want bool
	}{
		{uintptr(0), uintptr(0), false},
		{uintptr(0), uintptr(1), true},
		{uintptr(1), uintptr(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewUintptr(c.x).LT(types.NewUintptr(c.y)) == c.want)
	}
}

func TestUintptrGT(t *testing.T) {
	cases := []struct {
		x, y uintptr
		want bool
	}{
		{uintptr(0), uintptr(0), false},
		{uintptr(0), uintptr(1), false},
		{uintptr(1), uintptr(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewUintptr(c.x).GT(types.NewUintptr(c.y)) == c.want)
	}
}

func TestUintptrEQ(t *testing.T) {
	cases := []struct {
		x, y uintptr
		want bool
	}{
		{uintptr(0), uintptr(0), true},
		{uintptr(0), uintptr(1), false},
		{uintptr(1), uintptr(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewUintptr(c.x).EQ(types.NewUintptr(c.y)) == c.want)
	}
}

func TestUintptrLE(t *testing.T) {
	cases := []struct {
		x, y uintptr
		want bool
	}{
		{uintptr(0), uintptr(0), true},
		{uintptr(0), uintptr(1), true},
		{uintptr(1), uintptr(0), false},
	}
	for _, c := range cases {
		assert.True(t, types.NewUintptr(c.x).LE(types.NewUintptr(c.y)) == c.want)
	}
}

func TestUintptrGE(t *testing.T) {
	cases := []struct {
		x, y uintptr
		want bool
	}{
		{uintptr(0), uintptr(0), true},
		{uintptr(0), uintptr(1), false},
		{uintptr(1), uintptr(0), true},
	}
	for _, c := range cases {
		assert.True(t, types.NewUintptr(c.x).GE(types.NewUintptr(c.y)) == c.want)
	}
}

func TestUintptrString(t *testing.T) {
	cases := []struct {
		x    uintptr
		want string
	}{
		{uintptr(0), "0"},
		{uintptr(10) /*hex*/, "a"},
		{uintptr(100) /*hex*/, "64"},
	}
	for _, c := range cases {
		assert.True(t, types.NewUintptr(c.x).String() == c.want)
	}
}
