package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	all_types_recursive "emit_go/emit_gen/all_types_recursive"
)


func runModel_RecList(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("RecList mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecList.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_recursive.RecListCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_recursive.RecListCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecList.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecList json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecList.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_recursive.RecListCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_recursive.RecListCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecList.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecList unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecList.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_recursive.RecListCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_recursive.RecListCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecList.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecList gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecList.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_recursive.RecListCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_recursive.RecListCodec.Encode(w, obj)
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
		obj := all_types_recursive.RecTreeCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_recursive.RecTreeCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecTree.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecTree json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecTree.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_recursive.RecTreeCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_recursive.RecTreeCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecTree.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecTree unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecTree.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_recursive.RecTreeCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_recursive.RecTreeCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecTree.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecTree gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecTree.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_recursive.RecTreeCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_recursive.RecTreeCodec.Encode(w, obj)
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
		obj := all_types_recursive.RecChainCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_recursive.RecChainCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecChain.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecChain json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecChain.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_recursive.RecChainCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_recursive.RecChainCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecChain.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecChain unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecChain.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_recursive.RecChainCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_recursive.RecChainCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecChain.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecChain gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecChain.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_recursive.RecChainCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_recursive.RecChainCodec.Encode(w, obj)
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
		obj := all_types_recursive.RecWrapCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_recursive.RecWrapCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecWrap.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecWrap json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecWrap.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_recursive.RecWrapCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_recursive.RecWrapCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecWrap.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecWrap unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecWrap.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_recursive.RecWrapCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_recursive.RecWrapCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecWrap.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecWrap gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecWrap.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_recursive.RecWrapCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_recursive.RecWrapCodec.Encode(w, obj)
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
		obj := all_types_recursive.RecWideCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_recursive.RecWideCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecWide.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecWide json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecWide.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_recursive.RecWideCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_recursive.RecWideCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecWide.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecWide unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecWide.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_recursive.RecWideCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_recursive.RecWideCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecWide.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("RecWide gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "RecWide.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_recursive.RecWideCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_recursive.RecWideCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "RecWide.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypesRecursive(vecDir, outDir string) (passed, failed int) {
	var p, f int
	p, f = runModel_RecList(vecDir, outDir); passed += p; failed += f
	p, f = runModel_RecTree(vecDir, outDir); passed += p; failed += f
	p, f = runModel_RecChain(vecDir, outDir); passed += p; failed += f
	p, f = runModel_RecWrap(vecDir, outDir); passed += p; failed += f
	p, f = runModel_RecWide(vecDir, outDir); passed += p; failed += f

	return
}
