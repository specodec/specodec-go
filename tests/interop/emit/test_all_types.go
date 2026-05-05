package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	specodec_all_types "emit_go/emit_gen/specodec_all_types"
	specodec_all_types_nested "emit_go/emit_gen/specodec_all_types_nested"
	specodec_all_types_nested_deep "emit_go/emit_gen/specodec_all_types_nested_deep"
)


func runModel_OptInner(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptInner mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptInner.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptInnerCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptInner.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptInner json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptInner.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptInnerCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptInner.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptInner unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptInner.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptInnerCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptInner.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptInner gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptInner.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptInnerCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptInner.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_SingleString(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("SingleString mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleString.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.SingleStringCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.SingleStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleString.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleString json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleString.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleStringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleString.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleString unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleString.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleStringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleString.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleString gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleString.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.SingleStringCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.SingleStringCodec.Encode(w, obj)
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
		obj := specodec_all_types.SingleBooleanCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.SingleBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleBoolean.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleBoolean json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleBoolean.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleBoolean.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleBoolean unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleBoolean.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleBoolean.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleBoolean gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleBoolean.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.SingleBooleanCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.SingleBooleanCodec.Encode(w, obj)
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
		obj := specodec_all_types.SingleInt8Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.SingleInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt8.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt8 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt8.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleInt8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt8.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt8 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt8.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleInt8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt8.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt8 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt8.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.SingleInt8Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.SingleInt8Codec.Encode(w, obj)
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
		obj := specodec_all_types.SingleInt16Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.SingleInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt16.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt16 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt16.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleInt16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt16.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt16 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt16.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleInt16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt16.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt16 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt16.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.SingleInt16Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.SingleInt16Codec.Encode(w, obj)
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
		obj := specodec_all_types.SingleInt32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.SingleInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.SingleInt32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.SingleInt32Codec.Encode(w, obj)
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
		obj := specodec_all_types.SingleInt64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.SingleInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleInt64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleInt64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleInt64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.SingleInt64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.SingleInt64Codec.Encode(w, obj)
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
		obj := specodec_all_types.SingleUint8Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.SingleUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint8.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint8 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint8.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleUint8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint8.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint8 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint8.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleUint8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint8.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint8 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint8.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.SingleUint8Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.SingleUint8Codec.Encode(w, obj)
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
		obj := specodec_all_types.SingleUint16Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.SingleUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint16.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint16 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint16.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleUint16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint16.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint16 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint16.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleUint16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint16.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint16 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint16.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.SingleUint16Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.SingleUint16Codec.Encode(w, obj)
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
		obj := specodec_all_types.SingleUint32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.SingleUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleUint32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleUint32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.SingleUint32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.SingleUint32Codec.Encode(w, obj)
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
		obj := specodec_all_types.SingleUint64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.SingleUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleUint64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleUint64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleUint64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.SingleUint64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.SingleUint64Codec.Encode(w, obj)
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
		obj := specodec_all_types.SingleFloat32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.SingleFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleFloat32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleFloat32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleFloat32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleFloat32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleFloat32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleFloat32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleFloat32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleFloat32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleFloat32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleFloat32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleFloat32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.SingleFloat32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.SingleFloat32Codec.Encode(w, obj)
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
		obj := specodec_all_types.SingleFloat64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.SingleFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleFloat64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleFloat64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleFloat64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleFloat64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleFloat64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleFloat64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleFloat64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleFloat64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleFloat64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.SingleFloat64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.SingleFloat64Codec.Encode(w, obj)
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
		obj := specodec_all_types.SingleBytesCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.SingleBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleBytes.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleBytes json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleBytes.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleBytes.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleBytes unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleBytes.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.SingleBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.SingleBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleBytes.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("SingleBytes gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "SingleBytes.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.SingleBytesCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.SingleBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "SingleBytes.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptSingleString(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptSingleString mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleString.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptSingleStringCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptSingleStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleString.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleString json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleString.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleStringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleString.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleString unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleString.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleStringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleString.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleString gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleString.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptSingleStringCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptSingleStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleString.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptSingleBoolean(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptSingleBoolean mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleBoolean.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptSingleBooleanCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptSingleBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleBoolean.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleBoolean json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleBoolean.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleBoolean.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleBoolean unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleBoolean.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleBoolean.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleBoolean gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleBoolean.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptSingleBooleanCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptSingleBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleBoolean.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptSingleInt8(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptSingleInt8 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt8.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptSingleInt8Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptSingleInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt8.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt8 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt8.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleInt8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt8.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt8 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt8.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleInt8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt8.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt8 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt8.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptSingleInt8Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptSingleInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt8.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptSingleInt16(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptSingleInt16 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt16.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptSingleInt16Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptSingleInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt16.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt16 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt16.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleInt16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt16.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt16 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt16.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleInt16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt16.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt16 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt16.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptSingleInt16Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptSingleInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt16.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptSingleInt32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptSingleInt32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptSingleInt32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptSingleInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptSingleInt32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptSingleInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptSingleInt64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptSingleInt64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptSingleInt64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptSingleInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptSingleInt64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptSingleInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptSingleUint8(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptSingleUint8 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint8.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptSingleUint8Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptSingleUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint8.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint8 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint8.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleUint8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint8.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint8 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint8.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleUint8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint8.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint8 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint8.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptSingleUint8Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptSingleUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint8.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptSingleUint16(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptSingleUint16 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint16.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptSingleUint16Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptSingleUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint16.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint16 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint16.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleUint16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint16.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint16 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint16.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleUint16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint16.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint16 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint16.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptSingleUint16Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptSingleUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint16.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptSingleUint32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptSingleUint32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptSingleUint32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptSingleUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleUint32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleUint32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptSingleUint32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptSingleUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptSingleUint64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptSingleUint64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptSingleUint64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptSingleUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptSingleUint64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptSingleUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptSingleFloat32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptSingleFloat32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleFloat32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptSingleFloat32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptSingleFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleFloat32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleFloat32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleFloat32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleFloat32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleFloat32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleFloat32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleFloat32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleFloat32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleFloat32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleFloat32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleFloat32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptSingleFloat32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptSingleFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleFloat32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptSingleFloat64(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptSingleFloat64 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleFloat64.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptSingleFloat64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptSingleFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleFloat64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleFloat64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleFloat64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleFloat64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleFloat64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleFloat64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleFloat64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleFloat64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleFloat64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptSingleFloat64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptSingleFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleFloat64.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptSingleBytes(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptSingleBytes mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleBytes.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptSingleBytesCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptSingleBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleBytes.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleBytes json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleBytes.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleBytes.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleBytes unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleBytes.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptSingleBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptSingleBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleBytes.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleBytes gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleBytes.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptSingleBytesCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptSingleBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleBytes.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_PairString(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("PairString mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairString.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.PairStringCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.PairStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairString.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairString json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairString.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairStringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairString.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairString unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairString.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairStringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairString.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairString gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairString.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.PairStringCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.PairStringCodec.Encode(w, obj)
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
		obj := specodec_all_types.PairBooleanCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.PairBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairBoolean.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairBoolean json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairBoolean.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairBoolean.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairBoolean unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairBoolean.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairBoolean.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairBoolean gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairBoolean.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.PairBooleanCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.PairBooleanCodec.Encode(w, obj)
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
		obj := specodec_all_types.PairInt8Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.PairInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt8.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt8 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt8.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairInt8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt8.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt8 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt8.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairInt8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt8.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt8 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt8.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.PairInt8Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.PairInt8Codec.Encode(w, obj)
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
		obj := specodec_all_types.PairInt16Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.PairInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt16.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt16 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt16.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairInt16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt16.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt16 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt16.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairInt16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt16.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt16 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt16.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.PairInt16Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.PairInt16Codec.Encode(w, obj)
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
		obj := specodec_all_types.PairInt32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.PairInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.PairInt32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.PairInt32Codec.Encode(w, obj)
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
		obj := specodec_all_types.PairInt64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.PairInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairInt64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairInt64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairInt64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.PairInt64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.PairInt64Codec.Encode(w, obj)
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
		obj := specodec_all_types.PairUint8Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.PairUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint8.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint8 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint8.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairUint8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint8.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint8 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint8.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairUint8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint8.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint8 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint8.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.PairUint8Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.PairUint8Codec.Encode(w, obj)
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
		obj := specodec_all_types.PairUint16Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.PairUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint16.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint16 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint16.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairUint16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint16.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint16 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint16.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairUint16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint16.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint16 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint16.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.PairUint16Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.PairUint16Codec.Encode(w, obj)
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
		obj := specodec_all_types.PairUint32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.PairUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairUint32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairUint32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.PairUint32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.PairUint32Codec.Encode(w, obj)
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
		obj := specodec_all_types.PairUint64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.PairUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairUint64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairUint64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairUint64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.PairUint64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.PairUint64Codec.Encode(w, obj)
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
		obj := specodec_all_types.PairFloat32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.PairFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairFloat32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairFloat32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairFloat32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairFloat32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairFloat32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairFloat32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairFloat32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairFloat32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairFloat32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairFloat32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairFloat32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.PairFloat32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.PairFloat32Codec.Encode(w, obj)
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
		obj := specodec_all_types.PairFloat64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.PairFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairFloat64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairFloat64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairFloat64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairFloat64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairFloat64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairFloat64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairFloat64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairFloat64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairFloat64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.PairFloat64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.PairFloat64Codec.Encode(w, obj)
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
		obj := specodec_all_types.PairBytesCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.PairBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairBytes.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairBytes json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairBytes.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairBytes.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairBytes unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairBytes.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.PairBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.PairBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "PairBytes.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("PairBytes gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "PairBytes.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.PairBytesCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.PairBytesCodec.Encode(w, obj)
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
		obj := specodec_all_types.DualStringInt32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualStringInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringInt32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringInt32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringInt32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualStringInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualStringInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringInt32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringInt32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringInt32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualStringInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualStringInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringInt32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringInt32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringInt32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualStringInt32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualStringInt32Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualStringBooleanCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualStringBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringBoolean.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringBoolean json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringBoolean.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualStringBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualStringBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringBoolean.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringBoolean unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringBoolean.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualStringBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualStringBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringBoolean.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringBoolean gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringBoolean.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualStringBooleanCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualStringBooleanCodec.Encode(w, obj)
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
		obj := specodec_all_types.DualStringFloat64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualStringFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringFloat64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringFloat64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringFloat64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualStringFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualStringFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringFloat64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringFloat64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringFloat64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualStringFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualStringFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringFloat64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringFloat64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringFloat64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualStringFloat64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualStringFloat64Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualStringBytesCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualStringBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringBytes.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringBytes json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringBytes.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualStringBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualStringBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringBytes.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringBytes unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringBytes.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualStringBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualStringBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringBytes.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringBytes gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringBytes.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualStringBytesCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualStringBytesCodec.Encode(w, obj)
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
		obj := specodec_all_types.DualInt32BooleanCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualInt32BooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Boolean.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Boolean json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Boolean.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt32BooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt32BooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Boolean.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Boolean unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Boolean.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt32BooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt32BooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Boolean.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Boolean gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Boolean.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualInt32BooleanCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualInt32BooleanCodec.Encode(w, obj)
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
		obj := specodec_all_types.DualInt32Float64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualInt32Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Float64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Float64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Float64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt32Float64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt32Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Float64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Float64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Float64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt32Float64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt32Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Float64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Float64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Float64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualInt32Float64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualInt32Float64Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualInt32Int64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualInt32Int64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Int64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Int64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Int64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt32Int64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt32Int64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Int64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Int64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Int64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt32Int64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt32Int64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Int64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Int64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Int64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualInt32Int64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualInt32Int64Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualInt32Uint32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualInt32Uint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Uint32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Uint32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Uint32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt32Uint32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt32Uint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Uint32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Uint32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Uint32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt32Uint32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt32Uint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Uint32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Uint32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Uint32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualInt32Uint32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualInt32Uint32Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualInt64Uint64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualInt64Uint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt64Uint64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt64Uint64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt64Uint64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt64Uint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt64Uint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt64Uint64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt64Uint64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt64Uint64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt64Uint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt64Uint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt64Uint64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt64Uint64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt64Uint64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualInt64Uint64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualInt64Uint64Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualFloat32Float64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualFloat32Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat32Float64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat32Float64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat32Float64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualFloat32Float64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualFloat32Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat32Float64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat32Float64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat32Float64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualFloat32Float64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualFloat32Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat32Float64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat32Float64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat32Float64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualFloat32Float64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualFloat32Float64Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualFloat64BooleanCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualFloat64BooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Boolean.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Boolean json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Boolean.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualFloat64BooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualFloat64BooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Boolean.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Boolean unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Boolean.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualFloat64BooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualFloat64BooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Boolean.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Boolean gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Boolean.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualFloat64BooleanCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualFloat64BooleanCodec.Encode(w, obj)
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
		obj := specodec_all_types.DualFloat64BytesCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualFloat64BytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Bytes.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Bytes json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Bytes.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualFloat64BytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualFloat64BytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Bytes.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Bytes unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Bytes.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualFloat64BytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualFloat64BytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Bytes.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Bytes gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Bytes.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualFloat64BytesCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualFloat64BytesCodec.Encode(w, obj)
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
		obj := specodec_all_types.DualUint32Uint64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualUint32Uint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint32Uint64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint32Uint64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint32Uint64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualUint32Uint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualUint32Uint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint32Uint64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint32Uint64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint32Uint64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualUint32Uint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualUint32Uint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint32Uint64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint32Uint64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint32Uint64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualUint32Uint64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualUint32Uint64Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualBooleanBytesCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualBooleanBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBooleanBytes.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBooleanBytes json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBooleanBytes.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualBooleanBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualBooleanBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBooleanBytes.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBooleanBytes unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBooleanBytes.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualBooleanBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualBooleanBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBooleanBytes.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBooleanBytes gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBooleanBytes.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualBooleanBytesCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualBooleanBytesCodec.Encode(w, obj)
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
		obj := specodec_all_types.DualInt8Uint8Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualInt8Uint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt8Uint8.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt8Uint8 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt8Uint8.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt8Uint8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt8Uint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt8Uint8.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt8Uint8 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt8Uint8.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt8Uint8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt8Uint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt8Uint8.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt8Uint8 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt8Uint8.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualInt8Uint8Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualInt8Uint8Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualInt16Uint16Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualInt16Uint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt16Uint16.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt16Uint16 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt16Uint16.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt16Uint16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt16Uint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt16Uint16.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt16Uint16 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt16Uint16.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt16Uint16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt16Uint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt16Uint16.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt16Uint16 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt16Uint16.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualInt16Uint16Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualInt16Uint16Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualStringInt64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualStringInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringInt64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringInt64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringInt64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualStringInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualStringInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringInt64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringInt64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringInt64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualStringInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualStringInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringInt64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringInt64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringInt64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualStringInt64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualStringInt64Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualStringUint64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualStringUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringUint64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringUint64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringUint64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualStringUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualStringUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringUint64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringUint64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringUint64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualStringUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualStringUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualStringUint64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualStringUint64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualStringUint64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualStringUint64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualStringUint64Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualInt32BytesCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualInt32BytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Bytes.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Bytes json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Bytes.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt32BytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt32BytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Bytes.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Bytes unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Bytes.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt32BytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt32BytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt32Bytes.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt32Bytes gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt32Bytes.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualInt32BytesCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualInt32BytesCodec.Encode(w, obj)
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
		obj := specodec_all_types.DualFloat64Int32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualFloat64Int32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Int32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Int32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Int32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualFloat64Int32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualFloat64Int32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Int32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Int32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Int32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualFloat64Int32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualFloat64Int32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualFloat64Int32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualFloat64Int32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualFloat64Int32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualFloat64Int32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualFloat64Int32Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualBooleanInt32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualBooleanInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBooleanInt32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBooleanInt32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBooleanInt32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualBooleanInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualBooleanInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBooleanInt32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBooleanInt32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBooleanInt32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualBooleanInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualBooleanInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBooleanInt32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBooleanInt32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBooleanInt32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualBooleanInt32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualBooleanInt32Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualBytesInt64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualBytesInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBytesInt64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBytesInt64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBytesInt64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualBytesInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualBytesInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBytesInt64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBytesInt64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBytesInt64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualBytesInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualBytesInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualBytesInt64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualBytesInt64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualBytesInt64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualBytesInt64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualBytesInt64Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualInt8Float32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualInt8Float32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt8Float32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt8Float32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt8Float32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt8Float32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt8Float32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt8Float32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt8Float32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt8Float32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt8Float32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt8Float32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt8Float32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt8Float32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt8Float32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualInt8Float32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualInt8Float32Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualUint8Int16Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualUint8Int16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint8Int16.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint8Int16 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint8Int16.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualUint8Int16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualUint8Int16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint8Int16.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint8Int16 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint8Int16.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualUint8Int16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualUint8Int16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint8Int16.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint8Int16 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint8Int16.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualUint8Int16Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualUint8Int16Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualInt64Float64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualInt64Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt64Float64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt64Float64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt64Float64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt64Float64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt64Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt64Float64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt64Float64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt64Float64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualInt64Float64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualInt64Float64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualInt64Float64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualInt64Float64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualInt64Float64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualInt64Float64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualInt64Float64Codec.Encode(w, obj)
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
		obj := specodec_all_types.DualUint64StringCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DualUint64StringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint64String.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint64String json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint64String.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualUint64StringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualUint64StringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint64String.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint64String unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint64String.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DualUint64StringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DualUint64StringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DualUint64String.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DualUint64String gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DualUint64String.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DualUint64StringCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DualUint64StringCodec.Encode(w, obj)
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
		obj := specodec_all_types.Triple01Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Triple01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple01.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple01 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple01.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple01Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple01.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple01 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple01.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple01Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple01.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple01 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple01.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Triple01Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Triple01Codec.Encode(w, obj)
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
		obj := specodec_all_types.Triple02Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Triple02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple02.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple02 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple02.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple02Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple02.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple02 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple02.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple02Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple02.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple02 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple02.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Triple02Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Triple02Codec.Encode(w, obj)
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
		obj := specodec_all_types.Triple03Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Triple03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple03.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple03 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple03.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple03Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple03.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple03 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple03.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple03Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple03.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple03 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple03.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Triple03Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Triple03Codec.Encode(w, obj)
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
		obj := specodec_all_types.Triple04Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Triple04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple04.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple04 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple04.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple04Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple04.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple04 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple04.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple04Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple04.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple04 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple04.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Triple04Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Triple04Codec.Encode(w, obj)
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
		obj := specodec_all_types.Triple05Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Triple05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple05.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple05 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple05.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple05Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple05.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple05 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple05.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple05Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple05.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple05 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple05.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Triple05Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Triple05Codec.Encode(w, obj)
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
		obj := specodec_all_types.Triple06Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Triple06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple06.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple06 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple06.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple06Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple06.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple06 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple06.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple06Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple06.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple06 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple06.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Triple06Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Triple06Codec.Encode(w, obj)
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
		obj := specodec_all_types.Triple07Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Triple07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple07.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple07 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple07.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple07Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple07.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple07 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple07.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple07Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple07.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple07 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple07.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Triple07Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Triple07Codec.Encode(w, obj)
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
		obj := specodec_all_types.Triple08Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Triple08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple08.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple08 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple08.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple08Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple08.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple08 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple08.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple08Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple08.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple08 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple08.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Triple08Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Triple08Codec.Encode(w, obj)
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
		obj := specodec_all_types.Triple09Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Triple09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple09.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple09 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple09.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple09Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple09.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple09 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple09.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple09Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple09.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple09 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple09.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Triple09Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Triple09Codec.Encode(w, obj)
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
		obj := specodec_all_types.Triple10Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Triple10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple10.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple10 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple10.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple10Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple10.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple10 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple10.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple10Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple10.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple10 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple10.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Triple10Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Triple10Codec.Encode(w, obj)
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
		obj := specodec_all_types.Triple11Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Triple11Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple11.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple11 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple11.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple11Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple11Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple11.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple11 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple11.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple11Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple11Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple11.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple11 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple11.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Triple11Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Triple11Codec.Encode(w, obj)
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
		obj := specodec_all_types.Triple12Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Triple12Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple12.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple12 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple12.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple12Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple12Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple12.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple12 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple12.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple12Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple12Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple12.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple12 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple12.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Triple12Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Triple12Codec.Encode(w, obj)
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
		obj := specodec_all_types.Triple13Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Triple13Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple13.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple13 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple13.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple13Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple13Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple13.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple13 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple13.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple13Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple13Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple13.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple13 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple13.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Triple13Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Triple13Codec.Encode(w, obj)
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
		obj := specodec_all_types.Triple14Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Triple14Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple14.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple14 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple14.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple14Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple14Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple14.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple14 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple14.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple14Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple14Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple14.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple14 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple14.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Triple14Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Triple14Codec.Encode(w, obj)
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
		obj := specodec_all_types.Triple15Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Triple15Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple15.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple15 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple15.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple15Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple15Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple15.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple15 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple15.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Triple15Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Triple15Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple15.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Triple15 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Triple15.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Triple15Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Triple15Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Triple15.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Five01(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Five01 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five01.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Five01Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Five01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five01.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five01 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five01.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five01Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five01.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five01 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five01.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five01Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five01.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five01 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five01.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Five01Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Five01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five01.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Five02(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Five02 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five02.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Five02Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Five02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five02.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five02 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five02.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five02Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five02.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five02 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five02.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five02Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five02.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five02 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five02.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Five02Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Five02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five02.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Five03(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Five03 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five03.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Five03Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Five03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five03.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five03 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five03.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five03Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five03.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five03 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five03.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five03Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five03.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five03 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five03.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Five03Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Five03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five03.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Five04(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Five04 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five04.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Five04Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Five04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five04.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five04 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five04.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five04Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five04.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five04 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five04.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five04Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five04.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five04 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five04.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Five04Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Five04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five04.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Five05(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Five05 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five05.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Five05Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Five05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five05.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five05 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five05.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five05Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five05.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five05 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five05.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five05Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five05.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five05 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five05.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Five05Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Five05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five05.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Five06(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Five06 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five06.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Five06Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Five06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five06.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five06 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five06.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five06Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five06.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five06 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five06.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five06Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five06.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five06 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five06.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Five06Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Five06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five06.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Five07(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Five07 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five07.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Five07Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Five07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five07.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five07 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five07.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five07Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five07.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five07 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five07.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five07Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five07.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five07 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five07.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Five07Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Five07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five07.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Five08(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Five08 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five08.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Five08Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Five08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five08.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five08 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five08.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five08Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five08.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five08 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five08.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five08Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five08.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five08 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five08.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Five08Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Five08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five08.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Five09(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Five09 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five09.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Five09Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Five09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five09.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five09 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five09.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five09Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five09.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five09 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five09.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five09Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five09.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five09 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five09.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Five09Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Five09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five09.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Five10(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Five10 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five10.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Five10Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Five10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five10.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five10 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five10.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five10Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five10.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five10 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five10.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Five10Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Five10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five10.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five10 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five10.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Five10Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Five10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five10.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Ten01(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Ten01 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten01.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Ten01Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Ten01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten01.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten01 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten01.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Ten01Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Ten01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten01.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten01 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten01.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Ten01Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Ten01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten01.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten01 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten01.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Ten01Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Ten01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten01.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Ten02(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Ten02 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten02.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Ten02Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Ten02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten02.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten02 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten02.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Ten02Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Ten02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten02.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten02 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten02.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Ten02Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Ten02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten02.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten02 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten02.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Ten02Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Ten02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten02.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Ten03(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Ten03 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten03.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Ten03Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Ten03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten03.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten03 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten03.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Ten03Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Ten03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten03.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten03 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten03.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Ten03Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Ten03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten03.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten03 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten03.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Ten03Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Ten03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten03.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Ten04(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Ten04 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten04.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Ten04Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Ten04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten04.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten04 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten04.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Ten04Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Ten04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten04.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten04 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten04.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Ten04Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Ten04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten04.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten04 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten04.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Ten04Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Ten04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten04.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Ten05(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Ten05 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten05.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Ten05Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Ten05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten05.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten05 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten05.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Ten05Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Ten05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten05.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten05 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten05.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Ten05Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Ten05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten05.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten05 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten05.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Ten05Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Ten05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten05.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ArrString(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ArrString mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrString.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.ArrStringCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.ArrStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrString.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrString json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrString.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ArrStringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ArrStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrString.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrString unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrString.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ArrStringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ArrStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrString.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrString gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrString.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.ArrStringCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.ArrStringCodec.Encode(w, obj)
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
		obj := specodec_all_types.ArrInt32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.ArrInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrInt32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrInt32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrInt32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ArrInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ArrInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrInt32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrInt32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrInt32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ArrInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ArrInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrInt32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrInt32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrInt32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.ArrInt32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.ArrInt32Codec.Encode(w, obj)
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
		obj := specodec_all_types.ArrBooleanCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.ArrBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrBoolean.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrBoolean json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrBoolean.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ArrBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ArrBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrBoolean.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrBoolean unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrBoolean.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ArrBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ArrBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrBoolean.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrBoolean gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrBoolean.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.ArrBooleanCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.ArrBooleanCodec.Encode(w, obj)
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
		obj := specodec_all_types.ArrFloat64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.ArrFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrFloat64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrFloat64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrFloat64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ArrFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ArrFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrFloat64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrFloat64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrFloat64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ArrFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ArrFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrFloat64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrFloat64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrFloat64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.ArrFloat64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.ArrFloat64Codec.Encode(w, obj)
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
		obj := specodec_all_types.ArrBytesCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.ArrBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrBytes.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrBytes json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrBytes.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ArrBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ArrBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrBytes.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrBytes unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrBytes.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ArrBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ArrBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrBytes.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrBytes gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrBytes.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.ArrBytesCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.ArrBytesCodec.Encode(w, obj)
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
		obj := specodec_all_types.ArrInt64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.ArrInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrInt64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrInt64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrInt64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ArrInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ArrInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrInt64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrInt64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrInt64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ArrInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ArrInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrInt64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrInt64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrInt64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.ArrInt64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.ArrInt64Codec.Encode(w, obj)
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
		obj := specodec_all_types.ArrUint64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.ArrUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrUint64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrUint64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrUint64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ArrUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ArrUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrUint64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrUint64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrUint64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ArrUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ArrUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ArrUint64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ArrUint64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ArrUint64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.ArrUint64Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.ArrUint64Codec.Encode(w, obj)
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
		obj := specodec_all_types.MultiArr1Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.MultiArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr1.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr1 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr1.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.MultiArr1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.MultiArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr1.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr1 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr1.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.MultiArr1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.MultiArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr1.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr1 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr1.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.MultiArr1Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.MultiArr1Codec.Encode(w, obj)
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
		obj := specodec_all_types.MultiArr2Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.MultiArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr2.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr2 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr2.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.MultiArr2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.MultiArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr2.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr2 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr2.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.MultiArr2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.MultiArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr2.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr2 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr2.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.MultiArr2Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.MultiArr2Codec.Encode(w, obj)
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
		obj := specodec_all_types.MultiArr3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.MultiArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.MultiArr3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.MultiArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.MultiArr3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.MultiArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.MultiArr3Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.MultiArr3Codec.Encode(w, obj)
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
		obj := specodec_all_types.MultiArr4Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.MultiArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr4.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr4 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr4.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.MultiArr4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.MultiArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr4.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr4 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr4.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.MultiArr4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.MultiArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr4.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr4 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr4.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.MultiArr4Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.MultiArr4Codec.Encode(w, obj)
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
		obj := specodec_all_types.MultiArr5Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.MultiArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr5.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr5 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr5.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.MultiArr5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.MultiArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr5.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr5 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr5.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.MultiArr5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.MultiArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MultiArr5.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MultiArr5 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MultiArr5.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.MultiArr5Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.MultiArr5Codec.Encode(w, obj)
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
		obj := specodec_all_types.OptCombo1Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptCombo1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo1.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo1 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo1.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo1.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo1 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo1.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo1.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo1 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo1.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptCombo1Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptCombo1Codec.Encode(w, obj)
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
		obj := specodec_all_types.OptCombo2Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptCombo2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo2.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo2 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo2.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo2.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo2 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo2.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo2.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo2 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo2.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptCombo2Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptCombo2Codec.Encode(w, obj)
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
		obj := specodec_all_types.OptCombo3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptCombo3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptCombo3Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptCombo3Codec.Encode(w, obj)
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
		obj := specodec_all_types.OptCombo4Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptCombo4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo4.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo4 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo4.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo4.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo4 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo4.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo4.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo4 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo4.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptCombo4Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptCombo4Codec.Encode(w, obj)
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
		obj := specodec_all_types.OptCombo5Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptCombo5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo5.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo5 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo5.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo5.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo5 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo5.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo5.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo5 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo5.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptCombo5Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptCombo5Codec.Encode(w, obj)
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
		obj := specodec_all_types.OptCombo6Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptCombo6Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo6.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo6 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo6.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo6Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo6Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo6.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo6 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo6.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo6Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo6Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo6.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo6 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo6.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptCombo6Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptCombo6Codec.Encode(w, obj)
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
		obj := specodec_all_types.OptCombo7Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptCombo7Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo7.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo7 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo7.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo7Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo7Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo7.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo7 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo7.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo7Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo7Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo7.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo7 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo7.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptCombo7Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptCombo7Codec.Encode(w, obj)
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
		obj := specodec_all_types.OptCombo8Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptCombo8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo8.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo8 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo8.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo8.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo8 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo8.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo8.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo8 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo8.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptCombo8Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptCombo8Codec.Encode(w, obj)
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
		obj := specodec_all_types.OptCombo9Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptCombo9Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo9.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo9 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo9.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo9Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo9Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo9.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo9 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo9.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo9Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo9Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo9.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo9 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo9.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptCombo9Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptCombo9Codec.Encode(w, obj)
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
		obj := specodec_all_types.OptCombo10Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptCombo10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo10.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo10 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo10.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo10Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo10.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo10 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo10.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptCombo10Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptCombo10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo10.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptCombo10 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptCombo10.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptCombo10Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptCombo10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptCombo10.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestInner(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestInner mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestInner.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.NestInnerCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.NestInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestInner.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestInner json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestInner.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestInnerCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestInner.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestInner unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestInner.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestInnerCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestInner.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestInner gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestInner.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.NestInnerCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.NestInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestInner.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestCoord(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestCoord mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestCoord.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.NestCoordCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.NestCoordCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestCoord.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestCoord json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestCoord.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestCoordCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestCoordCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestCoord.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestCoord unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestCoord.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestCoordCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestCoordCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestCoord.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestCoord gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestCoord.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.NestCoordCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.NestCoordCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestCoord.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestIdVal(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestIdVal mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestIdVal.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.NestIdValCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.NestIdValCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestIdVal.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestIdVal json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestIdVal.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestIdValCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestIdValCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestIdVal.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestIdVal unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestIdVal.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestIdValCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestIdValCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestIdVal.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestIdVal gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestIdVal.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.NestIdValCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.NestIdValCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestIdVal.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestLabel(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestLabel mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestLabel.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.NestLabelCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.NestLabelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestLabel.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestLabel json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestLabel.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestLabelCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestLabelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestLabel.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestLabel unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestLabel.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestLabelCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestLabelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestLabel.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestLabel gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestLabel.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.NestLabelCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.NestLabelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestLabel.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestMoney(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestMoney mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestMoney.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.NestMoneyCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.NestMoneyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestMoney.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestMoney json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestMoney.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestMoneyCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestMoneyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestMoney.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestMoney unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestMoney.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestMoneyCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestMoneyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestMoney.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestMoney gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestMoney.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.NestMoneyCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.NestMoneyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestMoney.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestRange32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestRange32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestRange32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.NestRange32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.NestRange32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestRange32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestRange32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestRange32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestRange32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestRange32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestRange32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestRange32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestRange32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestRange32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestRange32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestRange32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestRange32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestRange32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.NestRange32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.NestRange32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestRange32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestAddr(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestAddr mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestAddr.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.NestAddrCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.NestAddrCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestAddr.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestAddr json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestAddr.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestAddrCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestAddrCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestAddr.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestAddr unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestAddr.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestAddrCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestAddrCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestAddr.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestAddr gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestAddr.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.NestAddrCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.NestAddrCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestAddr.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestPoint3(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestPoint3 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestPoint3.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.NestPoint3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.NestPoint3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestPoint3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestPoint3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestPoint3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestPoint3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestPoint3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestPoint3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestPoint3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestPoint3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestPoint3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestPoint3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestPoint3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestPoint3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestPoint3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.NestPoint3Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.NestPoint3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestPoint3.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptNestInner(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptNestInner mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestInner.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptNestInnerCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptNestInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestInner.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestInner json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestInner.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptNestInnerCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptNestInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestInner.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestInner unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestInner.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptNestInnerCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptNestInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestInner.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestInner gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestInner.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptNestInnerCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptNestInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestInner.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptNestCoord(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptNestCoord mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestCoord.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptNestCoordCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptNestCoordCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestCoord.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestCoord json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestCoord.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptNestCoordCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptNestCoordCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestCoord.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestCoord unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestCoord.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptNestCoordCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptNestCoordCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestCoord.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestCoord gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestCoord.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptNestCoordCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptNestCoordCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestCoord.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptNestIdVal(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptNestIdVal mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestIdVal.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptNestIdValCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptNestIdValCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestIdVal.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestIdVal json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestIdVal.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptNestIdValCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptNestIdValCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestIdVal.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestIdVal unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestIdVal.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptNestIdValCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptNestIdValCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestIdVal.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestIdVal gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestIdVal.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptNestIdValCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptNestIdValCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestIdVal.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptNestLabel(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptNestLabel mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestLabel.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptNestLabelCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptNestLabelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestLabel.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestLabel json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestLabel.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptNestLabelCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptNestLabelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestLabel.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestLabel unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestLabel.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptNestLabelCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptNestLabelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestLabel.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestLabel gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestLabel.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptNestLabelCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptNestLabelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestLabel.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptNestMoney(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptNestMoney mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestMoney.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptNestMoneyCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptNestMoneyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestMoney.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestMoney json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestMoney.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptNestMoneyCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptNestMoneyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestMoney.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestMoney unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestMoney.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptNestMoneyCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptNestMoneyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestMoney.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestMoney gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestMoney.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptNestMoneyCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptNestMoneyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestMoney.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptNestRange32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptNestRange32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestRange32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptNestRange32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptNestRange32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestRange32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestRange32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestRange32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptNestRange32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptNestRange32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestRange32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestRange32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestRange32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptNestRange32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptNestRange32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestRange32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestRange32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestRange32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptNestRange32Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptNestRange32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestRange32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptNestAddr(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptNestAddr mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestAddr.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptNestAddrCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptNestAddrCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestAddr.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestAddr json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestAddr.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptNestAddrCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptNestAddrCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestAddr.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestAddr unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestAddr.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptNestAddrCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptNestAddrCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestAddr.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestAddr gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestAddr.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptNestAddrCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptNestAddrCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestAddr.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptNestPoint3(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptNestPoint3 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestPoint3.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptNestPoint3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptNestPoint3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestPoint3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestPoint3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestPoint3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptNestPoint3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptNestPoint3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestPoint3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestPoint3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestPoint3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptNestPoint3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptNestPoint3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestPoint3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestPoint3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestPoint3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptNestPoint3Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptNestPoint3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestPoint3.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ModelArr1(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ModelArr1 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr1.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.ModelArr1Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.ModelArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr1.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr1 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr1.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ModelArr1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ModelArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr1.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr1 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr1.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ModelArr1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ModelArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr1.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr1 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr1.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.ModelArr1Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.ModelArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr1.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ModelArr2(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ModelArr2 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr2.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.ModelArr2Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.ModelArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr2.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr2 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr2.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ModelArr2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ModelArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr2.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr2 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr2.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ModelArr2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ModelArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr2.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr2 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr2.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.ModelArr2Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.ModelArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr2.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ModelArr3(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ModelArr3 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr3.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.ModelArr3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.ModelArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ModelArr3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ModelArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ModelArr3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ModelArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.ModelArr3Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.ModelArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr3.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ModelArr4(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ModelArr4 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr4.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.ModelArr4Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.ModelArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr4.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr4 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr4.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ModelArr4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ModelArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr4.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr4 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr4.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ModelArr4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ModelArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr4.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr4 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr4.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.ModelArr4Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.ModelArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr4.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ModelArr5(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ModelArr5 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr5.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.ModelArr5Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.ModelArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr5.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr5 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr5.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ModelArr5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ModelArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr5.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr5 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr5.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ModelArr5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ModelArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr5.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr5 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr5.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.ModelArr5Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.ModelArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr5.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Mix01(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Mix01 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix01.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Mix01Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Mix01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix01.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix01 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix01.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix01Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix01.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix01 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix01.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix01Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix01.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix01 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix01.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Mix01Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Mix01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix01.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Mix02(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Mix02 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix02.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Mix02Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Mix02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix02.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix02 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix02.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix02Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix02.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix02 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix02.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix02Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix02.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix02 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix02.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Mix02Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Mix02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix02.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Mix03(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Mix03 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix03.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Mix03Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Mix03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix03.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix03 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix03.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix03Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix03.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix03 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix03.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix03Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix03.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix03 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix03.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Mix03Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Mix03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix03.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Mix04(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Mix04 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix04.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Mix04Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Mix04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix04.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix04 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix04.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix04Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix04.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix04 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix04.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix04Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix04.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix04 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix04.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Mix04Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Mix04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix04.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Mix05(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Mix05 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix05.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Mix05Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Mix05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix05.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix05 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix05.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix05Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix05.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix05 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix05.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix05Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix05.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix05 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix05.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Mix05Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Mix05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix05.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Mix06(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Mix06 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix06.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Mix06Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Mix06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix06.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix06 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix06.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix06Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix06.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix06 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix06.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix06Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix06.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix06 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix06.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Mix06Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Mix06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix06.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Mix07(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Mix07 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix07.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Mix07Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Mix07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix07.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix07 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix07.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix07Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix07.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix07 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix07.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix07Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix07.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix07 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix07.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Mix07Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Mix07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix07.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Mix08(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Mix08 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix08.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Mix08Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Mix08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix08.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix08 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix08.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix08Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix08.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix08 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix08.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix08Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix08.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix08 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix08.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Mix08Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Mix08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix08.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Mix09(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Mix09 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix09.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Mix09Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Mix09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix09.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix09 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix09.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix09Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix09.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix09 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix09.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix09Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix09.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix09 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix09.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Mix09Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Mix09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix09.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Mix10(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Mix10 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix10.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Mix10Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Mix10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix10.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix10 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix10.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix10Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix10.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix10 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix10.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix10Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix10.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix10 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix10.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Mix10Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Mix10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix10.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Mix11(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Mix11 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix11.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Mix11Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Mix11Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix11.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix11 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix11.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix11Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix11Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix11.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix11 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix11.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix11Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix11Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix11.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix11 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix11.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Mix11Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Mix11Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix11.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Mix12(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Mix12 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix12.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Mix12Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Mix12Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix12.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix12 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix12.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix12Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix12Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix12.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix12 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix12.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix12Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix12Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix12.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix12 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix12.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Mix12Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Mix12Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix12.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Mix13(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Mix13 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix13.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Mix13Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Mix13Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix13.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix13 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix13.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix13Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix13Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix13.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix13 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix13.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix13Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix13Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix13.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix13 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix13.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Mix13Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Mix13Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix13.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Mix14(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Mix14 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix14.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Mix14Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Mix14Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix14.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix14 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix14.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix14Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix14Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix14.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix14 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix14.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix14Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix14Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix14.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix14 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix14.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Mix14Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Mix14Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix14.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Mix15(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Mix15 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix15.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Mix15Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Mix15Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix15.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix15 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix15.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix15Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix15Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix15.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix15 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix15.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Mix15Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Mix15Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix15.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix15 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix15.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Mix15Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Mix15Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix15.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_AllOpt1(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("AllOpt1 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt1.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.AllOpt1Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.AllOpt1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt1.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt1 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt1.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.AllOpt1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.AllOpt1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt1.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt1 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt1.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.AllOpt1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.AllOpt1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt1.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt1 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt1.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.AllOpt1Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.AllOpt1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt1.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_AllOpt2(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("AllOpt2 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt2.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.AllOpt2Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.AllOpt2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt2.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt2 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt2.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.AllOpt2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.AllOpt2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt2.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt2 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt2.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.AllOpt2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.AllOpt2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt2.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt2 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt2.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.AllOpt2Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.AllOpt2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt2.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_AllOpt3(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("AllOpt3 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt3.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.AllOpt3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.AllOpt3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.AllOpt3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.AllOpt3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.AllOpt3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.AllOpt3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.AllOpt3Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.AllOpt3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt3.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_AllOpt4(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("AllOpt4 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt4.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.AllOpt4Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.AllOpt4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt4.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt4 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt4.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.AllOpt4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.AllOpt4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt4.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt4 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt4.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.AllOpt4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.AllOpt4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt4.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt4 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt4.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.AllOpt4Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.AllOpt4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt4.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_AllOpt5(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("AllOpt5 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt5.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.AllOpt5Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.AllOpt5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt5.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt5 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt5.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.AllOpt5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.AllOpt5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt5.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt5 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt5.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.AllOpt5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.AllOpt5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt5.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt5 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt5.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.AllOpt5Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.AllOpt5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt5.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_RecList(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("RecList mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecList.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.RecListCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.RecListCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecList.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecList json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecList.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.RecListCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.RecListCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecList.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecList unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecList.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.RecListCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.RecListCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecList.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecList gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecList.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.RecListCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.RecListCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecList.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_RecTree(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("RecTree mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecTree.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.RecTreeCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.RecTreeCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecTree.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecTree json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecTree.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.RecTreeCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.RecTreeCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecTree.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecTree unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecTree.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.RecTreeCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.RecTreeCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecTree.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecTree gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecTree.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.RecTreeCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.RecTreeCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecTree.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_RecChain(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("RecChain mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecChain.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.RecChainCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.RecChainCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecChain.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecChain json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecChain.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.RecChainCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.RecChainCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecChain.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecChain unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecChain.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.RecChainCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.RecChainCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecChain.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecChain gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecChain.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.RecChainCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.RecChainCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecChain.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_RecWrap(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("RecWrap mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecWrap.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.RecWrapCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.RecWrapCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecWrap.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecWrap json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecWrap.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.RecWrapCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.RecWrapCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecWrap.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecWrap unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecWrap.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.RecWrapCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.RecWrapCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecWrap.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecWrap gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecWrap.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.RecWrapCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.RecWrapCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecWrap.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_RecWide(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("RecWide mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecWide.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.RecWideCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.RecWideCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecWide.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecWide json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecWide.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.RecWideCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.RecWideCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecWide.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecWide unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecWide.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.RecWideCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.RecWideCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecWide.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecWide gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecWide.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.RecWideCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.RecWideCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecWide.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Wide20(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Wide20 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide20.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Wide20Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Wide20Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide20.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide20 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide20.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Wide20Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Wide20Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide20.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide20 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide20.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Wide20Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Wide20Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide20.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide20 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide20.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Wide20Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Wide20Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide20.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Wide25(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Wide25 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide25.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Wide25Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Wide25Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide25.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide25 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide25.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Wide25Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Wide25Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide25.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide25 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide25.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Wide25Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Wide25Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide25.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide25 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide25.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Wide25Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Wide25Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide25.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Wide30(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Wide30 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide30.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Wide30Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Wide30Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide30.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide30 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide30.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Wide30Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Wide30Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide30.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide30 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide30.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Wide30Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Wide30Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide30.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide30 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide30.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Wide30Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Wide30Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide30.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Wide35(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Wide35 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide35.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Wide35Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Wide35Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide35.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide35 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide35.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Wide35Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Wide35Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide35.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide35 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide35.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Wide35Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Wide35Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide35.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide35 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide35.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Wide35Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Wide35Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide35.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_Wide40(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Wide40 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide40.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.Wide40Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.Wide40Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide40.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide40 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide40.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Wide40Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Wide40Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide40.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide40 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide40.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.Wide40Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.Wide40Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide40.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide40 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide40.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.Wide40Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.Wide40Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide40.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeEmpty(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeEmpty mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeEmpty.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.EdgeEmptyCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.EdgeEmptyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeEmpty.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeEmpty json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeEmpty.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeEmptyCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeEmptyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeEmpty.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeEmpty unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeEmpty.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeEmptyCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeEmptyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeEmpty.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeEmpty gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeEmpty.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.EdgeEmptyCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.EdgeEmptyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeEmpty.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeOneOpt(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeOneOpt mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeOneOpt.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.EdgeOneOptCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.EdgeOneOptCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeOneOpt.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeOneOpt json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeOneOpt.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeOneOptCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeOneOptCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeOneOpt.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeOneOpt unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeOneOpt.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeOneOptCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeOneOptCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeOneOpt.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeOneOpt gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeOneOpt.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.EdgeOneOptCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.EdgeOneOptCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeOneOpt.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeBigNums(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeBigNums mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBigNums.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.EdgeBigNumsCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.EdgeBigNumsCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBigNums.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBigNums json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBigNums.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeBigNumsCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeBigNumsCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBigNums.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBigNums unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBigNums.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeBigNumsCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeBigNumsCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBigNums.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBigNums gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBigNums.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.EdgeBigNumsCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.EdgeBigNumsCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBigNums.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeZeroVals(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeZeroVals mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeZeroVals.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.EdgeZeroValsCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.EdgeZeroValsCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeZeroVals.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeZeroVals json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeZeroVals.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeZeroValsCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeZeroValsCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeZeroVals.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeZeroVals unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeZeroVals.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeZeroValsCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeZeroValsCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeZeroVals.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeZeroVals gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeZeroVals.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.EdgeZeroValsCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.EdgeZeroValsCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeZeroVals.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeNullable(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeNullable mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNullable.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.EdgeNullableCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.EdgeNullableCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNullable.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNullable json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNullable.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeNullableCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeNullableCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNullable.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNullable unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNullable.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeNullableCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeNullableCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNullable.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNullable gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNullable.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.EdgeNullableCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.EdgeNullableCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNullable.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeNegZero(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeNegZero mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNegZero.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.EdgeNegZeroCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.EdgeNegZeroCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNegZero.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNegZero json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNegZero.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeNegZeroCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeNegZeroCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNegZero.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNegZero unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNegZero.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeNegZeroCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeNegZeroCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNegZero.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNegZero gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNegZero.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.EdgeNegZeroCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.EdgeNegZeroCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNegZero.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeNullByte(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeNullByte mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNullByte.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.EdgeNullByteCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.EdgeNullByteCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNullByte.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNullByte json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNullByte.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeNullByteCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeNullByteCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNullByte.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNullByte unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNullByte.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeNullByteCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeNullByteCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNullByte.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNullByte gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNullByte.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.EdgeNullByteCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.EdgeNullByteCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNullByte.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeBoundary(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeBoundary mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBoundary.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.EdgeBoundaryCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.EdgeBoundaryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBoundary.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBoundary json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBoundary.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeBoundaryCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeBoundaryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBoundary.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBoundary unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBoundary.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeBoundaryCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeBoundaryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBoundary.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBoundary gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBoundary.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.EdgeBoundaryCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.EdgeBoundaryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBoundary.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeStrLen(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeStrLen mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeStrLen.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.EdgeStrLenCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.EdgeStrLenCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeStrLen.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeStrLen json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeStrLen.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeStrLenCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeStrLenCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeStrLen.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeStrLen unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeStrLen.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeStrLenCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeStrLenCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeStrLen.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeStrLen gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeStrLen.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.EdgeStrLenCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.EdgeStrLenCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeStrLen.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeBytesLen(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeBytesLen mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBytesLen.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.EdgeBytesLenCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.EdgeBytesLenCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBytesLen.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBytesLen json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBytesLen.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeBytesLenCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeBytesLenCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBytesLen.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBytesLen unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBytesLen.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeBytesLenCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeBytesLenCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBytesLen.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBytesLen gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBytesLen.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.EdgeBytesLenCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.EdgeBytesLenCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBytesLen.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeArrEmpty(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeArrEmpty mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeArrEmpty.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.EdgeArrEmptyCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.EdgeArrEmptyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeArrEmpty.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeArrEmpty json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeArrEmpty.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeArrEmptyCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeArrEmptyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeArrEmpty.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeArrEmpty unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeArrEmpty.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeArrEmptyCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeArrEmptyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeArrEmpty.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeArrEmpty gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeArrEmpty.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.EdgeArrEmptyCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.EdgeArrEmptyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeArrEmpty.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeArrBoundary(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeArrBoundary mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeArrBoundary.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.EdgeArrBoundaryCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.EdgeArrBoundaryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeArrBoundary.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeArrBoundary json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeArrBoundary.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeArrBoundaryCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeArrBoundaryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeArrBoundary.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeArrBoundary unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeArrBoundary.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.EdgeArrBoundaryCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.EdgeArrBoundaryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeArrBoundary.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeArrBoundary gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeArrBoundary.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.EdgeArrBoundaryCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.EdgeArrBoundaryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeArrBoundary.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptArr1(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptArr1 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr1.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptArr1Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr1.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr1 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr1.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptArr1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr1.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr1 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr1.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptArr1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr1.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr1 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr1.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptArr1Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr1.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptArr2(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptArr2 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr2.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptArr2Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr2.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr2 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr2.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptArr2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr2.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr2 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr2.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptArr2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr2.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr2 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr2.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptArr2Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr2.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptArr3(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptArr3 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr3.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptArr3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptArr3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptArr3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptArr3Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr3.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptArr4(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptArr4 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr4.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptArr4Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr4.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr4 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr4.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptArr4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr4.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr4 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr4.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptArr4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr4.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr4 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr4.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptArr4Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr4.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptArr5(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptArr5 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr5.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.OptArr5Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.OptArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr5.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr5 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr5.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptArr5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr5.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr5 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr5.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.OptArr5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.OptArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr5.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr5 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr5.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.OptArr5Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.OptArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr5.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestOpt1(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestOpt1 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt1.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.NestOpt1Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.NestOpt1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt1.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt1 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt1.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestOpt1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestOpt1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt1.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt1 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt1.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestOpt1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestOpt1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt1.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt1 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt1.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.NestOpt1Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.NestOpt1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt1.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestOpt2(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestOpt2 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt2.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.NestOpt2Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.NestOpt2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt2.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt2 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt2.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestOpt2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestOpt2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt2.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt2 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt2.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestOpt2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestOpt2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt2.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt2 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt2.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.NestOpt2Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.NestOpt2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt2.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestOpt3(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestOpt3 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt3.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.NestOpt3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.NestOpt3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestOpt3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestOpt3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestOpt3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestOpt3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.NestOpt3Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.NestOpt3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt3.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestOpt4(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestOpt4 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt4.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.NestOpt4Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.NestOpt4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt4.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt4 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt4.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestOpt4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestOpt4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt4.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt4 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt4.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestOpt4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestOpt4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt4.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt4 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt4.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.NestOpt4Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.NestOpt4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt4.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestOpt5(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestOpt5 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt5.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.NestOpt5Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.NestOpt5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt5.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt5 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt5.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestOpt5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestOpt5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt5.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt5 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt5.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestOpt5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestOpt5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt5.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt5 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt5.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.NestOpt5Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.NestOpt5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt5.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestOptInner1(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestOptInner1 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner1.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.NestOptInner1Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.NestOptInner1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner1.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner1 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner1.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestOptInner1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestOptInner1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner1.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner1 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner1.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestOptInner1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestOptInner1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner1.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner1 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner1.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.NestOptInner1Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.NestOptInner1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner1.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestOptInner2(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestOptInner2 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner2.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.NestOptInner2Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.NestOptInner2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner2.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner2 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner2.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestOptInner2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestOptInner2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner2.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner2 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner2.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestOptInner2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestOptInner2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner2.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner2 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner2.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.NestOptInner2Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.NestOptInner2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner2.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestOptInner3(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestOptInner3 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner3.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.NestOptInner3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.NestOptInner3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestOptInner3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestOptInner3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.NestOptInner3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.NestOptInner3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.NestOptInner3Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.NestOptInner3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner3.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DeepNest1(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DeepNest1 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest1.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.DeepNest1Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DeepNest1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest1.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest1 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest1.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DeepNest1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DeepNest1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest1.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest1 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest1.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DeepNest1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DeepNest1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest1.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest1 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest1.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DeepNest1Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DeepNest1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest1.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DeepNest2(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DeepNest2 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest2.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.DeepNest2Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DeepNest2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest2.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest2 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest2.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DeepNest2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DeepNest2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest2.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest2 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest2.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DeepNest2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DeepNest2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest2.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest2 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest2.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DeepNest2Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DeepNest2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest2.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DeepNest3(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DeepNest3 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest3.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.DeepNest3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DeepNest3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DeepNest3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DeepNest3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DeepNest3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DeepNest3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DeepNest3Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DeepNest3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest3.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DeepNest4(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DeepNest4 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest4.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.DeepNest4Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DeepNest4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest4.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest4 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest4.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DeepNest4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DeepNest4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest4.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest4 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest4.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DeepNest4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DeepNest4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest4.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest4 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest4.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DeepNest4Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DeepNest4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest4.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DeepNest5(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DeepNest5 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest5.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.DeepNest5Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DeepNest5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest5.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest5 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest5.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DeepNest5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DeepNest5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest5.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest5 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest5.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DeepNest5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DeepNest5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest5.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest5 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest5.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DeepNest5Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DeepNest5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest5.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DeepNest6(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DeepNest6 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest6.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.DeepNest6Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DeepNest6Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest6.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest6 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest6.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DeepNest6Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DeepNest6Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest6.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest6 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest6.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DeepNest6Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DeepNest6Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest6.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest6 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest6.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DeepNest6Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DeepNest6Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest6.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DeepNest7(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DeepNest7 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest7.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.DeepNest7Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.DeepNest7Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest7.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest7 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest7.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DeepNest7Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DeepNest7Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest7.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest7 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest7.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.DeepNest7Codec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.DeepNest7Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest7.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest7 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest7.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.DeepNest7Codec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.DeepNest7Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest7.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_TimestampEntry(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("TimestampEntry mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "TimestampEntry.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.TimestampEntryCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.TimestampEntryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "TimestampEntry.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("TimestampEntry json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "TimestampEntry.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.TimestampEntryCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.TimestampEntryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "TimestampEntry.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("TimestampEntry unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "TimestampEntry.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.TimestampEntryCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.TimestampEntryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "TimestampEntry.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("TimestampEntry gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "TimestampEntry.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.TimestampEntryCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.TimestampEntryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "TimestampEntry.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ConfigEntry(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ConfigEntry mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ConfigEntry.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types.ConfigEntryCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types.ConfigEntryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ConfigEntry.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ConfigEntry json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ConfigEntry.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ConfigEntryCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ConfigEntryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ConfigEntry.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ConfigEntry unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ConfigEntry.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types.ConfigEntryCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types.ConfigEntryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ConfigEntry.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ConfigEntry gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ConfigEntry.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types.ConfigEntryCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types.ConfigEntryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ConfigEntry.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestedSimple(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestedSimple mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestedSimple.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types_nested.NestedSimpleCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types_nested.NestedSimpleCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestedSimple.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestedSimple json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestedSimple.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types_nested.NestedSimpleCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types_nested.NestedSimpleCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestedSimple.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestedSimple unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestedSimple.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types_nested.NestedSimpleCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types_nested.NestedSimpleCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestedSimple.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestedSimple gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestedSimple.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types_nested.NestedSimpleCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types_nested.NestedSimpleCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestedSimple.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_DeepModel(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DeepModel mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepModel.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := specodec_all_types_nested_deep.DeepModelCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		specodec_all_types_nested_deep.DeepModelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepModel.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepModel json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepModel.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types_nested_deep.DeepModelCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types_nested_deep.DeepModelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepModel.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepModel unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepModel.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := specodec_all_types_nested_deep.DeepModelCodec.Decode(r)
		w := specodec.NewJsonWriter()
		specodec_all_types_nested_deep.DeepModelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepModel.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepModel gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepModel.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := specodec_all_types_nested_deep.DeepModelCodec.Decode(r)
		w := specodec.NewGronWriter()
		specodec_all_types_nested_deep.DeepModelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepModel.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypes(vecDir, outDir string) (passed, failed int) {
	var p, f int
	p, f = runModel_OptInner(vecDir, outDir); passed += p; failed += f
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
	p, f = runModel_OptSingleString(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptSingleBoolean(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptSingleInt8(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptSingleInt16(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptSingleInt32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptSingleInt64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptSingleUint8(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptSingleUint16(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptSingleUint32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptSingleUint64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptSingleFloat32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptSingleFloat64(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptSingleBytes(vecDir, outDir); passed += p; failed += f
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
	p, f = runModel_Five01(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Five02(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Five03(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Five04(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Five05(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Five06(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Five07(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Five08(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Five09(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Five10(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Ten01(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Ten02(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Ten03(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Ten04(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Ten05(vecDir, outDir); passed += p; failed += f
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
	p, f = runModel_NestInner(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestCoord(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestIdVal(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestLabel(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestMoney(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestRange32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestAddr(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestPoint3(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptNestInner(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptNestCoord(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptNestIdVal(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptNestLabel(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptNestMoney(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptNestRange32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptNestAddr(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptNestPoint3(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ModelArr1(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ModelArr2(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ModelArr3(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ModelArr4(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ModelArr5(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Mix01(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Mix02(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Mix03(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Mix04(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Mix05(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Mix06(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Mix07(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Mix08(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Mix09(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Mix10(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Mix11(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Mix12(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Mix13(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Mix14(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Mix15(vecDir, outDir); passed += p; failed += f
	p, f = runModel_AllOpt1(vecDir, outDir); passed += p; failed += f
	p, f = runModel_AllOpt2(vecDir, outDir); passed += p; failed += f
	p, f = runModel_AllOpt3(vecDir, outDir); passed += p; failed += f
	p, f = runModel_AllOpt4(vecDir, outDir); passed += p; failed += f
	p, f = runModel_AllOpt5(vecDir, outDir); passed += p; failed += f
	p, f = runModel_RecList(vecDir, outDir); passed += p; failed += f
	p, f = runModel_RecTree(vecDir, outDir); passed += p; failed += f
	p, f = runModel_RecChain(vecDir, outDir); passed += p; failed += f
	p, f = runModel_RecWrap(vecDir, outDir); passed += p; failed += f
	p, f = runModel_RecWide(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Wide20(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Wide25(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Wide30(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Wide35(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Wide40(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeEmpty(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeOneOpt(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeBigNums(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeZeroVals(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeNullable(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeNegZero(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeNullByte(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeBoundary(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeStrLen(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeBytesLen(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeArrEmpty(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeArrBoundary(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptArr1(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptArr2(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptArr3(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptArr4(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptArr5(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestOpt1(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestOpt2(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestOpt3(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestOpt4(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestOpt5(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestOptInner1(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestOptInner2(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestOptInner3(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DeepNest1(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DeepNest2(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DeepNest3(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DeepNest4(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DeepNest5(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DeepNest6(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DeepNest7(vecDir, outDir); passed += p; failed += f
	p, f = runModel_TimestampEntry(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ConfigEntry(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestedSimple(vecDir, outDir); passed += p; failed += f
	p, f = runModel_DeepModel(vecDir, outDir); passed += p; failed += f

	return
}
