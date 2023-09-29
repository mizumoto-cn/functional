package types

// Int64 :: number[int64]
// Int is an alias of Int64.
// Int64 :: Ord[int64]

var _ Ord[int64] = Int64{int64(0)}
var _ Int = Int64{int64(0)}

// use alias declaration instead of type definition
type Int = Int64

func NewInt(x int64) Int {
	return Int{x}
}

// Int64 is a wrapper around int64 that implements Ord.
type Int64 = number[int64]

// NewInt64 returns a new Int64.
func NewInt64(x int64) Int64 {
	return Int64{x}
}

// Int32 :: number[int32]
// Int32 :: Ord[int32]
var _ Ord[int32] = Int32{int32(0)}

func NewInt32(x int32) Int32 {
	return Int32{x}
}

type Int32 = number[int32]

// Int16 :: number[int16]
// Int16 :: Ord[int16]
var _ Ord[int16] = Int16{int16(0)}

func NewInt16(x int16) Int16 {
	return Int16{x}
}

type Int16 = number[int16]

// Int8 :: number[int8]
// Int8 :: Ord[int8]
var _ Ord[int8] = Int8{int8(0)}

func NewInt8(x int8) Int8 {
	return Int8{x}
}

type Int8 = number[int8]
