package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	all_types_mixed "emit_go/emit_gen/all_types_mixed"
)


func runModel_ModelArr1(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ModelArr1 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr1.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_mixed.ModelArr1Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.ModelArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr1.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr1 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr1.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.ModelArr1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.ModelArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr1.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr1 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr1.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.ModelArr1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.ModelArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr1.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr1 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr1.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.ModelArr1Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.ModelArr1Codec.Encode(w, obj)
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
		obj := all_types_mixed.ModelArr2Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.ModelArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr2.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr2 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr2.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.ModelArr2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.ModelArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr2.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr2 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr2.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.ModelArr2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.ModelArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr2.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr2 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr2.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.ModelArr2Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.ModelArr2Codec.Encode(w, obj)
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
		obj := all_types_mixed.ModelArr3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.ModelArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.ModelArr3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.ModelArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.ModelArr3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.ModelArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.ModelArr3Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.ModelArr3Codec.Encode(w, obj)
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
		obj := all_types_mixed.ModelArr4Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.ModelArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr4.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr4 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr4.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.ModelArr4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.ModelArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr4.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr4 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr4.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.ModelArr4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.ModelArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr4.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr4 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr4.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.ModelArr4Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.ModelArr4Codec.Encode(w, obj)
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
		obj := all_types_mixed.ModelArr5Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.ModelArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr5.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr5 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr5.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.ModelArr5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.ModelArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr5.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr5 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr5.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.ModelArr5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.ModelArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ModelArr5.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ModelArr5 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ModelArr5.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.ModelArr5Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.ModelArr5Codec.Encode(w, obj)
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
		obj := all_types_mixed.Mix01Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.Mix01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix01.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix01 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix01.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix01Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix01.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix01 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix01.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix01Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix01Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix01.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix01 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix01.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.Mix01Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.Mix01Codec.Encode(w, obj)
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
		obj := all_types_mixed.Mix02Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.Mix02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix02.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix02 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix02.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix02Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix02.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix02 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix02.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix02Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix02Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix02.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix02 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix02.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.Mix02Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.Mix02Codec.Encode(w, obj)
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
		obj := all_types_mixed.Mix03Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.Mix03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix03.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix03 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix03.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix03Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix03.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix03 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix03.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix03Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix03Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix03.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix03 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix03.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.Mix03Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.Mix03Codec.Encode(w, obj)
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
		obj := all_types_mixed.Mix04Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.Mix04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix04.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix04 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix04.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix04Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix04.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix04 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix04.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix04Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix04Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix04.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix04 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix04.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.Mix04Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.Mix04Codec.Encode(w, obj)
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
		obj := all_types_mixed.Mix05Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.Mix05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix05.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix05 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix05.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix05Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix05.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix05 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix05.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix05Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix05Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix05.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix05 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix05.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.Mix05Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.Mix05Codec.Encode(w, obj)
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
		obj := all_types_mixed.Mix06Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.Mix06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix06.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix06 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix06.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix06Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix06.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix06 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix06.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix06Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix06Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix06.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix06 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix06.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.Mix06Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.Mix06Codec.Encode(w, obj)
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
		obj := all_types_mixed.Mix07Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.Mix07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix07.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix07 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix07.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix07Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix07.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix07 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix07.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix07Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix07Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix07.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix07 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix07.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.Mix07Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.Mix07Codec.Encode(w, obj)
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
		obj := all_types_mixed.Mix08Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.Mix08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix08.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix08 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix08.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix08Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix08.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix08 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix08.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix08Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix08Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix08.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix08 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix08.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.Mix08Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.Mix08Codec.Encode(w, obj)
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
		obj := all_types_mixed.Mix09Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.Mix09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix09.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix09 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix09.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix09Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix09.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix09 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix09.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix09Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix09Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix09.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix09 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix09.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.Mix09Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.Mix09Codec.Encode(w, obj)
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
		obj := all_types_mixed.Mix10Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.Mix10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix10.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix10 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix10.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix10Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix10.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix10 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix10.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix10Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix10Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix10.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix10 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix10.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.Mix10Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.Mix10Codec.Encode(w, obj)
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
		obj := all_types_mixed.Mix11Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.Mix11Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix11.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix11 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix11.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix11Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix11Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix11.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix11 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix11.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix11Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix11Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix11.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix11 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix11.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.Mix11Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.Mix11Codec.Encode(w, obj)
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
		obj := all_types_mixed.Mix12Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.Mix12Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix12.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix12 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix12.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix12Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix12Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix12.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix12 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix12.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix12Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix12Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix12.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix12 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix12.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.Mix12Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.Mix12Codec.Encode(w, obj)
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
		obj := all_types_mixed.Mix13Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.Mix13Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix13.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix13 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix13.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix13Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix13Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix13.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix13 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix13.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix13Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix13Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix13.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix13 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix13.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.Mix13Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.Mix13Codec.Encode(w, obj)
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
		obj := all_types_mixed.Mix14Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.Mix14Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix14.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix14 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix14.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix14Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix14Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix14.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix14 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix14.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix14Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix14Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix14.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix14 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix14.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.Mix14Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.Mix14Codec.Encode(w, obj)
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
		obj := all_types_mixed.Mix15Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.Mix15Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix15.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix15 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix15.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix15Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix15Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix15.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix15 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix15.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.Mix15Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.Mix15Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Mix15.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Mix15 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Mix15.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.Mix15Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.Mix15Codec.Encode(w, obj)
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
		obj := all_types_mixed.AllOpt1Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.AllOpt1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt1.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt1 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt1.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.AllOpt1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.AllOpt1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt1.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt1 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt1.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.AllOpt1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.AllOpt1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt1.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt1 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt1.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.AllOpt1Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.AllOpt1Codec.Encode(w, obj)
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
		obj := all_types_mixed.AllOpt2Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.AllOpt2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt2.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt2 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt2.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.AllOpt2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.AllOpt2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt2.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt2 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt2.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.AllOpt2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.AllOpt2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt2.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt2 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt2.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.AllOpt2Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.AllOpt2Codec.Encode(w, obj)
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
		obj := all_types_mixed.AllOpt3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.AllOpt3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.AllOpt3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.AllOpt3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.AllOpt3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.AllOpt3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.AllOpt3Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.AllOpt3Codec.Encode(w, obj)
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
		obj := all_types_mixed.AllOpt4Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.AllOpt4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt4.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt4 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt4.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.AllOpt4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.AllOpt4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt4.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt4 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt4.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.AllOpt4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.AllOpt4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt4.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt4 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt4.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.AllOpt4Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.AllOpt4Codec.Encode(w, obj)
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
		obj := all_types_mixed.AllOpt5Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_mixed.AllOpt5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt5.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt5 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt5.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.AllOpt5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.AllOpt5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt5.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt5 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt5.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_mixed.AllOpt5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_mixed.AllOpt5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt5.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("AllOpt5 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "AllOpt5.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_mixed.AllOpt5Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_mixed.AllOpt5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "AllOpt5.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypesMixed(vecDir, outDir string) (passed, failed int) {
	var p, f int
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

	return
}
