package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	all_types_arrays "emit_go/emit_gen/all_types_arrays"
)


func runModel_ArrString(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ArrString mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrString.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.ArrStringCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.ArrStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrString.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrString json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrString.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.ArrStringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.ArrStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrString.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrString unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrString.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.ArrStringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.ArrStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrString.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrString gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrString.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.ArrStringCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.ArrStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrString.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ArrInt32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ArrInt32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrInt32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.ArrInt32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.ArrInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrInt32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrInt32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrInt32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.ArrInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.ArrInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrInt32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrInt32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrInt32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.ArrInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.ArrInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrInt32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrInt32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrInt32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.ArrInt32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.ArrInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrInt32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ArrBoolean(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ArrBoolean mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrBoolean.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.ArrBooleanCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.ArrBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrBoolean.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrBoolean json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrBoolean.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.ArrBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.ArrBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrBoolean.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrBoolean unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrBoolean.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.ArrBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.ArrBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrBoolean.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrBoolean gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrBoolean.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.ArrBooleanCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.ArrBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrBoolean.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ArrFloat64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ArrFloat64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrFloat64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.ArrFloat64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.ArrFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrFloat64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrFloat64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrFloat64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.ArrFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.ArrFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrFloat64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrFloat64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrFloat64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.ArrFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.ArrFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrFloat64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrFloat64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrFloat64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.ArrFloat64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.ArrFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrFloat64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ArrBytes(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ArrBytes mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrBytes.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.ArrBytesCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.ArrBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrBytes.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrBytes json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrBytes.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.ArrBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.ArrBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrBytes.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrBytes unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrBytes.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.ArrBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.ArrBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrBytes.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrBytes gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrBytes.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.ArrBytesCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.ArrBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrBytes.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ArrInt64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ArrInt64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrInt64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.ArrInt64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.ArrInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrInt64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrInt64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrInt64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.ArrInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.ArrInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrInt64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrInt64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrInt64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.ArrInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.ArrInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrInt64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrInt64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrInt64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.ArrInt64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.ArrInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrInt64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ArrUint64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ArrUint64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrUint64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.ArrUint64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.ArrUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrUint64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrUint64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrUint64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.ArrUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.ArrUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrUint64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrUint64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrUint64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.ArrUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.ArrUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrUint64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrUint64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrUint64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.ArrUint64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.ArrUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrUint64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_MultiArr1(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("MultiArr1 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr1.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.MultiArr1Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.MultiArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr1.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr1 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr1.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.MultiArr1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.MultiArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr1.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr1 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr1.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.MultiArr1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.MultiArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr1.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr1 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr1.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.MultiArr1Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.MultiArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr1.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_MultiArr2(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("MultiArr2 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr2.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.MultiArr2Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.MultiArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr2.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr2 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr2.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.MultiArr2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.MultiArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr2.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr2 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr2.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.MultiArr2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.MultiArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr2.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr2 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr2.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.MultiArr2Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.MultiArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr2.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_MultiArr3(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("MultiArr3 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr3.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.MultiArr3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.MultiArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.MultiArr3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.MultiArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.MultiArr3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.MultiArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.MultiArr3Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.MultiArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr3.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_MultiArr4(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("MultiArr4 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr4.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.MultiArr4Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.MultiArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr4.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr4 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr4.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.MultiArr4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.MultiArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr4.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr4 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr4.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.MultiArr4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.MultiArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr4.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr4 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr4.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.MultiArr4Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.MultiArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr4.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_MultiArr5(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("MultiArr5 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr5.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.MultiArr5Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.MultiArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr5.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr5 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr5.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.MultiArr5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.MultiArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr5.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr5 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr5.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.MultiArr5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.MultiArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr5.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr5 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr5.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.MultiArr5Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.MultiArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr5.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptCombo1(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptCombo1 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo1.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.OptCombo1Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.OptCombo1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo1.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo1 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo1.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo1.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo1 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo1.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo1.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo1 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo1.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.OptCombo1Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.OptCombo1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo1.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptCombo2(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptCombo2 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo2.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.OptCombo2Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.OptCombo2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo2.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo2 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo2.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo2.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo2 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo2.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo2.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo2 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo2.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.OptCombo2Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.OptCombo2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo2.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptCombo3(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptCombo3 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo3.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.OptCombo3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.OptCombo3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.OptCombo3Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.OptCombo3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo3.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptCombo4(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptCombo4 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo4.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.OptCombo4Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.OptCombo4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo4.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo4 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo4.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo4.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo4 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo4.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo4.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo4 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo4.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.OptCombo4Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.OptCombo4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo4.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptCombo5(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptCombo5 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo5.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.OptCombo5Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.OptCombo5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo5.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo5 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo5.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo5.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo5 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo5.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo5.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo5 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo5.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.OptCombo5Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.OptCombo5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo5.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptCombo6(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptCombo6 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo6.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.OptCombo6Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.OptCombo6Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo6.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo6 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo6.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo6Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo6Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo6.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo6 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo6.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo6Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo6Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo6.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo6 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo6.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.OptCombo6Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.OptCombo6Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo6.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptCombo7(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptCombo7 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo7.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.OptCombo7Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.OptCombo7Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo7.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo7 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo7.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo7Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo7Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo7.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo7 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo7.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo7Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo7Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo7.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo7 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo7.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.OptCombo7Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.OptCombo7Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo7.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptCombo8(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptCombo8 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo8.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.OptCombo8Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.OptCombo8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo8.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo8 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo8.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo8.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo8 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo8.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo8.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo8 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo8.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.OptCombo8Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.OptCombo8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo8.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptCombo9(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptCombo9 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo9.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.OptCombo9Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.OptCombo9Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo9.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo9 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo9.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo9Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo9Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo9.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo9 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo9.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo9Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo9Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo9.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo9 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo9.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.OptCombo9Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.OptCombo9Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo9.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptCombo10(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptCombo10 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo10.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_arrays.OptCombo10Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_arrays.OptCombo10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo10.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo10 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo10.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo10Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo10.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo10 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo10.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_arrays.OptCombo10Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_arrays.OptCombo10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo10.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo10 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo10.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_arrays.OptCombo10Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_arrays.OptCombo10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo10.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypesArrays(vecDir, outDir string) (passed, failed int) {
	var p, f int
	p, f = runModel_ArrString(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ArrInt32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ArrBoolean(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ArrFloat64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ArrBytes(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ArrInt64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ArrUint64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_MultiArr1(vecDir, outDir); passed += p; failed += f
	p, f = runModel_MultiArr2(vecDir, outDir); passed += p; failed += f
	p, f = runModel_MultiArr3(vecDir, outDir); passed += p; failed += f
	p, f = runModel_MultiArr4(vecDir, outDir); passed += p; failed += f
	p, f = runModel_MultiArr5(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptCombo1(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptCombo2(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptCombo3(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptCombo4(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptCombo5(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptCombo6(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptCombo7(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptCombo8(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptCombo9(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptCombo10(vecDir, outDir); passed += p; failed += f

	return
}
