package specodec

import "strconv"

// FmtFloat32 returns the shortest decimal string that uniquely identifies
// the given float32 value and round-trips back to the same float32 bits.
//
// Uses strconv.FormatFloat with bitSize=32, which implements the Ryu algorithm
// internally (Go stdlib >= 1.15).
//
// TODO: if a standalone Ryu implementation is needed (e.g. for portability),
// replace this body with an explicit Ryu f32 port.
func fmtFloat32(value float32) string {
	return strconv.FormatFloat(float64(value), 'f', -1, 32)
}
