package main

import (
	"os"
	"path/filepath"
	specodec "github.com/specodec/specodec-runtime-golang"
	all_types_edge "emit_go/emit_gen/all_types_edge"
)


func runModel_EdgeEmpty(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeEmpty mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeEmpty.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_edge.EdgeEmptyCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_edge.EdgeEmptyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeEmpty.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeEmpty json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeEmpty.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeEmptyCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeEmptyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeEmpty.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeEmpty unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeEmpty.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeEmptyCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeEmptyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeEmpty.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeEmpty gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeEmpty.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_edge.EdgeEmptyCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_edge.EdgeEmptyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeEmpty.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeOneOpt(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeOneOpt mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeOneOpt.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_edge.EdgeOneOptCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_edge.EdgeOneOptCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeOneOpt.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeOneOpt json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeOneOpt.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeOneOptCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeOneOptCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeOneOpt.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeOneOpt unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeOneOpt.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeOneOptCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeOneOptCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeOneOpt.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeOneOpt gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeOneOpt.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_edge.EdgeOneOptCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_edge.EdgeOneOptCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeOneOpt.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeBigNums(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeBigNums mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBigNums.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_edge.EdgeBigNumsCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_edge.EdgeBigNumsCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBigNums.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBigNums json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBigNums.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeBigNumsCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeBigNumsCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBigNums.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBigNums unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBigNums.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeBigNumsCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeBigNumsCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBigNums.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBigNums gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBigNums.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_edge.EdgeBigNumsCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_edge.EdgeBigNumsCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBigNums.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeZeroVals(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeZeroVals mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeZeroVals.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_edge.EdgeZeroValsCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_edge.EdgeZeroValsCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeZeroVals.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeZeroVals json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeZeroVals.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeZeroValsCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeZeroValsCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeZeroVals.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeZeroVals unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeZeroVals.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeZeroValsCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeZeroValsCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeZeroVals.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeZeroVals gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeZeroVals.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_edge.EdgeZeroValsCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_edge.EdgeZeroValsCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeZeroVals.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeNullable(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeNullable mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNullable.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_edge.EdgeNullableCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_edge.EdgeNullableCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNullable.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNullable json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNullable.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeNullableCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeNullableCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNullable.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNullable unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNullable.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeNullableCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeNullableCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNullable.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNullable gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNullable.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_edge.EdgeNullableCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_edge.EdgeNullableCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNullable.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeNegZero(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeNegZero mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNegZero.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_edge.EdgeNegZeroCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_edge.EdgeNegZeroCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNegZero.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNegZero json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNegZero.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeNegZeroCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeNegZeroCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNegZero.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNegZero unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNegZero.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeNegZeroCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeNegZeroCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNegZero.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNegZero gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNegZero.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_edge.EdgeNegZeroCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_edge.EdgeNegZeroCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNegZero.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeNullByte(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeNullByte mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNullByte.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_edge.EdgeNullByteCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_edge.EdgeNullByteCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNullByte.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNullByte json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNullByte.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeNullByteCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeNullByteCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNullByte.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNullByte unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNullByte.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeNullByteCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeNullByteCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNullByte.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeNullByte gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeNullByte.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_edge.EdgeNullByteCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_edge.EdgeNullByteCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeNullByte.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeBoundary(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeBoundary mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBoundary.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_edge.EdgeBoundaryCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_edge.EdgeBoundaryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBoundary.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBoundary json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBoundary.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeBoundaryCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeBoundaryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBoundary.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBoundary unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBoundary.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeBoundaryCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeBoundaryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBoundary.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBoundary gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBoundary.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_edge.EdgeBoundaryCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_edge.EdgeBoundaryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBoundary.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeStrLen(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeStrLen mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeStrLen.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_edge.EdgeStrLenCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_edge.EdgeStrLenCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeStrLen.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeStrLen json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeStrLen.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeStrLenCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeStrLenCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeStrLen.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeStrLen unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeStrLen.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeStrLenCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeStrLenCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeStrLen.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeStrLen gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeStrLen.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_edge.EdgeStrLenCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_edge.EdgeStrLenCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeStrLen.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeBytesLen(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeBytesLen mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBytesLen.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_edge.EdgeBytesLenCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_edge.EdgeBytesLenCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBytesLen.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBytesLen json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBytesLen.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeBytesLenCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeBytesLenCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBytesLen.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBytesLen unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBytesLen.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeBytesLenCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeBytesLenCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBytesLen.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeBytesLen gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeBytesLen.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_edge.EdgeBytesLenCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_edge.EdgeBytesLenCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeBytesLen.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeArrEmpty(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeArrEmpty mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeArrEmpty.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_edge.EdgeArrEmptyCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_edge.EdgeArrEmptyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeArrEmpty.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeArrEmpty json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeArrEmpty.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeArrEmptyCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeArrEmptyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeArrEmpty.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeArrEmpty unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeArrEmpty.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeArrEmptyCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeArrEmptyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeArrEmpty.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeArrEmpty gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeArrEmpty.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_edge.EdgeArrEmptyCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_edge.EdgeArrEmptyCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeArrEmpty.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}

func runModel_EdgeArrBoundary(vecDir, outDir string) (passed, failed int) {
	p, f := tryTest("EdgeArrBoundary mp", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeArrBoundary.msgpack"))
		if err != nil { panic(err) }
		r := specodec.NewMsgPackReader(data)
		obj := all_types_edge.EdgeArrBoundaryCodec.Decode(r)
		w := specodec.NewMsgPackWriter()
		all_types_edge.EdgeArrBoundaryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeArrBoundary.msgpack"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeArrBoundary json", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeArrBoundary.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeArrBoundaryCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeArrBoundaryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeArrBoundary.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeArrBoundary unformatted", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeArrBoundary.unformatted.json"))
		if err != nil { panic(err) }
		r := specodec.NewJsonReader(data)
		obj := all_types_edge.EdgeArrBoundaryCodec.Decode(r)
		w := specodec.NewJsonWriter()
		all_types_edge.EdgeArrBoundaryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeArrBoundary.unformatted.json"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f

	p, f = tryTest("EdgeArrBoundary gron", func() {
		data, err := os.ReadFile(filepath.Join(vecDir, "EdgeArrBoundary.gron"))
		if err != nil { panic(err) }
		r := specodec.NewGronReader(data)
		obj := all_types_edge.EdgeArrBoundaryCodec.Decode(r)
		w := specodec.NewGronWriter()
		all_types_edge.EdgeArrBoundaryCodec.Encode(w, obj)
		err = os.WriteFile(filepath.Join(outDir, "EdgeArrBoundary.gron"), w.ToBytes(), 0644)
		if err != nil { panic(err) }
	})
	passed += p; failed += f
	return
}


func runAllTypesEdge(vecDir, outDir string) (passed, failed int) {
	var p, f int
	p, f = runModel_EdgeEmpty(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeOneOpt(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeBigNums(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeZeroVals(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeNullable(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeNegZero(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeNullByte(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeBoundary(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeStrLen(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeBytesLen(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeArrEmpty(vecDir, outDir); passed += p; failed += f
	p, f = runModel_EdgeArrBoundary(vecDir, outDir); passed += p; failed += f

	return
}
