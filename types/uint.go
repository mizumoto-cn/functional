package types

// Uint64 :: number[uint64]
// Uint64 :: Ord[uint64]
// Uint is an alias of Uint64.
var _ Ord[uint64] = Uint64{uint64(0)}

func NewUint64(x uint64) Uint64 {
	return Uint64{x}
}

func NewUint(x uint64) Uint {
	return Uint64{x}
}

type Uint64 = number[uint64]
type Uint = Uint64

// Uint32 :: number[uint32]
// Uint32 :: Ord[uint32]
var _ Ord[uint32] = Uint32{uint32(0)}

func NewUint32(x uint32) Uint32 {
	return Uint32{x}
}

type Uint32 = number[uint32]

// Uint16 :: number[uint16]
// Uint16 :: Ord[uint16]
var _ Ord[uint16] = Uint16{uint16(0)}

func NewUint16(x uint16) Uint16 {
	return Uint16{x}
}

type Uint16 = number[uint16]

// Uint8 :: number[uint8]
// Uint8 :: Ord[uint8]
var _ Ord[uint8] = Uint8{uint8(0)}

func NewUint8(x uint8) Uint8 {
	return Uint8{x}
}

type Uint8 = number[uint8]

// Uintptr :: number[uintptr]
// Uintptr :: Ord[uintptr]
var _ Ord[uintptr] = Uintptr{uintptr(0)}

func NewUintptr(x uintptr) Uintptr {
	return Uintptr{x}
}

type Uintptr = number[uintptr]
