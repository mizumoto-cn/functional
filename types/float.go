package types

// Float64 :: number[float64]
// Float64 :: Ord[float64]
// Float is an alias of Float64.

var _ Ord[float64] = Float64{float64(0)}
var _ Float = Float64{float64(0)}

func NewFloat(x float64) Float {
	return Float{x}
}

func NewFloat64(x float64) Float64 {
	return Float64{x}
}

type Float64 = number[float64]
type Float = Float64

// Float32 :: number[float32]
// Float32 :: Ord[float32]

var _ Ord[float32] = Float32{float32(0)}

func NewFloat32(x float32) Float32 {
	return Float32{x}
}

type Float32 = number[float32]
