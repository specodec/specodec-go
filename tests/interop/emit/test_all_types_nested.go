package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	all_types_nested "emit_go/emit_gen/all_types_nested"
)


func runModel_NestedSimple(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestedSimple mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestedSimple.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nested.NestedSimpleCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nested.NestedSimpleCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestedSimple.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestedSimple json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestedSimple.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nested.NestedSimpleCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nested.NestedSimpleCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestedSimple.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestedSimple unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestedSimple.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nested.NestedSimpleCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nested.NestedSimpleCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestedSimple.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestedSimple gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestedSimple.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nested.NestedSimpleCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nested.NestedSimpleCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestedSimple.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypesNested(vecDir, outDir string) (passed, failed int) {
	var p, f int
	p, f = runModel_NestedSimple(vecDir, outDir); passed += p; failed += f

	return
}
