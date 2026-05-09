package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	all_types_scalars "emit_go/emit_gen/all_types_scalars"
)


func runModel_SingleString(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("SingleString mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleString.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_scalars.SingleStringCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_scalars.SingleStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleString.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleString json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleString.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleStringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleString.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleString unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleString.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleStringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleString.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleString gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleString.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_scalars.SingleStringCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_scalars.SingleStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleString.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_SingleBoolean(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("SingleBoolean mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleBoolean.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_scalars.SingleBooleanCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_scalars.SingleBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleBoolean.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleBoolean json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleBoolean.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleBoolean.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleBoolean unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleBoolean.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleBoolean.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleBoolean gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleBoolean.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_scalars.SingleBooleanCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_scalars.SingleBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleBoolean.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_SingleInt8(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("SingleInt8 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt8.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_scalars.SingleInt8Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_scalars.SingleInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt8.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt8 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt8.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleInt8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt8.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt8 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt8.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleInt8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt8.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt8 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt8.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_scalars.SingleInt8Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_scalars.SingleInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt8.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_SingleInt16(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("SingleInt16 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt16.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_scalars.SingleInt16Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_scalars.SingleInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt16.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt16 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt16.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleInt16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt16.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt16 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt16.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleInt16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt16.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt16 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt16.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_scalars.SingleInt16Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_scalars.SingleInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt16.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_SingleInt32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("SingleInt32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_scalars.SingleInt32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_scalars.SingleInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_scalars.SingleInt32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_scalars.SingleInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_SingleInt64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("SingleInt64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_scalars.SingleInt64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_scalars.SingleInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_scalars.SingleInt64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_scalars.SingleInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_SingleUint8(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("SingleUint8 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint8.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_scalars.SingleUint8Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_scalars.SingleUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint8.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint8 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint8.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleUint8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint8.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint8 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint8.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleUint8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint8.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint8 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint8.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_scalars.SingleUint8Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_scalars.SingleUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint8.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_SingleUint16(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("SingleUint16 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint16.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_scalars.SingleUint16Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_scalars.SingleUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint16.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint16 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint16.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleUint16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint16.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint16 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint16.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleUint16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint16.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint16 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint16.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_scalars.SingleUint16Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_scalars.SingleUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint16.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_SingleUint32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("SingleUint32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_scalars.SingleUint32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_scalars.SingleUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleUint32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleUint32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_scalars.SingleUint32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_scalars.SingleUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_SingleUint64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("SingleUint64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_scalars.SingleUint64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_scalars.SingleUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_scalars.SingleUint64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_scalars.SingleUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_SingleFloat32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("SingleFloat32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleFloat32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_scalars.SingleFloat32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_scalars.SingleFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleFloat32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleFloat32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleFloat32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleFloat32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleFloat32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleFloat32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleFloat32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleFloat32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleFloat32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleFloat32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleFloat32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_scalars.SingleFloat32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_scalars.SingleFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleFloat32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_SingleFloat64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("SingleFloat64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleFloat64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_scalars.SingleFloat64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_scalars.SingleFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleFloat64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleFloat64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleFloat64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleFloat64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleFloat64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleFloat64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleFloat64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleFloat64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleFloat64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_scalars.SingleFloat64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_scalars.SingleFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleFloat64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_SingleBytes(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("SingleBytes mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleBytes.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_scalars.SingleBytesCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_scalars.SingleBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleBytes.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleBytes json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleBytes.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleBytes.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleBytes unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleBytes.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_scalars.SingleBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_scalars.SingleBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleBytes.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleBytes gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleBytes.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_scalars.SingleBytesCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_scalars.SingleBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleBytes.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypesScalars(vecDir, outDir string) (passed, failed int) {
	var p, f int
	p, f = runModel_SingleString(vecDir, outDir); passed += p; failed += f
	p, f = runModel_SingleBoolean(vecDir, outDir); passed += p; failed += f
	p, f = runModel_SingleInt8(vecDir, outDir); passed += p; failed += f
	p, f = runModel_SingleInt16(vecDir, outDir); passed += p; failed += f
	p, f = runModel_SingleInt32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_SingleInt64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_SingleUint8(vecDir, outDir); passed += p; failed += f
	p, f = runModel_SingleUint16(vecDir, outDir); passed += p; failed += f
	p, f = runModel_SingleUint32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_SingleUint64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_SingleFloat32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_SingleFloat64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_SingleBytes(vecDir, outDir); passed += p; failed += f

	return
}
