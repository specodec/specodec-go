package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	all_types "emit_go/emit_gen/all_types"
	all_types_unions "emit_go/emit_gen/all_types_unions"
)


func runModel_OptInner(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptInner mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptInner.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types.OptInnerCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types.OptInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptInner.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptInner json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptInner.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types.OptInnerCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types.OptInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptInner.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptInner unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptInner.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types.OptInnerCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types.OptInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptInner.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptInner gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptInner.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types.OptInnerCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types.OptInnerCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptInner.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ShapeCircle(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Shape_circle mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Shape_circle.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.ShapeCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.ShapeCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Shape_circle.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Shape_circle json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Shape_circle.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.ShapeCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.ShapeCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Shape_circle.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Shape_circle unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Shape_circle.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.ShapeCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.ShapeCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Shape_circle.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Shape_circle gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Shape_circle.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.ShapeCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.ShapeCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Shape_circle.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ShapeRect(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Shape_rect mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Shape_rect.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.ShapeCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.ShapeCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Shape_rect.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Shape_rect json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Shape_rect.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.ShapeCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.ShapeCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Shape_rect.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Shape_rect unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Shape_rect.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.ShapeCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.ShapeCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Shape_rect.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Shape_rect gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Shape_rect.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.ShapeCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.ShapeCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Shape_rect.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_IdentName(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Ident_name mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ident_name.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.IdentCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.IdentCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ident_name.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ident_name json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ident_name.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.IdentCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.IdentCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ident_name.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ident_name unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ident_name.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.IdentCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.IdentCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ident_name.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ident_name gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ident_name.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.IdentCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.IdentCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ident_name.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_IdentNumber(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Ident_number mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ident_number.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.IdentCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.IdentCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ident_number.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ident_number json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ident_number.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.IdentCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.IdentCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ident_number.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ident_number unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ident_number.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.IdentCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.IdentCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ident_number.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Ident_number gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Ident_number.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.IdentCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.IdentCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Ident_number.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ResultMsgOk(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ResultMsg_ok mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ResultMsg_ok.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.ResultMsgCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.ResultMsgCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ResultMsg_ok.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ResultMsg_ok json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ResultMsg_ok.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.ResultMsgCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.ResultMsgCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ResultMsg_ok.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ResultMsg_ok unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ResultMsg_ok.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.ResultMsgCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.ResultMsgCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ResultMsg_ok.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ResultMsg_ok gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ResultMsg_ok.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.ResultMsgCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.ResultMsgCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ResultMsg_ok.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ResultMsgErr(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ResultMsg_err mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ResultMsg_err.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.ResultMsgCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.ResultMsgCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ResultMsg_err.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ResultMsg_err json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ResultMsg_err.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.ResultMsgCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.ResultMsgCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ResultMsg_err.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ResultMsg_err unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ResultMsg_err.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.ResultMsgCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.ResultMsgCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ResultMsg_err.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ResultMsg_err gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ResultMsg_err.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.ResultMsgCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.ResultMsgCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ResultMsg_err.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_TaggedTag(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Tagged_tag mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Tagged_tag.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.TaggedCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.TaggedCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Tagged_tag.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Tagged_tag json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Tagged_tag.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.TaggedCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.TaggedCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Tagged_tag.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Tagged_tag unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Tagged_tag.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.TaggedCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.TaggedCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Tagged_tag.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Tagged_tag gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Tagged_tag.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.TaggedCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.TaggedCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Tagged_tag.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_TaggedScore(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Tagged_score mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Tagged_score.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.TaggedCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.TaggedCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Tagged_score.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Tagged_score json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Tagged_score.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.TaggedCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.TaggedCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Tagged_score.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Tagged_score unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Tagged_score.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.TaggedCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.TaggedCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Tagged_score.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Tagged_score gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Tagged_score.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.TaggedCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.TaggedCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Tagged_score.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_TaggedActive(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("Tagged_active mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Tagged_active.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.TaggedCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.TaggedCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Tagged_active.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Tagged_active json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Tagged_active.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.TaggedCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.TaggedCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Tagged_active.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Tagged_active unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Tagged_active.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.TaggedCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.TaggedCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Tagged_active.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("Tagged_active gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "Tagged_active.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.TaggedCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.TaggedCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "Tagged_active.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptUnionHolderShape(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptUnionHolder_shape mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptUnionHolder_shape.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.OptUnionHolderCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.OptUnionHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptUnionHolder_shape.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptUnionHolder_shape json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptUnionHolder_shape.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.OptUnionHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.OptUnionHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptUnionHolder_shape.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptUnionHolder_shape unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptUnionHolder_shape.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.OptUnionHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.OptUnionHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptUnionHolder_shape.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptUnionHolder_shape gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptUnionHolder_shape.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.OptUnionHolderCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.OptUnionHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptUnionHolder_shape.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_OptUnionHolderId(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("OptUnionHolder_id mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptUnionHolder_id.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.OptUnionHolderCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.OptUnionHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptUnionHolder_id.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptUnionHolder_id json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptUnionHolder_id.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.OptUnionHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.OptUnionHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptUnionHolder_id.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptUnionHolder_id unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptUnionHolder_id.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.OptUnionHolderCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.OptUnionHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptUnionHolder_id.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("OptUnionHolder_id gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "OptUnionHolder_id.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.OptUnionHolderCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.OptUnionHolderCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "OptUnionHolder_id.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_MixedUnionPoint(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("MixedUnion_point mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MixedUnion_point.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.MixedUnionCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.MixedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MixedUnion_point.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MixedUnion_point json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MixedUnion_point.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.MixedUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.MixedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MixedUnion_point.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MixedUnion_point unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MixedUnion_point.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.MixedUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.MixedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MixedUnion_point.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MixedUnion_point gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MixedUnion_point.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.MixedUnionCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.MixedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MixedUnion_point.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_MixedUnionLabel(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("MixedUnion_label mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MixedUnion_label.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.MixedUnionCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.MixedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MixedUnion_label.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MixedUnion_label json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MixedUnion_label.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.MixedUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.MixedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MixedUnion_label.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MixedUnion_label unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MixedUnion_label.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.MixedUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.MixedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MixedUnion_label.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MixedUnion_label gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MixedUnion_label.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.MixedUnionCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.MixedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MixedUnion_label.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_MixedUnionCount(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("MixedUnion_count mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MixedUnion_count.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.MixedUnionCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.MixedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MixedUnion_count.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MixedUnion_count json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MixedUnion_count.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.MixedUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.MixedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MixedUnion_count.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MixedUnion_count unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MixedUnion_count.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.MixedUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.MixedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MixedUnion_count.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("MixedUnion_count gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "MixedUnion_count.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.MixedUnionCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.MixedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "MixedUnion_count.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestedUnionResult(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestedUnion_result mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestedUnion_result.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.NestedUnionCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.NestedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestedUnion_result.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestedUnion_result json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestedUnion_result.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.NestedUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.NestedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestedUnion_result.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestedUnion_result unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestedUnion_result.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.NestedUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.NestedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestedUnion_result.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestedUnion_result gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestedUnion_result.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.NestedUnionCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.NestedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestedUnion_result.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_NestedUnionShape(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("NestedUnion_shape mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestedUnion_shape.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.NestedUnionCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.NestedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestedUnion_shape.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestedUnion_shape json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestedUnion_shape.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.NestedUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.NestedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestedUnion_shape.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestedUnion_shape unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestedUnion_shape.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.NestedUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.NestedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestedUnion_shape.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("NestedUnion_shape gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "NestedUnion_shape.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.NestedUnionCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.NestedUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "NestedUnion_shape.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ScalarUnionS(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ScalarUnion_s mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ScalarUnion_s.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.ScalarUnionCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.ScalarUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ScalarUnion_s.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ScalarUnion_s json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ScalarUnion_s.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.ScalarUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.ScalarUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ScalarUnion_s.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ScalarUnion_s unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ScalarUnion_s.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.ScalarUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.ScalarUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ScalarUnion_s.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ScalarUnion_s gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ScalarUnion_s.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.ScalarUnionCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.ScalarUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ScalarUnion_s.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ScalarUnionI(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ScalarUnion_i mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ScalarUnion_i.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.ScalarUnionCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.ScalarUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ScalarUnion_i.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ScalarUnion_i json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ScalarUnion_i.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.ScalarUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.ScalarUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ScalarUnion_i.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ScalarUnion_i unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ScalarUnion_i.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.ScalarUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.ScalarUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ScalarUnion_i.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ScalarUnion_i gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ScalarUnion_i.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.ScalarUnionCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.ScalarUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ScalarUnion_i.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ScalarUnionF(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ScalarUnion_f mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ScalarUnion_f.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.ScalarUnionCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.ScalarUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ScalarUnion_f.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ScalarUnion_f json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ScalarUnion_f.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.ScalarUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.ScalarUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ScalarUnion_f.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ScalarUnion_f unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ScalarUnion_f.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.ScalarUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.ScalarUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ScalarUnion_f.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ScalarUnion_f gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ScalarUnion_f.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.ScalarUnionCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.ScalarUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ScalarUnion_f.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_ScalarUnionB(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("ScalarUnion_b mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ScalarUnion_b.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_unions.ScalarUnionCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_unions.ScalarUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ScalarUnion_b.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ScalarUnion_b json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ScalarUnion_b.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.ScalarUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.ScalarUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ScalarUnion_b.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ScalarUnion_b unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ScalarUnion_b.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_unions.ScalarUnionCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_unions.ScalarUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ScalarUnion_b.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("ScalarUnion_b gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "ScalarUnion_b.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_unions.ScalarUnionCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_unions.ScalarUnionCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "ScalarUnion_b.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypes(vecDir, outDir string) (passed, failed int) {
	var p, f int
	p, f = runModel_OptInner(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ShapeCircle(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ShapeRect(vecDir, outDir); passed += p; failed += f
	p, f = runModel_IdentName(vecDir, outDir); passed += p; failed += f
	p, f = runModel_IdentNumber(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ResultMsgOk(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ResultMsgErr(vecDir, outDir); passed += p; failed += f
	p, f = runModel_TaggedTag(vecDir, outDir); passed += p; failed += f
	p, f = runModel_TaggedScore(vecDir, outDir); passed += p; failed += f
	p, f = runModel_TaggedActive(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptUnionHolderShape(vecDir, outDir); passed += p; failed += f
	p, f = runModel_OptUnionHolderId(vecDir, outDir); passed += p; failed += f
	p, f = runModel_MixedUnionPoint(vecDir, outDir); passed += p; failed += f
	p, f = runModel_MixedUnionLabel(vecDir, outDir); passed += p; failed += f
	p, f = runModel_MixedUnionCount(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestedUnionResult(vecDir, outDir); passed += p; failed += f
	p, f = runModel_NestedUnionShape(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ScalarUnionS(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ScalarUnionI(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ScalarUnionF(vecDir, outDir); passed += p; failed += f
	p, f = runModel_ScalarUnionB(vecDir, outDir); passed += p; failed += f

	return
}
