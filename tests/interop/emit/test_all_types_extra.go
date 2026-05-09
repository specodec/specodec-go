package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	all_types_extra "emit_go/emit_gen/all_types_extra"
)


func runModel_OptArr1(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptArr1 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr1.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_extra.OptArr1Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.OptArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr1.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr1 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr1.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.OptArr1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.OptArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr1.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr1 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr1.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.OptArr1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.OptArr1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr1.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr1 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr1.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.OptArr1Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.OptArr1Codec.Encode(w, obj)
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
		obj := all_types_extra.OptArr2Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.OptArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr2.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr2 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr2.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.OptArr2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.OptArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr2.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr2 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr2.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.OptArr2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.OptArr2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr2.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr2 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr2.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.OptArr2Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.OptArr2Codec.Encode(w, obj)
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
		obj := all_types_extra.OptArr3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.OptArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.OptArr3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.OptArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.OptArr3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.OptArr3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.OptArr3Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.OptArr3Codec.Encode(w, obj)
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
		obj := all_types_extra.OptArr4Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.OptArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr4.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr4 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr4.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.OptArr4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.OptArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr4.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr4 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr4.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.OptArr4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.OptArr4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr4.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr4 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr4.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.OptArr4Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.OptArr4Codec.Encode(w, obj)
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
		obj := all_types_extra.OptArr5Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.OptArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr5.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr5 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr5.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.OptArr5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.OptArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr5.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr5 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr5.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.OptArr5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.OptArr5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptArr5.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptArr5 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptArr5.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.OptArr5Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.OptArr5Codec.Encode(w, obj)
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
		obj := all_types_extra.NestOpt1Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.NestOpt1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt1.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt1 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt1.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.NestOpt1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.NestOpt1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt1.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt1 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt1.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.NestOpt1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.NestOpt1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt1.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt1 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt1.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.NestOpt1Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.NestOpt1Codec.Encode(w, obj)
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
		obj := all_types_extra.NestOpt2Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.NestOpt2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt2.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt2 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt2.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.NestOpt2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.NestOpt2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt2.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt2 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt2.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.NestOpt2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.NestOpt2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt2.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt2 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt2.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.NestOpt2Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.NestOpt2Codec.Encode(w, obj)
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
		obj := all_types_extra.NestOpt3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.NestOpt3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.NestOpt3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.NestOpt3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.NestOpt3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.NestOpt3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.NestOpt3Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.NestOpt3Codec.Encode(w, obj)
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
		obj := all_types_extra.NestOpt4Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.NestOpt4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt4.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt4 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt4.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.NestOpt4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.NestOpt4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt4.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt4 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt4.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.NestOpt4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.NestOpt4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt4.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt4 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt4.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.NestOpt4Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.NestOpt4Codec.Encode(w, obj)
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
		obj := all_types_extra.NestOpt5Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.NestOpt5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt5.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt5 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt5.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.NestOpt5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.NestOpt5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt5.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt5 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt5.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.NestOpt5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.NestOpt5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOpt5.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOpt5 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOpt5.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.NestOpt5Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.NestOpt5Codec.Encode(w, obj)
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
		obj := all_types_extra.NestOptInner1Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.NestOptInner1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner1.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner1 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner1.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.NestOptInner1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.NestOptInner1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner1.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner1 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner1.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.NestOptInner1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.NestOptInner1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner1.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner1 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner1.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.NestOptInner1Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.NestOptInner1Codec.Encode(w, obj)
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
		obj := all_types_extra.NestOptInner2Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.NestOptInner2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner2.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner2 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner2.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.NestOptInner2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.NestOptInner2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner2.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner2 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner2.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.NestOptInner2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.NestOptInner2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner2.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner2 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner2.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.NestOptInner2Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.NestOptInner2Codec.Encode(w, obj)
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
		obj := all_types_extra.NestOptInner3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.NestOptInner3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.NestOptInner3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.NestOptInner3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.NestOptInner3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.NestOptInner3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestOptInner3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestOptInner3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestOptInner3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.NestOptInner3Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.NestOptInner3Codec.Encode(w, obj)
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
		obj := all_types_extra.DeepNest1Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.DeepNest1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest1.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest1 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest1.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.DeepNest1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.DeepNest1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest1.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest1 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest1.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.DeepNest1Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.DeepNest1Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest1.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest1 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest1.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.DeepNest1Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.DeepNest1Codec.Encode(w, obj)
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
		obj := all_types_extra.DeepNest2Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.DeepNest2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest2.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest2 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest2.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.DeepNest2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.DeepNest2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest2.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest2 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest2.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.DeepNest2Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.DeepNest2Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest2.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest2 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest2.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.DeepNest2Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.DeepNest2Codec.Encode(w, obj)
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
		obj := all_types_extra.DeepNest3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.DeepNest3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.DeepNest3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.DeepNest3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.DeepNest3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.DeepNest3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.DeepNest3Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.DeepNest3Codec.Encode(w, obj)
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
		obj := all_types_extra.DeepNest4Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.DeepNest4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest4.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest4 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest4.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.DeepNest4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.DeepNest4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest4.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest4 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest4.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.DeepNest4Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.DeepNest4Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest4.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest4 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest4.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.DeepNest4Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.DeepNest4Codec.Encode(w, obj)
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
		obj := all_types_extra.DeepNest5Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.DeepNest5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest5.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest5 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest5.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.DeepNest5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.DeepNest5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest5.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest5 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest5.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.DeepNest5Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.DeepNest5Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest5.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest5 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest5.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.DeepNest5Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.DeepNest5Codec.Encode(w, obj)
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
		obj := all_types_extra.DeepNest6Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.DeepNest6Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest6.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest6 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest6.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.DeepNest6Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.DeepNest6Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest6.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest6 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest6.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.DeepNest6Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.DeepNest6Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest6.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest6 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest6.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.DeepNest6Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.DeepNest6Codec.Encode(w, obj)
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
		obj := all_types_extra.DeepNest7Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.DeepNest7Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest7.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest7 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest7.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.DeepNest7Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.DeepNest7Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest7.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest7 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest7.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.DeepNest7Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.DeepNest7Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepNest7.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepNest7 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepNest7.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.DeepNest7Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.DeepNest7Codec.Encode(w, obj)
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
		obj := all_types_extra.TimestampEntryCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.TimestampEntryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "TimestampEntry.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("TimestampEntry json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "TimestampEntry.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.TimestampEntryCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.TimestampEntryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "TimestampEntry.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("TimestampEntry unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "TimestampEntry.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.TimestampEntryCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.TimestampEntryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "TimestampEntry.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("TimestampEntry gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "TimestampEntry.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.TimestampEntryCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.TimestampEntryCodec.Encode(w, obj)
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
		obj := all_types_extra.ConfigEntryCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_extra.ConfigEntryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ConfigEntry.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ConfigEntry json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ConfigEntry.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.ConfigEntryCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.ConfigEntryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ConfigEntry.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ConfigEntry unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ConfigEntry.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_extra.ConfigEntryCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_extra.ConfigEntryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ConfigEntry.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ConfigEntry gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ConfigEntry.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_extra.ConfigEntryCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_extra.ConfigEntryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ConfigEntry.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypesExtra(vecDir, outDir string) (passed, failed int) {
	var p, f int
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

	return
}
