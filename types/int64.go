package types

import (
	"strconv"
)

var _ Ord[int64] = Int64{int64(0)}
var _ Int = Int64{int64(0)}

// use alias declaration instead of type definition
type Int = Int64

func NewInt(x int64) Int {
	return Int{x}
}

// Int64 is a wrapper around int64 that implements Ord.
type Int64 struct {
	value int64
}

// NewInt64 returns a new Int64.
func NewInt64(x int64) Int64 {
	return Int64{x}
}

// Compare :: Int64 -> Int64 -> Ordering
func (x Int64) Compare(y Ord[int64]) Ordering {
	if x.value < y.Value() {
		return LT
	} else if x.value > y.Value() {
		return GT
	}
	return EQ
}

// LT :: Int64 -> Int64 -> Bool
func (x Int64) LT(y Ord[int64]) bool {
	return x.Compare(y) == LT
}

// GT :: Int64 -> Int64 -> Bool
func (x Int64) GT(y Ord[int64]) bool {
	return x.Compare(y) == GT
}

// EQ :: Int64 -> Int64 -> Bool
func (x Int64) EQ(y Ord[int64]) bool {
	return x.Compare(y) == EQ
}

// LE :: Int64 -> Int64 -> Bool
func (x Int64) LE(y Ord[int64]) bool {
	return x.Compare(y) != GT
}

// GE :: Int64 -> Int64 -> Bool
func (x Int64) GE(y Ord[int64]) bool {
	return x.Compare(y) != LT
}

// Value :: Int64 -> int64
func (x Int64) Value() int64 {
	return x.value
}

// String :: Int64 -> string
func (x Int64) String() string {
	return strconv.FormatInt(x.value, 10)
}
