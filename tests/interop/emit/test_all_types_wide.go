package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	all_types_wide "emit_go/emit_gen/all_types_wide"
)


func runModel_Wide20(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Wide20 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide20.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_wide.Wide20Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_wide.Wide20Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide20.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide20 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide20.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_wide.Wide20Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_wide.Wide20Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide20.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide20 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide20.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_wide.Wide20Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_wide.Wide20Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide20.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide20 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide20.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_wide.Wide20Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_wide.Wide20Codec.Encode(w, obj)
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
		obj := all_types_wide.Wide25Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_wide.Wide25Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide25.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide25 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide25.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_wide.Wide25Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_wide.Wide25Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide25.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide25 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide25.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_wide.Wide25Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_wide.Wide25Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide25.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide25 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide25.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_wide.Wide25Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_wide.Wide25Codec.Encode(w, obj)
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
		obj := all_types_wide.Wide30Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_wide.Wide30Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide30.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide30 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide30.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_wide.Wide30Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_wide.Wide30Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide30.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide30 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide30.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_wide.Wide30Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_wide.Wide30Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide30.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide30 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide30.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_wide.Wide30Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_wide.Wide30Codec.Encode(w, obj)
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
		obj := all_types_wide.Wide35Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_wide.Wide35Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide35.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide35 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide35.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_wide.Wide35Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_wide.Wide35Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide35.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide35 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide35.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_wide.Wide35Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_wide.Wide35Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide35.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide35 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide35.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_wide.Wide35Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_wide.Wide35Codec.Encode(w, obj)
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
		obj := all_types_wide.Wide40Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_wide.Wide40Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide40.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide40 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide40.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_wide.Wide40Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_wide.Wide40Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide40.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide40 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide40.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_wide.Wide40Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_wide.Wide40Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide40.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Wide40 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Wide40.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_wide.Wide40Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_wide.Wide40Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Wide40.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypesWide(vecDir, outDir string) (passed, failed int) {
	var p, f int
	p, f = runModel_Wide20(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Wide25(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Wide30(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Wide35(vecDir, outDir); passed += p; failed += f
	p, f = runModel_Wide40(vecDir, outDir); passed += p; failed += f

	return
}
