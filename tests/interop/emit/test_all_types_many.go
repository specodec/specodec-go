package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	all_types_many "emit_go/emit_gen/all_types_many"
)


func runModel_Five01(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Five01 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five01.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_many.Five01Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_many.Five01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five01.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five01 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five01.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five01Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five01.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five01 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five01.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five01Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five01.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five01 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five01.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_many.Five01Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_many.Five01Codec.Encode(w, obj)
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
		obj := all_types_many.Five02Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_many.Five02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five02.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five02 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five02.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five02Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five02.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five02 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five02.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five02Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five02.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five02 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five02.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_many.Five02Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_many.Five02Codec.Encode(w, obj)
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
		obj := all_types_many.Five03Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_many.Five03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five03.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five03 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five03.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five03Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five03.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five03 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five03.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five03Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five03.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five03 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five03.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_many.Five03Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_many.Five03Codec.Encode(w, obj)
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
		obj := all_types_many.Five04Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_many.Five04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five04.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five04 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five04.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five04Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five04.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five04 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five04.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five04Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five04.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five04 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five04.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_many.Five04Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_many.Five04Codec.Encode(w, obj)
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
		obj := all_types_many.Five05Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_many.Five05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five05.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five05 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five05.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five05Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five05.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five05 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five05.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five05Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five05.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five05 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five05.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_many.Five05Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_many.Five05Codec.Encode(w, obj)
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
		obj := all_types_many.Five06Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_many.Five06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five06.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five06 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five06.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five06Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five06.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five06 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five06.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five06Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five06.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five06 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five06.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_many.Five06Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_many.Five06Codec.Encode(w, obj)
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
		obj := all_types_many.Five07Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_many.Five07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five07.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five07 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five07.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five07Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five07.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five07 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five07.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five07Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five07.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five07 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five07.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_many.Five07Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_many.Five07Codec.Encode(w, obj)
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
		obj := all_types_many.Five08Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_many.Five08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five08.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five08 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five08.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five08Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five08.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five08 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five08.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five08Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five08.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five08 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five08.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_many.Five08Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_many.Five08Codec.Encode(w, obj)
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
		obj := all_types_many.Five09Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_many.Five09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five09.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five09 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five09.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five09Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five09.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five09 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five09.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five09Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five09.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five09 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five09.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_many.Five09Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_many.Five09Codec.Encode(w, obj)
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
		obj := all_types_many.Five10Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_many.Five10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five10.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five10 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five10.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five10Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five10.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five10 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five10.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Five10Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Five10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Five10.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Five10 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Five10.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_many.Five10Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_many.Five10Codec.Encode(w, obj)
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
		obj := all_types_many.Ten01Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_many.Ten01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten01.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten01 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten01.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Ten01Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Ten01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten01.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten01 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten01.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Ten01Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Ten01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten01.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten01 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten01.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_many.Ten01Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_many.Ten01Codec.Encode(w, obj)
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
		obj := all_types_many.Ten02Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_many.Ten02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten02.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten02 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten02.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Ten02Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Ten02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten02.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten02 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten02.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Ten02Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Ten02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten02.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten02 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten02.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_many.Ten02Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_many.Ten02Codec.Encode(w, obj)
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
		obj := all_types_many.Ten03Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_many.Ten03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten03.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten03 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten03.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Ten03Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Ten03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten03.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten03 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten03.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Ten03Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Ten03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten03.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten03 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten03.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_many.Ten03Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_many.Ten03Codec.Encode(w, obj)
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
		obj := all_types_many.Ten04Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_many.Ten04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten04.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten04 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten04.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Ten04Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Ten04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten04.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten04 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten04.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Ten04Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Ten04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten04.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten04 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten04.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_many.Ten04Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_many.Ten04Codec.Encode(w, obj)
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
		obj := all_types_many.Ten05Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_many.Ten05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten05.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten05 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten05.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Ten05Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Ten05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten05.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten05 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten05.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_many.Ten05Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_many.Ten05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten05.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ten05 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ten05.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_many.Ten05Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_many.Ten05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ten05.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypesMany(vecDir, outDir string) (passed, failed int) {
	var p, f int
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

	return
}
