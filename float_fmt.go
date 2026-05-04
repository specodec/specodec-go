package specodec

import "github.com/specodec/specodec-runtime-golang/ryu"

func formatFloat32(value float32) string {
	return ryu.Float32ToString(value)
}

func formatFloat64(value float64) string {
	return ryu.Float64ToString(value)
}
