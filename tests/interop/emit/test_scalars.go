package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
)


func runScalar_Int8Min(vecDir, outDir string) (passed, failed int) {
	return tryTest("int8_min mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/int8_min.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadInt32()
		w := specodec.NewMsgPackWriter()
		w.WriteInt32(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/int8_min.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Int8Max(vecDir, outDir string) (passed, failed int) {
	return tryTest("int8_max mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/int8_max.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadInt32()
		w := specodec.NewMsgPackWriter()
		w.WriteInt32(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/int8_max.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Int16Min(vecDir, outDir string) (passed, failed int) {
	return tryTest("int16_min mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/int16_min.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadInt32()
		w := specodec.NewMsgPackWriter()
		w.WriteInt32(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/int16_min.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Int16Max(vecDir, outDir string) (passed, failed int) {
	return tryTest("int16_max mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/int16_max.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadInt32()
		w := specodec.NewMsgPackWriter()
		w.WriteInt32(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/int16_max.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Int32Min(vecDir, outDir string) (passed, failed int) {
	return tryTest("int32_min mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/int32_min.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadInt32()
		w := specodec.NewMsgPackWriter()
		w.WriteInt32(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/int32_min.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Int32Max(vecDir, outDir string) (passed, failed int) {
	return tryTest("int32_max mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/int32_max.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadInt32()
		w := specodec.NewMsgPackWriter()
		w.WriteInt32(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/int32_max.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Int64Min(vecDir, outDir string) (passed, failed int) {
	return tryTest("int64_min mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/int64_min.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadInt64()
		w := specodec.NewMsgPackWriter()
		w.WriteInt64(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/int64_min.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Int64Max(vecDir, outDir string) (passed, failed int) {
	return tryTest("int64_max mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/int64_max.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadInt64()
		w := specodec.NewMsgPackWriter()
		w.WriteInt64(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/int64_max.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Uint8Max(vecDir, outDir string) (passed, failed int) {
	return tryTest("uint8_max mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/uint8_max.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadUint32()
		w := specodec.NewMsgPackWriter()
		w.WriteUint32(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/uint8_max.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Uint16Max(vecDir, outDir string) (passed, failed int) {
	return tryTest("uint16_max mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/uint16_max.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadUint32()
		w := specodec.NewMsgPackWriter()
		w.WriteUint32(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/uint16_max.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Uint32Max(vecDir, outDir string) (passed, failed int) {
	return tryTest("uint32_max mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/uint32_max.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadUint32()
		w := specodec.NewMsgPackWriter()
		w.WriteUint32(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/uint32_max.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Uint64Max(vecDir, outDir string) (passed, failed int) {
	return tryTest("uint64_max mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/uint64_max.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadUint64()
		w := specodec.NewMsgPackWriter()
		w.WriteUint64(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/uint64_max.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Float3215(vecDir, outDir string) (passed, failed int) {
	return tryTest("float32_1.5 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/float32_1.5.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadFloat32()
		w := specodec.NewMsgPackWriter()
		w.WriteFloat32(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/float32_1.5.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Float32NegZero(vecDir, outDir string) (passed, failed int) {
	return tryTest("float32_neg_zero mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/float32_neg_zero.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadFloat32()
		w := specodec.NewMsgPackWriter()
		w.WriteFloat32(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/float32_neg_zero.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Float32Inf(vecDir, outDir string) (passed, failed int) {
	return tryTest("float32_inf mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/float32_inf.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadFloat32()
		w := specodec.NewMsgPackWriter()
		w.WriteFloat32(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/float32_inf.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Float32NegInf(vecDir, outDir string) (passed, failed int) {
	return tryTest("float32_neg_inf mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/float32_neg_inf.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadFloat32()
		w := specodec.NewMsgPackWriter()
		w.WriteFloat32(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/float32_neg_inf.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Float32Nan(vecDir, outDir string) (passed, failed int) {
	return tryTest("float32_nan mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/float32_nan.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadFloat32()
		w := specodec.NewMsgPackWriter()
		w.WriteFloat32(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/float32_nan.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Float64Pi(vecDir, outDir string) (passed, failed int) {
	return tryTest("float64_pi mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/float64_pi.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadFloat64()
		w := specodec.NewMsgPackWriter()
		w.WriteFloat64(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/float64_pi.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Float64NegZero(vecDir, outDir string) (passed, failed int) {
	return tryTest("float64_neg_zero mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/float64_neg_zero.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadFloat64()
		w := specodec.NewMsgPackWriter()
		w.WriteFloat64(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/float64_neg_zero.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Float64Inf(vecDir, outDir string) (passed, failed int) {
	return tryTest("float64_inf mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/float64_inf.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadFloat64()
		w := specodec.NewMsgPackWriter()
		w.WriteFloat64(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/float64_inf.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Float64NegInf(vecDir, outDir string) (passed, failed int) {
	return tryTest("float64_neg_inf mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/float64_neg_inf.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadFloat64()
		w := specodec.NewMsgPackWriter()
		w.WriteFloat64(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/float64_neg_inf.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Float64Nan(vecDir, outDir string) (passed, failed int) {
	return tryTest("float64_nan mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/float64_nan.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadFloat64()
		w := specodec.NewMsgPackWriter()
		w.WriteFloat64(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/float64_nan.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_StrEmpty(vecDir, outDir string) (passed, failed int) {
	return tryTest("str_empty mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/str_empty.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadString()
		w := specodec.NewMsgPackWriter()
		w.WriteString(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/str_empty.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_StrAscii(vecDir, outDir string) (passed, failed int) {
	return tryTest("str_ascii mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/str_ascii.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadString()
		w := specodec.NewMsgPackWriter()
		w.WriteString(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/str_ascii.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_StrNullByte(vecDir, outDir string) (passed, failed int) {
	return tryTest("str_null_byte mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/str_null_byte.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadString()
		w := specodec.NewMsgPackWriter()
		w.WriteString(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/str_null_byte.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_StrEscape(vecDir, outDir string) (passed, failed int) {
	return tryTest("str_escape mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/str_escape.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadString()
		w := specodec.NewMsgPackWriter()
		w.WriteString(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/str_escape.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_StrUnicode(vecDir, outDir string) (passed, failed int) {
	return tryTest("str_unicode mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/str_unicode.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadString()
		w := specodec.NewMsgPackWriter()
		w.WriteString(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/str_unicode.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Str31(vecDir, outDir string) (passed, failed int) {
	return tryTest("str_31 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/str_31.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadString()
		w := specodec.NewMsgPackWriter()
		w.WriteString(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/str_31.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Str32(vecDir, outDir string) (passed, failed int) {
	return tryTest("str_32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/str_32.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadString()
		w := specodec.NewMsgPackWriter()
		w.WriteString(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/str_32.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Str255(vecDir, outDir string) (passed, failed int) {
	return tryTest("str_255 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/str_255.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadString()
		w := specodec.NewMsgPackWriter()
		w.WriteString(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/str_255.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Str256(vecDir, outDir string) (passed, failed int) {
	return tryTest("str_256 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/str_256.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadString()
		w := specodec.NewMsgPackWriter()
		w.WriteString(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/str_256.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_BytesEmpty(vecDir, outDir string) (passed, failed int) {
	return tryTest("bytes_empty mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/bytes_empty.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadBytes()
		w := specodec.NewMsgPackWriter()
		w.WriteBytes(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/bytes_empty.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_BytesSmall(vecDir, outDir string) (passed, failed int) {
	return tryTest("bytes_small mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/bytes_small.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadBytes()
		w := specodec.NewMsgPackWriter()
		w.WriteBytes(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/bytes_small.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Bytes31(vecDir, outDir string) (passed, failed int) {
	return tryTest("bytes_31 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/bytes_31.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadBytes()
		w := specodec.NewMsgPackWriter()
		w.WriteBytes(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/bytes_31.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Bytes32(vecDir, outDir string) (passed, failed int) {
	return tryTest("bytes_32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/bytes_32.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadBytes()
		w := specodec.NewMsgPackWriter()
		w.WriteBytes(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/bytes_32.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Bytes255(vecDir, outDir string) (passed, failed int) {
	return tryTest("bytes_255 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/bytes_255.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadBytes()
		w := specodec.NewMsgPackWriter()
		w.WriteBytes(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/bytes_255.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_Bytes256(vecDir, outDir string) (passed, failed int) {
	return tryTest("bytes_256 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/bytes_256.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadBytes()
		w := specodec.NewMsgPackWriter()
		w.WriteBytes(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/bytes_256.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_BytesZeros(vecDir, outDir string) (passed, failed int) {
	return tryTest("bytes_zeros mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/bytes_zeros.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadBytes()
		w := specodec.NewMsgPackWriter()
		w.WriteBytes(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/bytes_zeros.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_BytesFf(vecDir, outDir string) (passed, failed int) {
	return tryTest("bytes_ff mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/bytes_ff.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadBytes()
		w := specodec.NewMsgPackWriter()
		w.WriteBytes(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/bytes_ff.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_BoolTrue(vecDir, outDir string) (passed, failed int) {
	return tryTest("bool_true mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/bool_true.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadBool()
		w := specodec.NewMsgPackWriter()
		w.WriteBool(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/bool_true.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}

func runScalar_BoolFalse(vecDir, outDir string) (passed, failed int) {
	return tryTest("bool_false mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "scalars/bool_false.mp"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		val := r.ReadBool()
		w := specodec.NewMsgPackWriter()
		w.WriteBool(val)
		err = os.WriteFile(filepath.Join(outDir, "scalars/bool_false.mp"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
}


func runScalars(vecDir, outDir string) (passed, failed int) {
	var p, f int
	p, f = runScalar_Int8Min(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Int8Max(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Int16Min(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Int16Max(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Int32Min(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Int32Max(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Int64Min(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Int64Max(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Uint8Max(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Uint16Max(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Uint32Max(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Uint64Max(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Float3215(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Float32NegZero(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Float32Inf(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Float32NegInf(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Float32Nan(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Float64Pi(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Float64NegZero(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Float64Inf(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Float64NegInf(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Float64Nan(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_StrEmpty(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_StrAscii(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_StrNullByte(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_StrEscape(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_StrUnicode(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Str31(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Str32(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Str255(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Str256(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_BytesEmpty(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_BytesSmall(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Bytes31(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Bytes32(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Bytes255(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_Bytes256(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_BytesZeros(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_BytesFf(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_BoolTrue(vecDir, outDir); passed += p; failed += f
	p, f = runScalar_BoolFalse(vecDir, outDir); passed += p; failed += f

	return
}
