package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	all_types_nested_deep "emit_go/emit_gen/all_types_nested_deep"
)


func runModel_DeepModel(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("DeepModel mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepModel.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nested_deep.DeepModelCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nested_deep.DeepModelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepModel.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepModel json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepModel.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nested_deep.DeepModelCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nested_deep.DeepModelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepModel.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepModel unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepModel.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nested_deep.DeepModelCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nested_deep.DeepModelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepModel.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("DeepModel gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "DeepModel.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nested_deep.DeepModelCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nested_deep.DeepModelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "DeepModel.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypesNestedDeep(vecDir, outDir string) (passed, failed int) {
	var p, f int
	p, f = runModel_DeepModel(vecDir, outDir); passed += p; failed += f

	return
}
