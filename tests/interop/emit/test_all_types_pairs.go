package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	all_types_pairs "emit_go/emit_gen/all_types_pairs"
)


func runModel_PairString(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("PairString mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairString.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.PairStringCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.PairStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairString.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairString json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairString.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairStringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairString.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairString unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairString.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairStringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairString.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairString gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairString.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.PairStringCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.PairStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairString.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_PairBoolean(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("PairBoolean mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairBoolean.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.PairBooleanCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.PairBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairBoolean.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairBoolean json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairBoolean.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairBoolean.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairBoolean unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairBoolean.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairBoolean.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairBoolean gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairBoolean.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.PairBooleanCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.PairBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairBoolean.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_PairInt8(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("PairInt8 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt8.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.PairInt8Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.PairInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt8.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt8 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt8.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairInt8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt8.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt8 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt8.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairInt8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt8.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt8 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt8.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.PairInt8Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.PairInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt8.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_PairInt16(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("PairInt16 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt16.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.PairInt16Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.PairInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt16.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt16 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt16.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairInt16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt16.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt16 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt16.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairInt16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt16.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt16 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt16.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.PairInt16Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.PairInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt16.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_PairInt32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("PairInt32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.PairInt32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.PairInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.PairInt32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.PairInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_PairInt64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("PairInt64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.PairInt64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.PairInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.PairInt64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.PairInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_PairUint8(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("PairUint8 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint8.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.PairUint8Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.PairUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint8.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint8 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint8.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairUint8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint8.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint8 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint8.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairUint8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint8.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint8 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint8.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.PairUint8Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.PairUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint8.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_PairUint16(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("PairUint16 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint16.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.PairUint16Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.PairUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint16.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint16 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint16.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairUint16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint16.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint16 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint16.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairUint16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint16.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint16 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint16.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.PairUint16Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.PairUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint16.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_PairUint32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("PairUint32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.PairUint32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.PairUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairUint32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairUint32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.PairUint32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.PairUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_PairUint64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("PairUint64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.PairUint64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.PairUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.PairUint64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.PairUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_PairFloat32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("PairFloat32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairFloat32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.PairFloat32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.PairFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairFloat32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairFloat32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairFloat32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairFloat32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairFloat32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairFloat32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairFloat32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairFloat32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairFloat32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairFloat32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairFloat32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.PairFloat32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.PairFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairFloat32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_PairFloat64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("PairFloat64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairFloat64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.PairFloat64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.PairFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairFloat64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairFloat64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairFloat64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairFloat64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairFloat64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairFloat64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairFloat64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairFloat64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairFloat64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.PairFloat64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.PairFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairFloat64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_PairBytes(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("PairBytes mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairBytes.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.PairBytesCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.PairBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairBytes.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairBytes json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairBytes.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairBytes.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairBytes unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairBytes.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.PairBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.PairBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairBytes.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairBytes gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairBytes.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.PairBytesCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.PairBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairBytes.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualStringInt32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualStringInt32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringInt32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualStringInt32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualStringInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringInt32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringInt32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringInt32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualStringInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualStringInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringInt32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringInt32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringInt32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualStringInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualStringInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringInt32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringInt32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringInt32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualStringInt32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualStringInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringInt32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualStringBoolean(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualStringBoolean mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringBoolean.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualStringBooleanCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualStringBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringBoolean.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringBoolean json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringBoolean.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualStringBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualStringBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringBoolean.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringBoolean unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringBoolean.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualStringBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualStringBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringBoolean.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringBoolean gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringBoolean.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualStringBooleanCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualStringBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringBoolean.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualStringFloat64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualStringFloat64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringFloat64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualStringFloat64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualStringFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringFloat64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringFloat64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringFloat64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualStringFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualStringFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringFloat64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringFloat64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringFloat64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualStringFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualStringFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringFloat64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringFloat64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringFloat64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualStringFloat64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualStringFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringFloat64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualStringBytes(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualStringBytes mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringBytes.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualStringBytesCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualStringBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringBytes.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringBytes json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringBytes.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualStringBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualStringBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringBytes.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringBytes unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringBytes.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualStringBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualStringBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringBytes.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringBytes gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringBytes.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualStringBytesCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualStringBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringBytes.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualInt32Boolean(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualInt32Boolean mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Boolean.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualInt32BooleanCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualInt32BooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Boolean.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Boolean json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Boolean.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt32BooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt32BooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Boolean.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Boolean unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Boolean.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt32BooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt32BooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Boolean.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Boolean gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Boolean.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualInt32BooleanCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualInt32BooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Boolean.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualInt32Float64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualInt32Float64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Float64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualInt32Float64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualInt32Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Float64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Float64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Float64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt32Float64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt32Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Float64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Float64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Float64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt32Float64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt32Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Float64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Float64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Float64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualInt32Float64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualInt32Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Float64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualInt32Int64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualInt32Int64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Int64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualInt32Int64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualInt32Int64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Int64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Int64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Int64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt32Int64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt32Int64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Int64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Int64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Int64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt32Int64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt32Int64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Int64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Int64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Int64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualInt32Int64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualInt32Int64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Int64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualInt32Uint32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualInt32Uint32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Uint32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualInt32Uint32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualInt32Uint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Uint32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Uint32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Uint32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt32Uint32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt32Uint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Uint32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Uint32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Uint32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt32Uint32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt32Uint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Uint32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Uint32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Uint32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualInt32Uint32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualInt32Uint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Uint32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualInt64Uint64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualInt64Uint64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt64Uint64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualInt64Uint64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualInt64Uint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt64Uint64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt64Uint64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt64Uint64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt64Uint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt64Uint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt64Uint64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt64Uint64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt64Uint64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt64Uint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt64Uint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt64Uint64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt64Uint64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt64Uint64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualInt64Uint64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualInt64Uint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt64Uint64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualFloat32Float64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualFloat32Float64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat32Float64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualFloat32Float64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualFloat32Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat32Float64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat32Float64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat32Float64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualFloat32Float64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualFloat32Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat32Float64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat32Float64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat32Float64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualFloat32Float64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualFloat32Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat32Float64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat32Float64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat32Float64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualFloat32Float64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualFloat32Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat32Float64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualFloat64Boolean(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualFloat64Boolean mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Boolean.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualFloat64BooleanCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualFloat64BooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Boolean.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Boolean json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Boolean.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualFloat64BooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualFloat64BooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Boolean.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Boolean unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Boolean.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualFloat64BooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualFloat64BooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Boolean.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Boolean gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Boolean.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualFloat64BooleanCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualFloat64BooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Boolean.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualFloat64Bytes(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualFloat64Bytes mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Bytes.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualFloat64BytesCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualFloat64BytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Bytes.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Bytes json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Bytes.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualFloat64BytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualFloat64BytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Bytes.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Bytes unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Bytes.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualFloat64BytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualFloat64BytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Bytes.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Bytes gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Bytes.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualFloat64BytesCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualFloat64BytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Bytes.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualUint32Uint64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualUint32Uint64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint32Uint64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualUint32Uint64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualUint32Uint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint32Uint64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint32Uint64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint32Uint64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualUint32Uint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualUint32Uint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint32Uint64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint32Uint64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint32Uint64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualUint32Uint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualUint32Uint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint32Uint64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint32Uint64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint32Uint64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualUint32Uint64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualUint32Uint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint32Uint64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualBooleanBytes(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualBooleanBytes mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBooleanBytes.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualBooleanBytesCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualBooleanBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBooleanBytes.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBooleanBytes json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBooleanBytes.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualBooleanBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualBooleanBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBooleanBytes.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBooleanBytes unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBooleanBytes.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualBooleanBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualBooleanBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBooleanBytes.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBooleanBytes gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBooleanBytes.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualBooleanBytesCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualBooleanBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBooleanBytes.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualInt8Uint8(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualInt8Uint8 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt8Uint8.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualInt8Uint8Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualInt8Uint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt8Uint8.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt8Uint8 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt8Uint8.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt8Uint8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt8Uint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt8Uint8.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt8Uint8 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt8Uint8.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt8Uint8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt8Uint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt8Uint8.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt8Uint8 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt8Uint8.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualInt8Uint8Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualInt8Uint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt8Uint8.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualInt16Uint16(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualInt16Uint16 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt16Uint16.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualInt16Uint16Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualInt16Uint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt16Uint16.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt16Uint16 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt16Uint16.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt16Uint16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt16Uint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt16Uint16.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt16Uint16 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt16Uint16.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt16Uint16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt16Uint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt16Uint16.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt16Uint16 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt16Uint16.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualInt16Uint16Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualInt16Uint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt16Uint16.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualStringInt64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualStringInt64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringInt64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualStringInt64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualStringInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringInt64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringInt64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringInt64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualStringInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualStringInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringInt64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringInt64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringInt64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualStringInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualStringInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringInt64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringInt64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringInt64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualStringInt64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualStringInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringInt64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualStringUint64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualStringUint64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringUint64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualStringUint64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualStringUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringUint64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringUint64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringUint64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualStringUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualStringUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringUint64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringUint64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringUint64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualStringUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualStringUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringUint64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringUint64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringUint64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualStringUint64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualStringUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringUint64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualInt32Bytes(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualInt32Bytes mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Bytes.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualInt32BytesCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualInt32BytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Bytes.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Bytes json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Bytes.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt32BytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt32BytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Bytes.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Bytes unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Bytes.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt32BytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt32BytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Bytes.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Bytes gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Bytes.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualInt32BytesCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualInt32BytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Bytes.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualFloat64Int32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualFloat64Int32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Int32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualFloat64Int32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualFloat64Int32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Int32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Int32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Int32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualFloat64Int32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualFloat64Int32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Int32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Int32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Int32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualFloat64Int32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualFloat64Int32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Int32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Int32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Int32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualFloat64Int32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualFloat64Int32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Int32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualBooleanInt32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualBooleanInt32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBooleanInt32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualBooleanInt32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualBooleanInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBooleanInt32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBooleanInt32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBooleanInt32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualBooleanInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualBooleanInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBooleanInt32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBooleanInt32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBooleanInt32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualBooleanInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualBooleanInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBooleanInt32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBooleanInt32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBooleanInt32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualBooleanInt32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualBooleanInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBooleanInt32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualBytesInt64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualBytesInt64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBytesInt64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualBytesInt64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualBytesInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBytesInt64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBytesInt64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBytesInt64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualBytesInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualBytesInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBytesInt64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBytesInt64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBytesInt64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualBytesInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualBytesInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBytesInt64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBytesInt64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBytesInt64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualBytesInt64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualBytesInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBytesInt64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualInt8Float32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualInt8Float32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt8Float32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualInt8Float32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualInt8Float32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt8Float32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt8Float32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt8Float32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt8Float32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt8Float32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt8Float32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt8Float32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt8Float32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt8Float32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt8Float32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt8Float32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt8Float32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt8Float32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualInt8Float32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualInt8Float32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt8Float32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualUint8Int16(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualUint8Int16 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint8Int16.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualUint8Int16Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualUint8Int16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint8Int16.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint8Int16 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint8Int16.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualUint8Int16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualUint8Int16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint8Int16.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint8Int16 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint8Int16.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualUint8Int16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualUint8Int16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint8Int16.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint8Int16 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint8Int16.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualUint8Int16Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualUint8Int16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint8Int16.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualInt64Float64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualInt64Float64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt64Float64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualInt64Float64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualInt64Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt64Float64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt64Float64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt64Float64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt64Float64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt64Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt64Float64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt64Float64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt64Float64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualInt64Float64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualInt64Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt64Float64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt64Float64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt64Float64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualInt64Float64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualInt64Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt64Float64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DualUint64String(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DualUint64String mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint64String.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.DualUint64StringCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.DualUint64StringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint64String.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint64String json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint64String.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualUint64StringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualUint64StringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint64String.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint64String unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint64String.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.DualUint64StringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.DualUint64StringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint64String.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint64String gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint64String.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.DualUint64StringCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.DualUint64StringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint64String.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Triple01(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Triple01 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple01.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.Triple01Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.Triple01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple01.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple01 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple01.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple01Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple01.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple01 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple01.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple01Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple01.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple01 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple01.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.Triple01Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.Triple01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple01.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Triple02(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Triple02 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple02.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.Triple02Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.Triple02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple02.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple02 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple02.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple02Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple02.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple02 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple02.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple02Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple02.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple02 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple02.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.Triple02Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.Triple02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple02.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Triple03(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Triple03 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple03.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.Triple03Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.Triple03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple03.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple03 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple03.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple03Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple03.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple03 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple03.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple03Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple03.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple03 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple03.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.Triple03Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.Triple03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple03.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Triple04(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Triple04 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple04.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.Triple04Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.Triple04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple04.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple04 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple04.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple04Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple04.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple04 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple04.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple04Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple04.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple04 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple04.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.Triple04Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.Triple04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple04.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Triple05(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Triple05 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple05.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.Triple05Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.Triple05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple05.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple05 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple05.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple05Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple05.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple05 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple05.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple05Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple05.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple05 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple05.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.Triple05Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.Triple05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple05.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Triple06(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Triple06 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple06.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.Triple06Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.Triple06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple06.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple06 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple06.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple06Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple06.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple06 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple06.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple06Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple06.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple06 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple06.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.Triple06Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.Triple06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple06.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Triple07(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Triple07 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple07.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.Triple07Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.Triple07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple07.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple07 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple07.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple07Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple07.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple07 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple07.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple07Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple07.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple07 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple07.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.Triple07Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.Triple07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple07.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Triple08(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Triple08 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple08.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.Triple08Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.Triple08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple08.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple08 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple08.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple08Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple08.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple08 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple08.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple08Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple08.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple08 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple08.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.Triple08Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.Triple08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple08.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Triple09(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Triple09 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple09.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.Triple09Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.Triple09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple09.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple09 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple09.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple09Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple09.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple09 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple09.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple09Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple09.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple09 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple09.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.Triple09Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.Triple09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple09.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Triple10(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Triple10 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple10.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.Triple10Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.Triple10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple10.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple10 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple10.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple10Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple10.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple10 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple10.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple10Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple10.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple10 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple10.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.Triple10Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.Triple10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple10.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Triple11(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Triple11 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple11.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.Triple11Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.Triple11Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple11.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple11 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple11.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple11Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple11Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple11.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple11 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple11.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple11Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple11Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple11.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple11 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple11.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.Triple11Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.Triple11Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple11.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Triple12(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Triple12 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple12.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.Triple12Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.Triple12Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple12.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple12 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple12.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple12Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple12Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple12.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple12 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple12.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple12Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple12Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple12.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple12 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple12.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.Triple12Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.Triple12Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple12.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Triple13(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Triple13 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple13.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.Triple13Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.Triple13Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple13.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple13 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple13.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple13Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple13Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple13.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple13 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple13.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple13Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple13Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple13.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple13 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple13.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.Triple13Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.Triple13Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple13.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Triple14(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Triple14 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple14.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.Triple14Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.Triple14Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple14.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple14 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple14.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple14Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple14Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple14.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple14 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple14.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple14Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple14Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple14.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple14 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple14.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.Triple14Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.Triple14Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple14.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Triple15(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Triple15 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple15.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_pairs.Triple15Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_pairs.Triple15Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple15.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple15 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple15.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple15Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple15Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple15.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple15 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple15.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_pairs.Triple15Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_pairs.Triple15Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple15.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple15 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple15.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_pairs.Triple15Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_pairs.Triple15Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple15.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypesPairs(vecDir, outDir string) (passed, failed int) {
	var p, f int
	p, f = runModel_PairString(vecDir, outDir); passed += p; failed += f
	p, f = runModel_PairBoolean(vecDir, outDir); passed += p; failed += f
	p, f = runModel_PairInt8(vecDir, outDir); passed += p; failed += f
	p, f = runModel_PairInt16(vecDir, outDir); passed += p; failed += f
	p, f = runModel_PairInt32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_PairInt64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_PairUint8(vecDir, outDir); passed += p; failed += f
	p, f = runModel_PairUint16(vecDir, outDir); passed += p; failed += f
	p, f = runModel_PairUint32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_PairUint64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_PairFloat32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_PairFloat64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_PairBytes(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualStringInt32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualStringBoolean(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualStringFloat64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualStringBytes(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualInt32Boolean(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualInt32Float64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualInt32Int64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualInt32Uint32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualInt64Uint64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualFloat32Float64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualFloat64Boolean(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualFloat64Bytes(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualUint32Uint64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualBooleanBytes(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualInt8Uint8(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualInt16Uint16(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualStringInt64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualStringUint64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualInt32Bytes(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualFloat64Int32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualBooleanInt32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualBytesInt64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualInt8Float32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualUint8Int16(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualInt64Float64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DualUint64String(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Triple01(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Triple02(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Triple03(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Triple04(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Triple05(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Triple06(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Triple07(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Triple08(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Triple09(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Triple10(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Triple11(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Triple12(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Triple13(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Triple14(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Triple15(vecDir, outDir); passed += p; failed += f

	return
}
