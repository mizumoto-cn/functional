package types

import (
	"reflect"
	"strconv"
)

type number[T Number] struct {
	value T
}

// Compare :: number[T] -> number[T] -> Ordering
func (x number[T]) Compare(y Ord[T]) Ordering {
	if x.value < y.Value() {
		return LT
	} else if x.value > y.Value() {
		return GT
	}
	return EQ
}

// LT :: number[T] -> number[T] -> Bool
func (x number[T]) LT(y Ord[T]) bool {
	return x.Compare(y) == LT
}

// GT :: number[T] -> number[T] -> Bool
func (x number[T]) GT(y Ord[T]) bool {
	return x.Compare(y) == GT
}

// EQ :: number[T] -> number[T] -> Bool
func (x number[T]) EQ(y Ord[T]) bool {
	return x.Compare(y) == EQ
}

// LE :: number[T] -> number[T] -> Bool
func (x number[T]) LE(y Ord[T]) bool {
	return x.Compare(y) != GT
}

// GE :: number[T] -> number[T] -> Bool
func (x number[T]) GE(y Ord[T]) bool {
	return x.Compare(y) != LT
}

// Value :: number[T] -> T
func (x number[T]) Value() T {
	return x.value
}

// String :: number[T] -> string
func (x number[T]) String() string {
	t := reflect.TypeOf(x.value)
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(int64(x.value), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(uint64(x.value), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(float64(x.value), 'f', -1, 64)
	case reflect.Uintptr: // print as hex
		return strconv.FormatUint(uint64(x.value), 16)
	default:
		return "NaN"
	}
}
