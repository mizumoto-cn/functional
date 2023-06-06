package types

import (
	"fmt"

	"github.com/mizumoto-cn/functional/pkg/cmp"
)

// Enum Ordering (types.LT | types.EQ | types.GT)
type Ordering int

const (
	LT Ordering = iota
	EQ
	GT
)

// type Ord interface {

// 2023-06-06 Update: Will change to official Go "cmp" package
// where "Ordered" interface is defined as:
// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
// type Ordered interface { ... }

// // Min returns the smallest of a and b. If a and b are equal, Min returns a.
// func Min[T Ordered](a, b T) T

// // Max returns the largest of a and b., If a and b are equal, Max returns a.
// func Max[T Ordered](a, b T) T

type Ord[T cmp.Ordered] interface {
	// Compare :: a -> a -> Ordering
	Compare(Ord[T]) Ordering

	// LT :: a -> a -> Bool
	LT(Ord[T]) bool

	// GT :: a -> a -> Bool
	GT(Ord[T]) bool

	// EQ :: a -> a -> Bool
	EQ(Ord[T]) bool

	// LE :: a -> a -> Bool
	LE(Ord[T]) bool

	// GE :: a -> a -> Bool
	GE(Ord[T]) bool

	// Value :: a[T] -> T
	Value() T

	fmt.Stringer
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}
