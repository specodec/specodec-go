package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	all_types_unions "emit_go/emit_gen/all_types_unions"
)


func runModel_UnionFieldHolder(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("UnionFieldHolder mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "UnionFieldHolder.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.UnionFieldHolderCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.UnionFieldHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "UnionFieldHolder.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("UnionFieldHolder json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "UnionFieldHolder.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.UnionFieldHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.UnionFieldHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "UnionFieldHolder.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("UnionFieldHolder unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "UnionFieldHolder.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.UnionFieldHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.UnionFieldHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "UnionFieldHolder.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("UnionFieldHolder gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "UnionFieldHolder.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.UnionFieldHolderCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.UnionFieldHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "UnionFieldHolder.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptUnionFieldHolder(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptUnionFieldHolder mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptUnionFieldHolder.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.OptUnionFieldHolderCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.OptUnionFieldHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptUnionFieldHolder.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptUnionFieldHolder json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptUnionFieldHolder.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.OptUnionFieldHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.OptUnionFieldHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptUnionFieldHolder.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptUnionFieldHolder unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptUnionFieldHolder.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.OptUnionFieldHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.OptUnionFieldHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptUnionFieldHolder.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptUnionFieldHolder gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptUnionFieldHolder.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.OptUnionFieldHolderCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.OptUnionFieldHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptUnionFieldHolder.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_UnionArrayHolder(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("UnionArrayHolder mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "UnionArrayHolder.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.UnionArrayHolderCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.UnionArrayHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "UnionArrayHolder.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("UnionArrayHolder json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "UnionArrayHolder.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.UnionArrayHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.UnionArrayHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "UnionArrayHolder.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("UnionArrayHolder unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "UnionArrayHolder.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.UnionArrayHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.UnionArrayHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "UnionArrayHolder.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("UnionArrayHolder gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "UnionArrayHolder.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.UnionArrayHolderCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.UnionArrayHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "UnionArrayHolder.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_UnionMixedHolder(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("UnionMixedHolder mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "UnionMixedHolder.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.UnionMixedHolderCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.UnionMixedHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "UnionMixedHolder.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("UnionMixedHolder json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "UnionMixedHolder.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.UnionMixedHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.UnionMixedHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "UnionMixedHolder.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("UnionMixedHolder unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "UnionMixedHolder.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.UnionMixedHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.UnionMixedHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "UnionMixedHolder.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("UnionMixedHolder gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "UnionMixedHolder.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.UnionMixedHolderCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.UnionMixedHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "UnionMixedHolder.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_UnionScalarHolder(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("UnionScalarHolder mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "UnionScalarHolder.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.UnionScalarHolderCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.UnionScalarHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "UnionScalarHolder.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("UnionScalarHolder json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "UnionScalarHolder.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.UnionScalarHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.UnionScalarHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "UnionScalarHolder.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("UnionScalarHolder unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "UnionScalarHolder.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.UnionScalarHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.UnionScalarHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "UnionScalarHolder.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("UnionScalarHolder gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "UnionScalarHolder.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.UnionScalarHolderCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.UnionScalarHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "UnionScalarHolder.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypesUnions(vecDir, outDir string) (passed, failed int) {
	var p, f int
	p, f = runModel_UnionFieldHolder(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptUnionFieldHolder(vecDir, outDir); passed += p; failed += f
	p, f = runModel_UnionArrayHolder(vecDir, outDir); passed += p; failed += f
	p, f = runModel_UnionMixedHolder(vecDir, outDir); passed += p; failed += f
	p, f = runModel_UnionScalarHolder(vecDir, outDir); passed += p; failed += f

	return
}
