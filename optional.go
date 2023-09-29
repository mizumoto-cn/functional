package functional

import "reflect"

type Maybe[T any] interface {
	Just(any) Maybe[any]
	IsNothing() bool
	IsPresent() bool
	IsPtr() bool
	MatchesType(t reflect.Type) bool
	MatchesKind(k reflect.Kind) bool
	// Matches(func(any) bool) bool
}

func Just[T any](in T) Maybe[T] {
	isNothing := IsNothing(in)
	return some[T]{value: in, isNothing: isNothing, isPresent: !isNothing}
}

// some is a Maybe implementation
type some[T any] struct {
	value     T
	isNothing bool
	isPresent bool
}

type none struct {
	some[any]
}

var None none = none{some[any]{isNothing: true, isPresent: false}}

var _ Maybe[any] = some[any]{}

func (s some[T]) Just(x any) Maybe[any] {
	return Just(x)
}

func (s some[T]) IsNothing() bool {
	return s.isNothing
}

func (s some[T]) IsPresent() bool {
	return s.isPresent
}

func (s some[T]) IsPtr() bool {
	return IsPtr(s.value)
}

func (s some[T]) MatchesType(t reflect.Type) bool {
	return TypeOf(s.value) == t
}

func (s some[T]) MatchesKind(k reflect.Kind) bool {
	return KindOf(s.value) == k
}

// func (s some[T]) Matches(f func(any) bool) bool
