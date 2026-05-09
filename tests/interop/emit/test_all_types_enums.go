package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	all_types_enums "emit_go/emit_gen/all_types_enums"
)


func runModel_EnumHolder(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EnumHolder mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EnumHolder.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_enums.EnumHolderCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_enums.EnumHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EnumHolder.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EnumHolder json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EnumHolder.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_enums.EnumHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_enums.EnumHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EnumHolder.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EnumHolder unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EnumHolder.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_enums.EnumHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_enums.EnumHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EnumHolder.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EnumHolder gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EnumHolder.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_enums.EnumHolderCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_enums.EnumHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EnumHolder.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptEnumHolder(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptEnumHolder mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptEnumHolder.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_enums.OptEnumHolderCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_enums.OptEnumHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptEnumHolder.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptEnumHolder json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptEnumHolder.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_enums.OptEnumHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_enums.OptEnumHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptEnumHolder.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptEnumHolder unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptEnumHolder.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_enums.OptEnumHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_enums.OptEnumHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptEnumHolder.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptEnumHolder gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptEnumHolder.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_enums.OptEnumHolderCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_enums.OptEnumHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptEnumHolder.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EnumArrayHolder(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EnumArrayHolder mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EnumArrayHolder.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_enums.EnumArrayHolderCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_enums.EnumArrayHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EnumArrayHolder.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EnumArrayHolder json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EnumArrayHolder.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_enums.EnumArrayHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_enums.EnumArrayHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EnumArrayHolder.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EnumArrayHolder unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EnumArrayHolder.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_enums.EnumArrayHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_enums.EnumArrayHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EnumArrayHolder.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EnumArrayHolder gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EnumArrayHolder.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_enums.EnumArrayHolderCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_enums.EnumArrayHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EnumArrayHolder.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EnumMixedHolder(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EnumMixedHolder mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EnumMixedHolder.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_enums.EnumMixedHolderCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_enums.EnumMixedHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EnumMixedHolder.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EnumMixedHolder json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EnumMixedHolder.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_enums.EnumMixedHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_enums.EnumMixedHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EnumMixedHolder.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EnumMixedHolder unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EnumMixedHolder.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_enums.EnumMixedHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_enums.EnumMixedHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EnumMixedHolder.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EnumMixedHolder gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EnumMixedHolder.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_enums.EnumMixedHolderCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_enums.EnumMixedHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EnumMixedHolder.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypesEnums(vecDir, outDir string) (passed, failed int) {
	var p, f int
	p, f = runModel_EnumHolder(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptEnumHolder(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EnumArrayHolder(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EnumMixedHolder(vecDir, outDir); passed += p; failed += f

	return
}
