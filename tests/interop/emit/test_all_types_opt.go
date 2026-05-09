package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	all_types_opt "emit_go/emit_gen/all_types_opt"
)


func runModel_OptSingleString(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptSingleString mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleString.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_opt.OptSingleStringCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_opt.OptSingleStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleString.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleString json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleString.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleStringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleString.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleString unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleString.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleStringCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleStringCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleString.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleString gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleString.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_opt.OptSingleStringCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_opt.OptSingleStringCodec.Encode(w, obj)
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
		obj := all_types_opt.OptSingleBooleanCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_opt.OptSingleBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleBoolean.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleBoolean json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleBoolean.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleBoolean.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleBoolean unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleBoolean.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleBooleanCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleBooleanCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleBoolean.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleBoolean gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleBoolean.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_opt.OptSingleBooleanCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_opt.OptSingleBooleanCodec.Encode(w, obj)
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
		obj := all_types_opt.OptSingleInt8Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_opt.OptSingleInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt8.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt8 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt8.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleInt8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt8.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt8 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt8.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleInt8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleInt8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt8.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt8 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt8.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_opt.OptSingleInt8Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_opt.OptSingleInt8Codec.Encode(w, obj)
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
		obj := all_types_opt.OptSingleInt16Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_opt.OptSingleInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt16.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt16 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt16.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleInt16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt16.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt16 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt16.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleInt16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleInt16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt16.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt16 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt16.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_opt.OptSingleInt16Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_opt.OptSingleInt16Codec.Encode(w, obj)
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
		obj := all_types_opt.OptSingleInt32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_opt.OptSingleInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleInt32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleInt32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_opt.OptSingleInt32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_opt.OptSingleInt32Codec.Encode(w, obj)
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
		obj := all_types_opt.OptSingleInt64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_opt.OptSingleInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleInt64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleInt64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleInt64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleInt64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleInt64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_opt.OptSingleInt64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_opt.OptSingleInt64Codec.Encode(w, obj)
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
		obj := all_types_opt.OptSingleUint8Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_opt.OptSingleUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint8.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint8 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint8.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleUint8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint8.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint8 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint8.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleUint8Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleUint8Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint8.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint8 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint8.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_opt.OptSingleUint8Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_opt.OptSingleUint8Codec.Encode(w, obj)
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
		obj := all_types_opt.OptSingleUint16Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_opt.OptSingleUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint16.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint16 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint16.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleUint16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint16.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint16 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint16.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleUint16Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleUint16Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint16.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint16 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint16.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_opt.OptSingleUint16Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_opt.OptSingleUint16Codec.Encode(w, obj)
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
		obj := all_types_opt.OptSingleUint32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_opt.OptSingleUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleUint32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleUint32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleUint32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_opt.OptSingleUint32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_opt.OptSingleUint32Codec.Encode(w, obj)
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
		obj := all_types_opt.OptSingleUint64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_opt.OptSingleUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleUint64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleUint64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleUint64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleUint64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleUint64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_opt.OptSingleUint64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_opt.OptSingleUint64Codec.Encode(w, obj)
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
		obj := all_types_opt.OptSingleFloat32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_opt.OptSingleFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleFloat32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleFloat32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleFloat32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleFloat32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleFloat32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleFloat32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleFloat32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleFloat32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleFloat32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleFloat32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleFloat32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleFloat32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_opt.OptSingleFloat32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_opt.OptSingleFloat32Codec.Encode(w, obj)
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
		obj := all_types_opt.OptSingleFloat64Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_opt.OptSingleFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleFloat64.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleFloat64 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleFloat64.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleFloat64.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleFloat64 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleFloat64.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleFloat64Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleFloat64Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleFloat64.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleFloat64 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleFloat64.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_opt.OptSingleFloat64Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_opt.OptSingleFloat64Codec.Encode(w, obj)
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
		obj := all_types_opt.OptSingleBytesCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_opt.OptSingleBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleBytes.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleBytes json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleBytes.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleBytes.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleBytes unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleBytes.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_opt.OptSingleBytesCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_opt.OptSingleBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleBytes.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptSingleBytes gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptSingleBytes.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_opt.OptSingleBytesCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_opt.OptSingleBytesCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptSingleBytes.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypesOpt(vecDir, outDir string) (passed, failed int) {
	var p, f int
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

	return
}
