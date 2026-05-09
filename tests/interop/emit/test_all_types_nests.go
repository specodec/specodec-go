package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	all_types_nests "emit_go/emit_gen/all_types_nests"
)


func runModel_NestInner(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestInner mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestInner.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nests.NestInnerCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nests.NestInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestInner.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestInner json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestInner.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.NestInnerCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.NestInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestInner.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestInner unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestInner.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.NestInnerCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.NestInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestInner.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestInner gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestInner.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nests.NestInnerCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nests.NestInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestInner.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestCoord(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestCoord mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestCoord.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nests.NestCoordCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nests.NestCoordCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestCoord.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestCoord json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestCoord.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.NestCoordCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.NestCoordCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestCoord.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestCoord unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestCoord.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.NestCoordCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.NestCoordCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestCoord.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestCoord gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestCoord.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nests.NestCoordCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nests.NestCoordCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestCoord.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestIdVal(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestIdVal mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestIdVal.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nests.NestIdValCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nests.NestIdValCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestIdVal.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestIdVal json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestIdVal.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.NestIdValCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.NestIdValCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestIdVal.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestIdVal unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestIdVal.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.NestIdValCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.NestIdValCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestIdVal.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestIdVal gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestIdVal.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nests.NestIdValCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nests.NestIdValCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestIdVal.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestLabel(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestLabel mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestLabel.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nests.NestLabelCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nests.NestLabelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestLabel.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestLabel json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestLabel.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.NestLabelCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.NestLabelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestLabel.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestLabel unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestLabel.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.NestLabelCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.NestLabelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestLabel.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestLabel gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestLabel.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nests.NestLabelCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nests.NestLabelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestLabel.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestMoney(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestMoney mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestMoney.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nests.NestMoneyCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nests.NestMoneyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestMoney.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestMoney json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestMoney.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.NestMoneyCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.NestMoneyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestMoney.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestMoney unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestMoney.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.NestMoneyCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.NestMoneyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestMoney.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestMoney gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestMoney.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nests.NestMoneyCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nests.NestMoneyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestMoney.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestRange32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestRange32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestRange32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nests.NestRange32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nests.NestRange32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestRange32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestRange32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestRange32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.NestRange32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.NestRange32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestRange32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestRange32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestRange32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.NestRange32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.NestRange32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestRange32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestRange32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestRange32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nests.NestRange32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nests.NestRange32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestRange32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestAddr(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestAddr mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestAddr.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nests.NestAddrCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nests.NestAddrCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestAddr.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestAddr json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestAddr.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.NestAddrCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.NestAddrCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestAddr.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestAddr unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestAddr.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.NestAddrCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.NestAddrCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestAddr.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestAddr gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestAddr.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nests.NestAddrCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nests.NestAddrCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestAddr.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestPoint3(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestPoint3 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestPoint3.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nests.NestPoint3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nests.NestPoint3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestPoint3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestPoint3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestPoint3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.NestPoint3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.NestPoint3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestPoint3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestPoint3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestPoint3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.NestPoint3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.NestPoint3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestPoint3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestPoint3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestPoint3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nests.NestPoint3Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nests.NestPoint3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestPoint3.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptNestInner(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptNestInner mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestInner.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nests.OptNestInnerCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nests.OptNestInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestInner.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestInner json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestInner.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.OptNestInnerCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.OptNestInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestInner.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestInner unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestInner.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.OptNestInnerCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.OptNestInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestInner.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestInner gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestInner.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nests.OptNestInnerCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nests.OptNestInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestInner.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptNestCoord(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptNestCoord mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestCoord.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nests.OptNestCoordCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nests.OptNestCoordCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestCoord.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestCoord json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestCoord.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.OptNestCoordCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.OptNestCoordCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestCoord.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestCoord unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestCoord.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.OptNestCoordCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.OptNestCoordCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestCoord.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestCoord gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestCoord.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nests.OptNestCoordCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nests.OptNestCoordCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestCoord.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptNestIdVal(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptNestIdVal mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestIdVal.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nests.OptNestIdValCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nests.OptNestIdValCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestIdVal.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestIdVal json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestIdVal.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.OptNestIdValCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.OptNestIdValCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestIdVal.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestIdVal unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestIdVal.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.OptNestIdValCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.OptNestIdValCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestIdVal.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestIdVal gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestIdVal.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nests.OptNestIdValCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nests.OptNestIdValCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestIdVal.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptNestLabel(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptNestLabel mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestLabel.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nests.OptNestLabelCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nests.OptNestLabelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestLabel.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestLabel json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestLabel.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.OptNestLabelCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.OptNestLabelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestLabel.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestLabel unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestLabel.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.OptNestLabelCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.OptNestLabelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestLabel.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestLabel gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestLabel.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nests.OptNestLabelCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nests.OptNestLabelCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestLabel.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptNestMoney(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptNestMoney mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestMoney.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nests.OptNestMoneyCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nests.OptNestMoneyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestMoney.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestMoney json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestMoney.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.OptNestMoneyCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.OptNestMoneyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestMoney.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestMoney unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestMoney.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.OptNestMoneyCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.OptNestMoneyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestMoney.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestMoney gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestMoney.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nests.OptNestMoneyCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nests.OptNestMoneyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestMoney.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptNestRange32(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptNestRange32 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestRange32.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nests.OptNestRange32Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nests.OptNestRange32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestRange32.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestRange32 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestRange32.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.OptNestRange32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.OptNestRange32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestRange32.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestRange32 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestRange32.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.OptNestRange32Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.OptNestRange32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestRange32.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestRange32 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestRange32.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nests.OptNestRange32Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nests.OptNestRange32Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestRange32.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptNestAddr(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptNestAddr mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestAddr.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nests.OptNestAddrCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nests.OptNestAddrCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestAddr.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestAddr json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestAddr.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.OptNestAddrCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.OptNestAddrCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestAddr.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestAddr unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestAddr.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.OptNestAddrCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.OptNestAddrCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestAddr.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestAddr gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestAddr.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nests.OptNestAddrCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nests.OptNestAddrCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestAddr.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptNestPoint3(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptNestPoint3 mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestPoint3.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_nests.OptNestPoint3Codec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_nests.OptNestPoint3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestPoint3.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestPoint3 json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestPoint3.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.OptNestPoint3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.OptNestPoint3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestPoint3.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestPoint3 unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestPoint3.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_nests.OptNestPoint3Codec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_nests.OptNestPoint3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestPoint3.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptNestPoint3 gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptNestPoint3.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_nests.OptNestPoint3Codec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_nests.OptNestPoint3Codec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptNestPoint3.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypesNests(vecDir, outDir string) (passed, failed int) {
	var p, f int
	p, f = runModel_NestInner(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestCoord(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestIdVal(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestLabel(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestMoney(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestRange32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestAddr(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestPoint3(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptNestInner(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptNestCoord(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptNestIdVal(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptNestLabel(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptNestMoney(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptNestRange32(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptNestAddr(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptNestPoint3(vecDir, outDir); passed += p; failed += f

	return
}
